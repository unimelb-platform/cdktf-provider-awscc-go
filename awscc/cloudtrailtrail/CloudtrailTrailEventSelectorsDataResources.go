package cloudtrailtrail


type CloudtrailTrailEventSelectorsDataResources struct {
	// The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail#type CloudtrailTrail#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail#values CloudtrailTrail#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

