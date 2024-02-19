package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinitionStatic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#statement VerifiedpermissionsPolicy#statement}.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#description VerifiedpermissionsPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

