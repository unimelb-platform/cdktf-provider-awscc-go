package robomakerrobot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RobomakerRobotConfig struct {
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
	// The target architecture of the robot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot#architecture RobomakerRobot#architecture}
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Greengrass group id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot#greengrass_group_id RobomakerRobot#greengrass_group_id}
	GreengrassGroupId *string `field:"required" json:"greengrassGroupId" yaml:"greengrassGroupId"`
	// The Amazon Resource Name (ARN) of the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot#fleet RobomakerRobot#fleet}
	Fleet *string `field:"optional" json:"fleet" yaml:"fleet"`
	// The name for the robot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot#name RobomakerRobot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot#tags RobomakerRobot#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

