// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverless

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awstypes "github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_opensearchserverless_collection_group", name="Collection Group")
// @Tags(identifierAttribute="arn")
func newCollectionGroupDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &collectionGroupDataSource{}, nil
}

const (
	DSNameCollectionGroup = "Collection Group Data Source"
)

type collectionGroupDataSource struct {
	framework.DataSourceWithModel[collectionGroupDataSourceModel]
}

func (d *collectionGroupDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: framework.ARNAttributeComputedOnly(),
			"created_date": schema.StringAttribute{
				Description: "Date the Collection Group was created.",
				Computed:    true,
			},
			names.AttrDescription: schema.StringAttribute{
				Description: "Description of the collection group.",
				Computed:    true,
			},
			names.AttrID: schema.StringAttribute{
				Description: "ID of the collection group.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName(names.AttrName),
					),
					stringvalidator.ExactlyOneOf(
						path.MatchRelative().AtParent().AtName(names.AttrName),
					),
				},
			},
			names.AttrName: schema.StringAttribute{
				Description: "Name of the collection group.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName(names.AttrID),
					),
				},
			},
			"number_of_collections": schema.Int64Attribute{
				Description: "The number of collections associated with the collection group.",
				Computed:    true,
			},
			"standby_replicas": schema.StringAttribute{
				Description: "Indicates whether standby replicas should be used for the collection group.",
				Computed:    true,
			},
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
		},
	}
}

func (d *collectionGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	conn := d.Meta().OpenSearchServerlessClient(ctx)

	var data collectionGroupDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var out *awstypes.CollectionGroupDetail

	if !data.ID.IsNull() && !data.ID.IsUnknown() {
		output, err := FindCollectionGroupByID(ctx, conn, data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.OpenSearchServerless, create.ErrActionReading, DSNameCollectionGroup, data.ID.String(), err),
				err.Error(),
			)
			return
		}

		out = output
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		output, err := FindCollectionGroupByName(ctx, conn, data.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.OpenSearchServerless, create.ErrActionReading, DSNameCollectionGroup, data.Name.String(), err),
				err.Error(),
			)
			return
		}

		out = output
	}

	resp.Diagnostics.Append(flex.Flatten(ctx, out, &data, flex.WithIgnoredFieldNames([]string{"CreatedDate"}))...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Special handling for Unix time conversion
	data.CreatedDate = flex.StringValueToFramework(ctx, time.UnixMilli(aws.ToInt64(out.CreatedDate)).Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type collectionGroupDataSourceModel struct {
	framework.WithRegionModel
	ARN                 types.String `tfsdk:"arn"`
	CreatedDate         types.String `tfsdk:"created_date"`
	Description         types.String `tfsdk:"description"`
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	NumberOfCollections types.Int64  `tfsdk:"number_of_collections"`
	StandbyReplicas     types.String `tfsdk:"standby_replicas"`
	Tags                tftags.Map   `tfsdk:"tags"`
}
