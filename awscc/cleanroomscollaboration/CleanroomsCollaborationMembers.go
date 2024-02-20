package cleanroomscollaboration


type CleanroomsCollaborationMembers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#account_id CleanroomsCollaboration#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#display_name CleanroomsCollaboration#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#member_abilities CleanroomsCollaboration#member_abilities}.
	MemberAbilities *[]*string `field:"required" json:"memberAbilities" yaml:"memberAbilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#payment_configuration CleanroomsCollaboration#payment_configuration}.
	PaymentConfiguration *CleanroomsCollaborationMembersPaymentConfiguration `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
}

