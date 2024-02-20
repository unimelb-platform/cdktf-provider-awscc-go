package dmsmigrationproject


type DmsMigrationProjectSchemaConversionApplicationAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_migration_project#s3_bucket_path DmsMigrationProject#s3_bucket_path}.
	S3BucketPath *string `field:"optional" json:"s3BucketPath" yaml:"s3BucketPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_migration_project#s3_bucket_role_arn DmsMigrationProject#s3_bucket_role_arn}.
	S3BucketRoleArn *string `field:"optional" json:"s3BucketRoleArn" yaml:"s3BucketRoleArn"`
}

