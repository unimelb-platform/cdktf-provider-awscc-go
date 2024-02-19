package macieallowlist


type MacieAllowListCriteriaS3WordsListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#bucket_name MacieAllowList#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#object_key MacieAllowList#object_key}.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
}

