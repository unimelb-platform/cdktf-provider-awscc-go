package b2bicapability


type B2BiCapabilityConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_capability#edi B2BiCapability#edi}.
	Edi *B2BiCapabilityConfigurationEdi `field:"optional" json:"edi" yaml:"edi"`
}

