package connectphonenumber

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectPhoneNumberConfig struct {
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
	// The ARN of the target the phone number is claimed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#target_arn ConnectPhoneNumber#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The phone number country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#country_code ConnectPhoneNumber#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The description of the phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#description ConnectPhoneNumber#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The phone number prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#prefix ConnectPhoneNumber#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The source phone number arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#source_phone_number_arn ConnectPhoneNumber#source_phone_number_arn}
	SourcePhoneNumberArn *string `field:"optional" json:"sourcePhoneNumberArn" yaml:"sourcePhoneNumberArn"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#tags ConnectPhoneNumber#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The phone number type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_phone_number#type ConnectPhoneNumber#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

