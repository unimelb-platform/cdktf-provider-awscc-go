package wisdomaiguardrailversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WisdomAiGuardrailVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail_version#ai_guardrail_id WisdomAiGuardrailVersion#ai_guardrail_id}.
	AiGuardrailId *string `field:"required" json:"aiGuardrailId" yaml:"aiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail_version#assistant_id WisdomAiGuardrailVersion#assistant_id}.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail_version#modified_time_seconds WisdomAiGuardrailVersion#modified_time_seconds}.
	ModifiedTimeSeconds *float64 `field:"optional" json:"modifiedTimeSeconds" yaml:"modifiedTimeSeconds"`
}

