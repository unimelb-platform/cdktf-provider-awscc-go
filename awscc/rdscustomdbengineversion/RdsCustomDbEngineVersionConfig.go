package rdscustomdbengineversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsCustomDbEngineVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of an Amazon S3 bucket that contains database installation files for your CEV.
	//
	// For example, a valid bucket name is `my-custom-installation-files`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_bucket_name RdsCustomDbEngineVersion#database_installation_files_s3_bucket_name}
	DatabaseInstallationFilesS3BucketName *string `field:"required" json:"databaseInstallationFilesS3BucketName" yaml:"databaseInstallationFilesS3BucketName"`
	// The database engine to use for your custom engine version (CEV). The only supported value is `custom-oracle-ee`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#engine RdsCustomDbEngineVersion#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The name of your CEV.
	//
	// The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#engine_version RdsCustomDbEngineVersion#engine_version}
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// The Amazon S3 directory that contains the database installation files for your CEV.
	//
	// For example, a valid bucket name is `123456789012/cev1`. If this setting isn't specified, no prefix is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#database_installation_files_s3_prefix RdsCustomDbEngineVersion#database_installation_files_s3_prefix}
	DatabaseInstallationFilesS3Prefix *string `field:"optional" json:"databaseInstallationFilesS3Prefix" yaml:"databaseInstallationFilesS3Prefix"`
	// An optional description of your CEV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#description RdsCustomDbEngineVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS KMS key identifier for an encrypted CEV.
	//
	// A symmetric KMS key is required for RDS Custom, but optional for Amazon RDS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#kms_key_id RdsCustomDbEngineVersion#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#manifest RdsCustomDbEngineVersion#manifest}
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// The availability status to be assigned to the CEV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#status RdsCustomDbEngineVersion#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_custom_db_engine_version#tags RdsCustomDbEngineVersion#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

