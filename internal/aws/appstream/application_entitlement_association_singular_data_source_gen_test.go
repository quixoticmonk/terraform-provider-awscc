// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAppStreamApplicationEntitlementAssociationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AppStream::ApplicationEntitlementAssociation", "awscc_appstream_application_entitlement_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSAppStreamApplicationEntitlementAssociationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AppStream::ApplicationEntitlementAssociation", "awscc_appstream_application_entitlement_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}