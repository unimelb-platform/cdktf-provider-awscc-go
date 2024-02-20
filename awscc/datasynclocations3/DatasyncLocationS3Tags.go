package datasynclocations3


type DatasyncLocationS3Tags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#key DatasyncLocationS3#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#value DatasyncLocationS3#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

