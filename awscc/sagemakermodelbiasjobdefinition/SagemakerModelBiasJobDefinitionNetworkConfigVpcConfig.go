package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionNetworkConfigVpcConfig struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the Subnets field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#security_group_ids SagemakerModelBiasJobDefinition#security_group_ids}
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#subnets SagemakerModelBiasJobDefinition#subnets}
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

