package ec2networkinsightspath


type Ec2NetworkInsightsPathTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#key Ec2NetworkInsightsPath#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#value Ec2NetworkInsightsPath#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

