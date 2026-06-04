package bedrockblueprint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockBlueprintConfig struct {
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
	// Name of the Blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#blueprint_name BedrockBlueprint#blueprint_name}
	BlueprintName *string `field:"required" json:"blueprintName" yaml:"blueprintName"`
	// Schema of the blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#schema BedrockBlueprint#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Modality Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#type BedrockBlueprint#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// KMS encryption context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#kms_encryption_context BedrockBlueprint#kms_encryption_context}
	KmsEncryptionContext *map[string]*string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// KMS key identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#kms_key_id BedrockBlueprint#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_blueprint#tags BedrockBlueprint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

