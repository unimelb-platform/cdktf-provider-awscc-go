package identitystoregroupmembership


type IdentitystoreGroupMembershipMemberId struct {
	// The identifier for a user in the identity store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/identitystore_group_membership#user_id IdentitystoreGroupMembership#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

