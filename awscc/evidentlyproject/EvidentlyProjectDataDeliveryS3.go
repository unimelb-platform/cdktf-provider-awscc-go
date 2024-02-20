package evidentlyproject


type EvidentlyProjectDataDeliveryS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#bucket_name EvidentlyProject#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#prefix EvidentlyProject#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

