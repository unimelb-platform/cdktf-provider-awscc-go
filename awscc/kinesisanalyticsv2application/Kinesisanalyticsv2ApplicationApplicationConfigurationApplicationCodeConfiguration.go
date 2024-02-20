package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration struct {
	// The location and type of the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#code_content Kinesisanalyticsv2Application#code_content}
	CodeContent *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent `field:"required" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#code_content_type Kinesisanalyticsv2Application#code_content_type}
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
}

