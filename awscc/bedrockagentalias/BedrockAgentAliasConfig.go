package bedrockagentalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockAgentAliasConfig struct {
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
	// Name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent_alias#agent_alias_name BedrockAgentAlias#agent_alias_name}
	AgentAliasName *string `field:"required" json:"agentAliasName" yaml:"agentAliasName"`
	// Identifier for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent_alias#agent_id BedrockAgentAlias#agent_id}
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Description of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent_alias#description BedrockAgentAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Routing configuration for an Agent alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent_alias#routing_configuration BedrockAgentAlias#routing_configuration}
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent_alias#tags BedrockAgentAlias#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

