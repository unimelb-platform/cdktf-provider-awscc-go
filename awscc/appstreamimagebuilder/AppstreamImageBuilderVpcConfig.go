package appstreamimagebuilder


type AppstreamImageBuilderVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_image_builder#security_group_ids AppstreamImageBuilder#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_image_builder#subnet_ids AppstreamImageBuilder#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

