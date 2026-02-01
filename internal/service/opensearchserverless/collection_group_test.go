// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchserverless_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfopensearchserverless "github.com/hashicorp/terraform-provider-aws/internal/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccOpenSearchServerlessCollectionGroup_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var collectiongroup types.CollectionGroupDetail
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_opensearchserverless_collection_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.OpenSearchServerlessEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.OpenSearchServerlessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCollectionGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionGroupConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrARN),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "number_of_collections", "0"),
					resource.TestCheckResourceAttr(resourceName, "standby_replicas", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccOpenSearchServerlessCollectionGroup_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var collectiongroup types.CollectionGroupDetail
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_opensearchserverless_collection_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.OpenSearchServerlessEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.OpenSearchServerlessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCollectionGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionGroupConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					acctest.CheckFrameworkResourceDisappears(ctx, t, tfopensearchserverless.ResourceCollectionGroup, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOpenSearchServerlessCollectionGroup_description(t *testing.T) {
	ctx := acctest.Context(t)
	var collectiongroup types.CollectionGroupDetail
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_opensearchserverless_collection_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.OpenSearchServerlessEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.OpenSearchServerlessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCollectionGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionGroupConfig_description(rName, "description1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCollectionGroupConfig_description(rName, "description2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description2"),
				),
			},
		},
	})
}

func TestAccOpenSearchServerlessCollectionGroup_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var collectiongroup types.CollectionGroupDetail
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_opensearchserverless_collection_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.OpenSearchServerlessEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.OpenSearchServerlessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCollectionGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionGroupConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCollectionGroupConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccCollectionGroupConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionGroupExists(ctx, resourceName, &collectiongroup),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccOpenSearchServerlessCollectionGroup_nameValidation(t *testing.T) {
	ctx := acctest.Context(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.OpenSearchServerlessEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.OpenSearchServerlessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckCollectionGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccCollectionGroupConfig_basic("ab"),
				ExpectError: regexache.MustCompile(`string length must be between 3 and 32`),
			},
			{
				Config:      testAccCollectionGroupConfig_basic("Test-Invalid"),
				ExpectError: regexache.MustCompile(`must start with a lowercase letter`),
			},
		},
	})
}

func testAccCheckCollectionGroupDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).OpenSearchServerlessClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_opensearchserverless_collection_group" {
				continue
			}

			_, err := tfopensearchserverless.FindCollectionGroupByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("OpenSearch Serverless Collection Group %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckCollectionGroupExists(ctx context.Context, n string, v *types.CollectionGroupDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).OpenSearchServerlessClient(ctx)

		output, err := tfopensearchserverless.FindCollectionGroupByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCollectionGroupConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_opensearchserverless_collection_group" "test" {
  name             = %[1]q
  standby_replicas = "ENABLED"
}
`, rName)
}

func testAccCollectionGroupConfig_description(rName, description string) string {
	return fmt.Sprintf(`
resource "aws_opensearchserverless_collection_group" "test" {
  name             = %[1]q
  standby_replicas = "ENABLED"
  description      = %[2]q
}
`, rName, description)
}

func testAccCollectionGroupConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_opensearchserverless_collection_group" "test" {
  name             = %[1]q
  standby_replicas = "ENABLED"

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccCollectionGroupConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_opensearchserverless_collection_group" "test" {
  name             = %[1]q
  standby_replicas = "ENABLED"

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
