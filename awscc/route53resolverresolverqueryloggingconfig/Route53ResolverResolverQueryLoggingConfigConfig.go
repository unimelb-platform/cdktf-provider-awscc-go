package route53resolverresolverqueryloggingconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverResolverQueryLoggingConfigConfig struct {
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
	// destination arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_query_logging_config#destination_arn Route53ResolverResolverQueryLoggingConfig#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// ResolverQueryLogConfigName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_resolver_query_logging_config#name Route53ResolverResolverQueryLoggingConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

