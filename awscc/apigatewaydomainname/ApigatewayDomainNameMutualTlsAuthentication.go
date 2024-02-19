package apigatewaydomainname


type ApigatewayDomainNameMutualTlsAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_domain_name#truststore_uri ApigatewayDomainName#truststore_uri}.
	TruststoreUri *string `field:"optional" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_domain_name#truststore_version ApigatewayDomainName#truststore_version}.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

