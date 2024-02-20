package medialivemultiplex


type MedialiveMultiplexDestinations struct {
	// Multiplex MediaConnect output destination settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplex#multiplex_media_connect_output_destination_settings MedialiveMultiplex#multiplex_media_connect_output_destination_settings}
	MultiplexMediaConnectOutputDestinationSettings *MedialiveMultiplexDestinationsMultiplexMediaConnectOutputDestinationSettings `field:"optional" json:"multiplexMediaConnectOutputDestinationSettings" yaml:"multiplexMediaConnectOutputDestinationSettings"`
}

