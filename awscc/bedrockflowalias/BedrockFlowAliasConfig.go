package bedrockflowalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockFlowAliasConfig struct {
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
	// Arn representation of the Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#flow_arn BedrockFlowAlias#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// Name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#name BedrockFlowAlias#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Routing configuration for a Flow alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#routing_configuration BedrockFlowAlias#routing_configuration}
	RoutingConfiguration interface{} `field:"required" json:"routingConfiguration" yaml:"routingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#concurrency_configuration BedrockFlowAlias#concurrency_configuration}.
	ConcurrencyConfiguration *BedrockFlowAliasConcurrencyConfiguration `field:"optional" json:"concurrencyConfiguration" yaml:"concurrencyConfiguration"`
	// Description of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#description BedrockFlowAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_flow_alias#tags BedrockFlowAlias#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

