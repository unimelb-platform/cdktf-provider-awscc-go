package cloudformationpublictypeversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationPublicTypeVersionConfig struct {
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
	// The Amazon Resource Number (ARN) of the extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_public_type_version#arn CloudformationPublicTypeVersion#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// A url to the S3 bucket where logs for the testType run will be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_public_type_version#log_delivery_bucket CloudformationPublicTypeVersion#log_delivery_bucket}
	LogDeliveryBucket *string `field:"optional" json:"logDeliveryBucket" yaml:"logDeliveryBucket"`
	// The version number of a public third-party extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_public_type_version#public_version_number CloudformationPublicTypeVersion#public_version_number}
	PublicVersionNumber *string `field:"optional" json:"publicVersionNumber" yaml:"publicVersionNumber"`
	// The kind of extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_public_type_version#type CloudformationPublicTypeVersion#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_public_type_version#type_name CloudformationPublicTypeVersion#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

