package datasynclocationfsxlustre


type DatasyncLocationFsxLustreTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_lustre#key DatasyncLocationFsxLustre#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_lustre#value DatasyncLocationFsxLustre#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

