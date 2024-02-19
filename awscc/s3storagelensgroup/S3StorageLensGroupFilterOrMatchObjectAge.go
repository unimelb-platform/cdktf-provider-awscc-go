package s3storagelensgroup


type S3StorageLensGroupFilterOrMatchObjectAge struct {
	// Minimum object age to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#days_greater_than S3StorageLensGroup#days_greater_than}
	DaysGreaterThan *float64 `field:"optional" json:"daysGreaterThan" yaml:"daysGreaterThan"`
	// Maximum object age to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#days_less_than S3StorageLensGroup#days_less_than}
	DaysLessThan *float64 `field:"optional" json:"daysLessThan" yaml:"daysLessThan"`
}

