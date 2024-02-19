package apigatewayv2domainname


type Apigatewayv2DomainNameDomainNameConfigurations struct {
	// An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#certificate_arn Apigatewayv2DomainName#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#certificate_name Apigatewayv2DomainName#certificate_name}
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The endpoint type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#endpoint_type Apigatewayv2DomainName#endpoint_type}
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The Amazon resource name (ARN) for the public certificate issued by ACMlong.
	//
	// This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#ownership_verification_certificate_arn Apigatewayv2DomainName#ownership_verification_certificate_arn}
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// The Transport Layer Security (TLS) version of the security policy for this domain name.
	//
	// The valid values are ``TLS_1_0`` and ``TLS_1_2``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_domain_name#security_policy Apigatewayv2DomainName#security_policy}
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

