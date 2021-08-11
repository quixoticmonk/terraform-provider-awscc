// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSServiceCatalogServiceAction_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ServiceCatalog::ServiceAction", "aws_servicecatalog_service_action", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}