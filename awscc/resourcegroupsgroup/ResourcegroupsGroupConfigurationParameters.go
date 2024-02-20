package resourcegroupsgroup


type ResourcegroupsGroupConfigurationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#name ResourcegroupsGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#values ResourcegroupsGroup#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

