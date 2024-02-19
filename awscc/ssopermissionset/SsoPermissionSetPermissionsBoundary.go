package ssopermissionset


type SsoPermissionSetPermissionsBoundary struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#customer_managed_policy_reference SsoPermissionSet#customer_managed_policy_reference}.
	CustomerManagedPolicyReference *SsoPermissionSetPermissionsBoundaryCustomerManagedPolicyReference `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// The managed policy to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_permission_set#managed_policy_arn SsoPermissionSet#managed_policy_arn}
	ManagedPolicyArn *string `field:"optional" json:"managedPolicyArn" yaml:"managedPolicyArn"`
}

