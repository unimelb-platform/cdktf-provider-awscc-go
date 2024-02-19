package datasyncstoragesystem


type DatasyncStorageSystemServerCredentials struct {
	// The password for your storage system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#password DatasyncStorageSystem#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username for your storage system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#username DatasyncStorageSystem#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

