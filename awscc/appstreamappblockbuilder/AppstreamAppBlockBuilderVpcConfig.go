package appstreamappblockbuilder


type AppstreamAppBlockBuilderVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block_builder#security_group_ids AppstreamAppBlockBuilder#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block_builder#subnet_ids AppstreamAppBlockBuilder#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

