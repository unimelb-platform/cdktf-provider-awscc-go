package greengrassv2deployment


type Greengrassv2DeploymentComponentsConfigurationUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#merge Greengrassv2Deployment#merge}.
	Merge *string `field:"optional" json:"merge" yaml:"merge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#reset Greengrassv2Deployment#reset}.
	Reset *[]*string `field:"optional" json:"reset" yaml:"reset"`
}

