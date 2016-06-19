package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	//"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceAwsSubnet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceAwsSubnetConfig,
			},
		},
	})
}

const testAccDataSourceAwsSubnetConfig = `
resource "aws_vpc" "test" {
  cidr_block = "172.16.0.0/16"
}

resource "aws_subnet" "test" {
  vpc_id     = "${aws_vpc.test.id}"
  cidr_block = "172.16.1.0/24"
}

data "aws_subnet" "test" {
  id = "${aws_subnet.test.id}"
}
`
