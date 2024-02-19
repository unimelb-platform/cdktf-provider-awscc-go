package kendradatasource


type KendraDataSourceDataSourceConfigurationOneDriveConfigurationOneDriveUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#one_drive_user_list KendraDataSource#one_drive_user_list}.
	OneDriveUserList *[]*string `field:"optional" json:"oneDriveUserList" yaml:"oneDriveUserList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#one_drive_user_s3_path KendraDataSource#one_drive_user_s3_path}.
	OneDriveUserS3Path *KendraDataSourceDataSourceConfigurationOneDriveConfigurationOneDriveUsersOneDriveUserS3Path `field:"optional" json:"oneDriveUserS3Path" yaml:"oneDriveUserS3Path"`
}

