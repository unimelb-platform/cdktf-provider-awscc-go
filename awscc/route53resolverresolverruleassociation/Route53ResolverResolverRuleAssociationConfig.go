package route53resolverresolverruleassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverResolverRuleAssociationConfig struct {
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
	// The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_rule_association#resolver_rule_id Route53ResolverResolverRuleAssociation#resolver_rule_id}
	ResolverRuleId *string `field:"required" json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_rule_association#vpc_id Route53ResolverResolverRuleAssociation#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The name of an association between a Resolver rule and a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_rule_association#name Route53ResolverResolverRuleAssociation#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

