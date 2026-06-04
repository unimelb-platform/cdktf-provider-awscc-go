package bedrockapplicationinferenceprofile


type BedrockApplicationInferenceProfileModelSource struct {
	// Source arns for a custom inference profile to copy its regional load balancing config from.
	//
	// This
	// can either be a foundation model or predefined inference profile ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_application_inference_profile#copy_from BedrockApplicationInferenceProfile#copy_from}
	CopyFrom *string `field:"optional" json:"copyFrom" yaml:"copyFrom"`
}

