package memorydbuser


type MemorydbUserAuthenticationMode struct {
	// Passwords used for this user account. You can create up to two passwords for each user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_user#passwords MemorydbUser#passwords}
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// Type of authentication strategy for this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_user#type MemorydbUser#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

