package route53resolveroutpostresolver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverOutpostResolverConfig struct {
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
	// The OutpostResolver name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_outpost_resolver#name Route53ResolverOutpostResolver#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Outpost ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_outpost_resolver#outpost_arn Route53ResolverOutpostResolver#outpost_arn}
	OutpostArn *string `field:"required" json:"outpostArn" yaml:"outpostArn"`
	// The OutpostResolver instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_outpost_resolver#preferred_instance_type Route53ResolverOutpostResolver#preferred_instance_type}
	PreferredInstanceType *string `field:"required" json:"preferredInstanceType" yaml:"preferredInstanceType"`
	// The number of OutpostResolvers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_outpost_resolver#instance_count Route53ResolverOutpostResolver#instance_count}
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_outpost_resolver#tags Route53ResolverOutpostResolver#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

