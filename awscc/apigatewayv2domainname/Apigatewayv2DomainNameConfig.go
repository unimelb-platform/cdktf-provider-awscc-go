package apigatewayv2domainname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2DomainNameConfig struct {
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
	// The custom domain name for your API in Amazon API Gateway. Uppercase letters are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#domain_name Apigatewayv2DomainName#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The domain name configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#domain_name_configurations Apigatewayv2DomainName#domain_name_configurations}
	DomainNameConfigurations interface{} `field:"optional" json:"domainNameConfigurations" yaml:"domainNameConfigurations"`
	// The mutual TLS authentication configuration for a custom domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#mutual_tls_authentication Apigatewayv2DomainName#mutual_tls_authentication}
	MutualTlsAuthentication *Apigatewayv2DomainNameMutualTlsAuthentication `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The collection of tags associated with a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#tags Apigatewayv2DomainName#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

