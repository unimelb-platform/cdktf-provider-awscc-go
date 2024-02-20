package ec2networkinsightsaccessscope


type Ec2NetworkInsightsAccessScopeMatchPathsDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#packet_header_statement Ec2NetworkInsightsAccessScope#packet_header_statement}.
	PacketHeaderStatement *Ec2NetworkInsightsAccessScopeMatchPathsDestinationPacketHeaderStatement `field:"optional" json:"packetHeaderStatement" yaml:"packetHeaderStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#resource_statement Ec2NetworkInsightsAccessScope#resource_statement}.
	ResourceStatement *Ec2NetworkInsightsAccessScopeMatchPathsDestinationResourceStatement `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

