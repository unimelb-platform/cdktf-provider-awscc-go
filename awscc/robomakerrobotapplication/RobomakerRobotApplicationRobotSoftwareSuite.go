package robomakerrobotapplication


type RobomakerRobotApplicationRobotSoftwareSuite struct {
	// The name of robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot_application#name RobomakerRobotApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of robot software suite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot_application#version RobomakerRobotApplication#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

