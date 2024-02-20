package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinitionTemplateLinked struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#policy_template_id VerifiedpermissionsPolicy#policy_template_id}.
	PolicyTemplateId *string `field:"required" json:"policyTemplateId" yaml:"policyTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#principal VerifiedpermissionsPolicy#principal}.
	Principal *VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipal `field:"optional" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#resource VerifiedpermissionsPolicy#resource}.
	Resource *VerifiedpermissionsPolicyDefinitionTemplateLinkedResource `field:"optional" json:"resource" yaml:"resource"`
}

