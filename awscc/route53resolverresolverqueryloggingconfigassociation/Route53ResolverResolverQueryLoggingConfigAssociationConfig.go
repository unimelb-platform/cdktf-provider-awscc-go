package route53resolverresolverqueryloggingconfigassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverResolverQueryLoggingConfigAssociationConfig struct {
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
	// ResolverQueryLogConfigId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_query_logging_config_association#resolver_query_log_config_id Route53ResolverResolverQueryLoggingConfigAssociation#resolver_query_log_config_id}
	ResolverQueryLogConfigId *string `field:"optional" json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_query_logging_config_association#resource_id Route53ResolverResolverQueryLoggingConfigAssociation#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

