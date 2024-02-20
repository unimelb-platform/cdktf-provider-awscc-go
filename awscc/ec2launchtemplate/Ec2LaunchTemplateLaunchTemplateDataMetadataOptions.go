package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataMetadataOptions struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#http_endpoint Ec2LaunchTemplate#http_endpoint}
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#http_protocol_ipv_6 Ec2LaunchTemplate#http_protocol_ipv_6}
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#http_put_response_hop_limit Ec2LaunchTemplate#http_put_response_hop_limit}
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// IMDSv2 uses token-backed sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#http_tokens Ec2LaunchTemplate#http_tokens}
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Set to enabled to allow access to instance tags from the instance metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_metadata_tags Ec2LaunchTemplate#instance_metadata_tags}
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

