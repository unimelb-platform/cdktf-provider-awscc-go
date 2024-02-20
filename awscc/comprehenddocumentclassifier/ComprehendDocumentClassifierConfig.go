package comprehenddocumentclassifier

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComprehendDocumentClassifierConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#data_access_role_arn ComprehendDocumentClassifier#data_access_role_arn}.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#document_classifier_name ComprehendDocumentClassifier#document_classifier_name}.
	DocumentClassifierName *string `field:"required" json:"documentClassifierName" yaml:"documentClassifierName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#input_data_config ComprehendDocumentClassifier#input_data_config}.
	InputDataConfig *ComprehendDocumentClassifierInputDataConfig `field:"required" json:"inputDataConfig" yaml:"inputDataConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#language_code ComprehendDocumentClassifier#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#mode ComprehendDocumentClassifier#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#model_kms_key_id ComprehendDocumentClassifier#model_kms_key_id}.
	ModelKmsKeyId *string `field:"optional" json:"modelKmsKeyId" yaml:"modelKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#model_policy ComprehendDocumentClassifier#model_policy}.
	ModelPolicy *string `field:"optional" json:"modelPolicy" yaml:"modelPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#output_data_config ComprehendDocumentClassifier#output_data_config}.
	OutputDataConfig *ComprehendDocumentClassifierOutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#tags ComprehendDocumentClassifier#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#version_name ComprehendDocumentClassifier#version_name}.
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#volume_kms_key_id ComprehendDocumentClassifier#volume_kms_key_id}.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_document_classifier#vpc_config ComprehendDocumentClassifier#vpc_config}.
	VpcConfig *ComprehendDocumentClassifierVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

