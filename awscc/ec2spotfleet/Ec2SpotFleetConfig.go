package ec2spotfleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2SpotFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_spot_fleet#spot_fleet_request_config_data Ec2SpotFleet#spot_fleet_request_config_data}.
	SpotFleetRequestConfigData *Ec2SpotFleetSpotFleetRequestConfigData `field:"required" json:"spotFleetRequestConfigData" yaml:"spotFleetRequestConfigData"`
}

