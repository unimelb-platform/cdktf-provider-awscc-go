package datasyncstoragesystem


type DatasyncStorageSystemTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#key DatasyncStorageSystem#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#value DatasyncStorageSystem#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

