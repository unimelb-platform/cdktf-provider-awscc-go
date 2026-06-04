package connectinstance


type ConnectInstanceAttributes struct {
	// Mandatory element which enables inbound calls on new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#inbound_calls ConnectInstance#inbound_calls}
	InboundCalls interface{} `field:"required" json:"inboundCalls" yaml:"inboundCalls"`
	// Mandatory element which enables outbound calls on new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#outbound_calls ConnectInstance#outbound_calls}
	OutboundCalls interface{} `field:"required" json:"outboundCalls" yaml:"outboundCalls"`
	// Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#auto_resolve_best_voices ConnectInstance#auto_resolve_best_voices}
	AutoResolveBestVoices interface{} `field:"optional" json:"autoResolveBestVoices" yaml:"autoResolveBestVoices"`
	// Boolean flag which enables CONTACTFLOW_LOGS on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#contactflow_logs ConnectInstance#contactflow_logs}
	ContactflowLogs interface{} `field:"optional" json:"contactflowLogs" yaml:"contactflowLogs"`
	// Boolean flag which enables CONTACT_LENS on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#contact_lens ConnectInstance#contact_lens}
	ContactLens interface{} `field:"optional" json:"contactLens" yaml:"contactLens"`
	// Boolean flag which enables EARLY_MEDIA on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#early_media ConnectInstance#early_media}
	EarlyMedia interface{} `field:"optional" json:"earlyMedia" yaml:"earlyMedia"`
	// Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_instance#use_custom_tts_voices ConnectInstance#use_custom_tts_voices}
	UseCustomTtsVoices interface{} `field:"optional" json:"useCustomTtsVoices" yaml:"useCustomTtsVoices"`
}

