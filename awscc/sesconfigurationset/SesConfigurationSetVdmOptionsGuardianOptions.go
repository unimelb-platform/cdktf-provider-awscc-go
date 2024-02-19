package sesconfigurationset


type SesConfigurationSetVdmOptionsGuardianOptions struct {
	// Whether emails sent with this configuration set have optimized delivery algorithm enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#optimized_shared_delivery SesConfigurationSet#optimized_shared_delivery}
	OptimizedSharedDelivery *string `field:"required" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

