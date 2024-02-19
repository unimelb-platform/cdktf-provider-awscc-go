package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupEndpointConfigurations struct {
	// Id of the endpoint.
	//
	// For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#endpoint_id GlobalacceleratorEndpointGroup#endpoint_id}
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Attachment ARN that provides access control to the cross account endpoint.
	//
	// Not required for resources hosted in the same account as the endpoint group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#attachment_arn GlobalacceleratorEndpointGroup#attachment_arn}
	AttachmentArn *string `field:"optional" json:"attachmentArn" yaml:"attachmentArn"`
	// true if client ip should be preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#client_ip_preservation_enabled GlobalacceleratorEndpointGroup#client_ip_preservation_enabled}
	ClientIpPreservationEnabled interface{} `field:"optional" json:"clientIpPreservationEnabled" yaml:"clientIpPreservationEnabled"`
	// The weight for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#weight GlobalacceleratorEndpointGroup#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

