package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#event_type CloudfrontDistribution#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#include_body CloudfrontDistribution#include_body}.
	IncludeBody interface{} `field:"optional" json:"includeBody" yaml:"includeBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#lambda_function_arn CloudfrontDistribution#lambda_function_arn}.
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

