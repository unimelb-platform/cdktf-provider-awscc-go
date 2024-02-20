package ec2natgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NatGatewayConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#subnet_id Ec2NatGateway#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#allocation_id Ec2NatGateway#allocation_id}.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#connectivity_type Ec2NatGateway#connectivity_type}.
	ConnectivityType *string `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#max_drain_duration_seconds Ec2NatGateway#max_drain_duration_seconds}.
	MaxDrainDurationSeconds *float64 `field:"optional" json:"maxDrainDurationSeconds" yaml:"maxDrainDurationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#private_ip_address Ec2NatGateway#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#secondary_allocation_ids Ec2NatGateway#secondary_allocation_ids}.
	SecondaryAllocationIds *[]*string `field:"optional" json:"secondaryAllocationIds" yaml:"secondaryAllocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#secondary_private_ip_address_count Ec2NatGateway#secondary_private_ip_address_count}.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#secondary_private_ip_addresses Ec2NatGateway#secondary_private_ip_addresses}.
	SecondaryPrivateIpAddresses *[]*string `field:"optional" json:"secondaryPrivateIpAddresses" yaml:"secondaryPrivateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#tags Ec2NatGateway#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

