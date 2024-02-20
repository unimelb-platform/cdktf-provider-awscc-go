package sesvdmattributes


type SesVdmAttributesGuardianAttributes struct {
	// Whether emails sent from this account have optimized delivery algorithm enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_vdm_attributes#optimized_shared_delivery SesVdmAttributes#optimized_shared_delivery}
	OptimizedSharedDelivery *string `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

