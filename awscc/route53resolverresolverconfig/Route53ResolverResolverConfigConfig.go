package route53resolverresolverconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverResolverConfigConfig struct {
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
	// Represents the desired status of AutodefinedReverse.
	//
	// The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_config#autodefined_reverse_flag Route53ResolverResolverConfig#autodefined_reverse_flag}
	AutodefinedReverseFlag *string `field:"required" json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_config#resource_id Route53ResolverResolverConfig#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

