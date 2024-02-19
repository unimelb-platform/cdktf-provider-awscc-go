package b2bicapability


type B2BiCapabilityConfigurationEdi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_capability#input_location B2BiCapability#input_location}.
	InputLocation *B2BiCapabilityConfigurationEdiInputLocation `field:"required" json:"inputLocation" yaml:"inputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_capability#output_location B2BiCapability#output_location}.
	OutputLocation *B2BiCapabilityConfigurationEdiOutputLocation `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_capability#transformer_id B2BiCapability#transformer_id}.
	TransformerId *string `field:"required" json:"transformerId" yaml:"transformerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_capability#type B2BiCapability#type}.
	Type *B2BiCapabilityConfigurationEdiType `field:"required" json:"type" yaml:"type"`
}

