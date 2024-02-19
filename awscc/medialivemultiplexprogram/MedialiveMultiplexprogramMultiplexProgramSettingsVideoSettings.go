package medialivemultiplexprogram


type MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings struct {
	// The constant bitrate configuration for the video encode. When this field is defined, StatmuxSettings must be undefined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#constant_bitrate MedialiveMultiplexprogram#constant_bitrate}
	ConstantBitrate *float64 `field:"optional" json:"constantBitrate" yaml:"constantBitrate"`
	// Statmux rate control settings. When this field is defined, ConstantBitrate must be undefined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#statmux_settings MedialiveMultiplexprogram#statmux_settings}
	StatmuxSettings *MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettings `field:"optional" json:"statmuxSettings" yaml:"statmuxSettings"`
}

