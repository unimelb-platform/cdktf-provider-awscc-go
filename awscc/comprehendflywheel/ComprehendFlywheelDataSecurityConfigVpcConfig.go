package comprehendflywheel


type ComprehendFlywheelDataSecurityConfigVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#security_group_ids ComprehendFlywheel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#subnets ComprehendFlywheel#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

