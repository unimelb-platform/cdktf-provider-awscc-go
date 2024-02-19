package ec2networkinsightspath


type Ec2NetworkInsightsPathFilterAtSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#destination_address Ec2NetworkInsightsPath#destination_address}.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#destination_port_range Ec2NetworkInsightsPath#destination_port_range}.
	DestinationPortRange *Ec2NetworkInsightsPathFilterAtSourceDestinationPortRange `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#source_address Ec2NetworkInsightsPath#source_address}.
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_path#source_port_range Ec2NetworkInsightsPath#source_port_range}.
	SourcePortRange *Ec2NetworkInsightsPathFilterAtSourceSourcePortRange `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

