package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionNetworkConfigVpcConfig struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the Subnets field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_data_quality_job_definition#security_group_ids SagemakerDataQualityJobDefinition#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_data_quality_job_definition#subnets SagemakerDataQualityJobDefinition#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

