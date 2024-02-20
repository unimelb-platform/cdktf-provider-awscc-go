package pinpointinapptemplate


type PinpointInAppTemplateContentPrimaryBtn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#android PinpointInAppTemplate#android}.
	Android *PinpointInAppTemplateContentPrimaryBtnAndroid `field:"optional" json:"android" yaml:"android"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#default_config PinpointInAppTemplate#default_config}.
	DefaultConfig *PinpointInAppTemplateContentPrimaryBtnDefaultConfig `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#ios PinpointInAppTemplate#ios}.
	Ios *PinpointInAppTemplateContentPrimaryBtnIos `field:"optional" json:"ios" yaml:"ios"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pinpoint_in_app_template#web PinpointInAppTemplate#web}.
	Web *PinpointInAppTemplateContentPrimaryBtnWeb `field:"optional" json:"web" yaml:"web"`
}

