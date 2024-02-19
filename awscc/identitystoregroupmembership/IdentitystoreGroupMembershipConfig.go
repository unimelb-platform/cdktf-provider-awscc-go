package identitystoregroupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentitystoreGroupMembershipConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The unique identifier for a group in the identity store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/identitystore_group_membership#group_id IdentitystoreGroupMembership#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The globally unique identifier for the identity store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/identitystore_group_membership#identity_store_id IdentitystoreGroupMembership#identity_store_id}
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// An object containing the identifier of a group member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/identitystore_group_membership#member_id IdentitystoreGroupMembership#member_id}
	MemberId *IdentitystoreGroupMembershipMemberId `field:"required" json:"memberId" yaml:"memberId"`
}

