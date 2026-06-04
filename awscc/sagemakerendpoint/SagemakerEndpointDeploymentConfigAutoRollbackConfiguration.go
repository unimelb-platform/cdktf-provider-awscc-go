package sagemakerendpoint


type SagemakerEndpointDeploymentConfigAutoRollbackConfiguration struct {
	// List of CloudWatch alarms to monitor during the deployment. If any alarm goes off, the deployment is rolled back.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_endpoint#alarms SagemakerEndpoint#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
}

