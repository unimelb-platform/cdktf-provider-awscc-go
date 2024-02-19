package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#event_type CloudfrontDistribution#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#function_arn CloudfrontDistribution#function_arn}.
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
}

