package eksaccessentry


type EksAccessEntryAccessPoliciesAccessScope struct {
	// The type of the access scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_access_entry#type EksAccessEntry#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The namespaces to associate with the access scope. Only specify if Type is set to 'namespace'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_access_entry#namespaces EksAccessEntry#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

