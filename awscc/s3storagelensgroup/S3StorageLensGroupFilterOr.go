package s3storagelensgroup


type S3StorageLensGroupFilterOr struct {
	// Filter to match any of the specified prefixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#match_any_prefix S3StorageLensGroup#match_any_prefix}
	MatchAnyPrefix *[]*string `field:"optional" json:"matchAnyPrefix" yaml:"matchAnyPrefix"`
	// Filter to match any of the specified suffixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#match_any_suffix S3StorageLensGroup#match_any_suffix}
	MatchAnySuffix *[]*string `field:"optional" json:"matchAnySuffix" yaml:"matchAnySuffix"`
	// Filter to match any of the specified object tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#match_any_tag S3StorageLensGroup#match_any_tag}
	MatchAnyTag interface{} `field:"optional" json:"matchAnyTag" yaml:"matchAnyTag"`
	// Filter to match all of the specified values for the minimum and maximum object age.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#match_object_age S3StorageLensGroup#match_object_age}
	MatchObjectAge *S3StorageLensGroupFilterOrMatchObjectAge `field:"optional" json:"matchObjectAge" yaml:"matchObjectAge"`
	// Filter to match all of the specified values for the minimum and maximum object size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#match_object_size S3StorageLensGroup#match_object_size}
	MatchObjectSize *S3StorageLensGroupFilterOrMatchObjectSize `field:"optional" json:"matchObjectSize" yaml:"matchObjectSize"`
}

