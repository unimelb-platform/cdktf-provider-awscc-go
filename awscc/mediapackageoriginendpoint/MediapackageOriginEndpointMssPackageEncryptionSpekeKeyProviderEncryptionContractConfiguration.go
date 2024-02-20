package mediapackageoriginendpoint


type MediapackageOriginEndpointMssPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration struct {
	// A collection of audio encryption presets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#preset_speke_20_audio MediapackageOriginEndpoint#preset_speke_20_audio}
	PresetSpeke20Audio *string `field:"required" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// A collection of video encryption presets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#preset_speke_20_video MediapackageOriginEndpoint#preset_speke_20_video}
	PresetSpeke20Video *string `field:"required" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

