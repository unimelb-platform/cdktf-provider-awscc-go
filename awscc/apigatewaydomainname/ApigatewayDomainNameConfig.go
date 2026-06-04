package apigatewaydomainname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayDomainNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#certificate_arn ApigatewayDomainName#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#domain_name ApigatewayDomainName#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The ``EndpointConfiguration`` property type specifies the endpoint types of an Amazon API Gateway domain name.
	//
	// ``EndpointConfiguration`` is a property of the [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#endpoint_configuration ApigatewayDomainName#endpoint_configuration}
	EndpointConfiguration *ApigatewayDomainNameEndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#mutual_tls_authentication ApigatewayDomainName#mutual_tls_authentication}.
	MutualTlsAuthentication *ApigatewayDomainNameMutualTlsAuthentication `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The ARN of the public certificate issued by ACM to validate ownership of your custom domain.
	//
	// Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#ownership_verification_certificate_arn ApigatewayDomainName#ownership_verification_certificate_arn}
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#regional_certificate_arn ApigatewayDomainName#regional_certificate_arn}.
	RegionalCertificateArn *string `field:"optional" json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#routing_mode ApigatewayDomainName#routing_mode}.
	RoutingMode *string `field:"optional" json:"routingMode" yaml:"routingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#security_policy ApigatewayDomainName#security_policy}.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_domain_name#tags ApigatewayDomainName#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

