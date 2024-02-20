package rdsintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsIntegrationConfig struct {
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
	// The Amazon Resource Name (ARN) of the Aurora DB cluster to use as the source for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#source_arn RdsIntegration#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The ARN of the Redshift data warehouse to use as the target for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#target_arn RdsIntegration#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An optional set of non-secret key–value pairs that contains additional contextual information about the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#additional_encryption_context RdsIntegration#additional_encryption_context}
	AdditionalEncryptionContext *map[string]*string `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// The name of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#integration_name RdsIntegration#integration_name}
	IntegrationName *string `field:"optional" json:"integrationName" yaml:"integrationName"`
	// An optional AWS Key Management System (AWS KMS) key ARN for the key used to to encrypt the integration.
	//
	// The resource accepts the key ID and the key ARN forms. The key ID form can be used if the KMS key is owned by te same account. If the KMS key belongs to a different account than the calling account, the full key ARN must be specified. Do not use the key alias or the key alias ARN as this will cause a false drift of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#kms_key_id RdsIntegration#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_integration#tags RdsIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

