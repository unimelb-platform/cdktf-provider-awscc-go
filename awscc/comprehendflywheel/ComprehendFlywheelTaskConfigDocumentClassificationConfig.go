package comprehendflywheel


type ComprehendFlywheelTaskConfigDocumentClassificationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#mode ComprehendFlywheel#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#labels ComprehendFlywheel#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

