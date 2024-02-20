package osispipeline


type OsisPipelineVpcOptions struct {
	// A list of subnet IDs associated with the VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#subnet_ids OsisPipeline#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// A list of security groups associated with the VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#security_group_ids OsisPipeline#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

