package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes struct {
	// Indicates whether audio input is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allow_audio_input LexBot#allow_audio_input}
	AllowAudioInput interface{} `field:"required" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Indicates whether DTMF input is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allow_dtmf_input LexBot#allow_dtmf_input}
	AllowDtmfInput interface{} `field:"required" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}

