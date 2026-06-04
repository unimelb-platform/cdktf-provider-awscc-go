package wisdommessagetemplateversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WisdomMessageTemplateVersionConfig struct {
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
	// The unqualified Amazon Resource Name (ARN) of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template_version#message_template_arn WisdomMessageTemplateVersion#message_template_arn}
	MessageTemplateArn *string `field:"required" json:"messageTemplateArn" yaml:"messageTemplateArn"`
	// The content SHA256 of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template_version#message_template_content_sha_256 WisdomMessageTemplateVersion#message_template_content_sha_256}
	MessageTemplateContentSha256 *string `field:"optional" json:"messageTemplateContentSha256" yaml:"messageTemplateContentSha256"`
}

