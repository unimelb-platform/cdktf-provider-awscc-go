package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy struct {
	// Network ACL entry set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fms_policy#network_acl_entry_set FmsPolicy#network_acl_entry_set}
	NetworkAclEntrySet *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySet `field:"optional" json:"networkAclEntrySet" yaml:"networkAclEntrySet"`
}

