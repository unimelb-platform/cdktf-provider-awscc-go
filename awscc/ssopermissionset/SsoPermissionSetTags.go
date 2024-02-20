package ssopermissionset


type SsoPermissionSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#key SsoPermissionSet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#value SsoPermissionSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

