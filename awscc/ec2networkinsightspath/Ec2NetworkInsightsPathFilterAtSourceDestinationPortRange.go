package ec2networkinsightspath


type Ec2NetworkInsightsPathFilterAtSourceDestinationPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#from_port Ec2NetworkInsightsPath#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#to_port Ec2NetworkInsightsPath#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

