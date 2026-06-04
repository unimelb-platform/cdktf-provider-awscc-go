package wisdomaiprompt

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WisdomAiPromptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#api_format WisdomAiPrompt#api_format}.
	ApiFormat *string `field:"required" json:"apiFormat" yaml:"apiFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#model_id WisdomAiPrompt#model_id}.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#template_configuration WisdomAiPrompt#template_configuration}.
	TemplateConfiguration *WisdomAiPromptTemplateConfiguration `field:"required" json:"templateConfiguration" yaml:"templateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#template_type WisdomAiPrompt#template_type}.
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#type WisdomAiPrompt#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#assistant_id WisdomAiPrompt#assistant_id}.
	AssistantId *string `field:"optional" json:"assistantId" yaml:"assistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#description WisdomAiPrompt#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#name WisdomAiPrompt#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_prompt#tags WisdomAiPrompt#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

