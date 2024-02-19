package datazonedomain


type DatazoneDomainSingleSignOn struct {
	// The type of single sign-on in Amazon DataZone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#type DatazoneDomain#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The single sign-on user assignment in Amazon DataZone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#user_assignment DatazoneDomain#user_assignment}
	UserAssignment *string `field:"optional" json:"userAssignment" yaml:"userAssignment"`
}

