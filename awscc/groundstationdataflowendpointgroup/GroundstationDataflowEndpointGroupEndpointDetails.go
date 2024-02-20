package groundstationdataflowendpointgroup


type GroundstationDataflowEndpointGroupEndpointDetails struct {
	// Information about AwsGroundStationAgentEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#aws_ground_station_agent_endpoint GroundstationDataflowEndpointGroup#aws_ground_station_agent_endpoint}
	AwsGroundStationAgentEndpoint *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpoint `field:"optional" json:"awsGroundStationAgentEndpoint" yaml:"awsGroundStationAgentEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#endpoint GroundstationDataflowEndpointGroup#endpoint}.
	Endpoint *GroundstationDataflowEndpointGroupEndpointDetailsEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#security_details GroundstationDataflowEndpointGroup#security_details}.
	SecurityDetails *GroundstationDataflowEndpointGroupEndpointDetailsSecurityDetails `field:"optional" json:"securityDetails" yaml:"securityDetails"`
}

