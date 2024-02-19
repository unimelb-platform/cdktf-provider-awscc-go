package ec2networkinsightsaccessscope


type Ec2NetworkInsightsAccessScopeTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#key Ec2NetworkInsightsAccessScope#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#value Ec2NetworkInsightsAccessScope#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

