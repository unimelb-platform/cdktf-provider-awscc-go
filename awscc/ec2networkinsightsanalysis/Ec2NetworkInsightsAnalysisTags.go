package ec2networkinsightsanalysis


type Ec2NetworkInsightsAnalysisTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_analysis#key Ec2NetworkInsightsAnalysis#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_analysis#value Ec2NetworkInsightsAnalysis#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

