package robomakerrobotapplication


type RobomakerRobotApplicationSources struct {
	// The architecture of robot application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot_application#architecture RobomakerRobotApplication#architecture}
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Arn of the S3Bucket that stores the robot application source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot_application#s3_bucket RobomakerRobotApplication#s3_bucket}
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 key of robot application source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_robot_application#s3_key RobomakerRobotApplication#s3_key}
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

