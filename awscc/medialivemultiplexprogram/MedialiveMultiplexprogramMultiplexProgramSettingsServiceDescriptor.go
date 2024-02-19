package medialivemultiplexprogram


type MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor struct {
	// Name of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#provider_name MedialiveMultiplexprogram#provider_name}
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#service_name MedialiveMultiplexprogram#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

