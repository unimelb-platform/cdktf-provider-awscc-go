package apigatewayv2routingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2RoutingRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#actions Apigatewayv2RoutingRule#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#conditions Apigatewayv2RoutingRule#conditions}.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The amazon resource name (ARN) of the domain name resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#domain_name_arn Apigatewayv2RoutingRule#domain_name_arn}
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#priority Apigatewayv2RoutingRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

