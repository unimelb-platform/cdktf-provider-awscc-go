package robomakersimulationapplication


type RobomakerSimulationApplicationSources struct {
	// The target processor architecture for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application#architecture RobomakerSimulationApplication#architecture}
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Amazon S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application#s3_bucket RobomakerSimulationApplication#s3_bucket}
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 object key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application#s3_key RobomakerSimulationApplication#s3_key}
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

