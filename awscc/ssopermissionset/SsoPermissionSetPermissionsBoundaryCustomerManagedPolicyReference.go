package ssopermissionset


type SsoPermissionSetPermissionsBoundaryCustomerManagedPolicyReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#name SsoPermissionSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#path SsoPermissionSet#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

