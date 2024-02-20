package datasyncstoragesystem


type DatasyncStorageSystemServerConfiguration struct {
	// The domain name or IP address of the storage system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#server_hostname DatasyncStorageSystem#server_hostname}
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// The network port needed to access the system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#server_port DatasyncStorageSystem#server_port}
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
}

