package securityhubstandard


type SecurityhubStandardDisabledStandardsControls struct {
	// the Arn for the standard control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_standard#standards_control_arn SecurityhubStandard#standards_control_arn}
	StandardsControlArn *string `field:"required" json:"standardsControlArn" yaml:"standardsControlArn"`
	// the reason the standard control is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/securityhub_standard#reason SecurityhubStandard#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

