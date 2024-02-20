package gameliftalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameliftAliasConfig struct {
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
	// A descriptive label that is associated with an alias. Alias names do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_alias#name GameliftAlias#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_alias#routing_strategy GameliftAlias#routing_strategy}
	RoutingStrategy *GameliftAliasRoutingStrategy `field:"required" json:"routingStrategy" yaml:"routingStrategy"`
	// A human-readable description of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_alias#description GameliftAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

