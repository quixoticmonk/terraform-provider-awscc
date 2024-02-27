// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package arczonalshift_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSARCZonalShiftZonalAutoshiftConfigurationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ARCZonalShift::ZonalAutoshiftConfiguration", "awscc_arczonalshift_zonal_autoshift_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithEmptyResourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "id", td.ResourceName, "id"),
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "arn", td.ResourceName, "arn"),
			),
		},
	})
}

func TestAccAWSARCZonalShiftZonalAutoshiftConfigurationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ARCZonalShift::ZonalAutoshiftConfiguration", "awscc_arczonalshift_zonal_autoshift_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}