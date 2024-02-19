package datasynclocationsmb


type DatasyncLocationSmbTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#key DatasyncLocationSmb#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#value DatasyncLocationSmb#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

