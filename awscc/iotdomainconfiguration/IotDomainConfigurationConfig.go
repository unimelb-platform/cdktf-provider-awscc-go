package iotdomainconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotDomainConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#authorizer_config IotDomainConfiguration#authorizer_config}.
	AuthorizerConfig *IotDomainConfigurationAuthorizerConfig `field:"optional" json:"authorizerConfig" yaml:"authorizerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#domain_configuration_name IotDomainConfiguration#domain_configuration_name}.
	DomainConfigurationName *string `field:"optional" json:"domainConfigurationName" yaml:"domainConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#domain_configuration_status IotDomainConfiguration#domain_configuration_status}.
	DomainConfigurationStatus *string `field:"optional" json:"domainConfigurationStatus" yaml:"domainConfigurationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#domain_name IotDomainConfiguration#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#server_certificate_arns IotDomainConfiguration#server_certificate_arns}.
	ServerCertificateArns *[]*string `field:"optional" json:"serverCertificateArns" yaml:"serverCertificateArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#service_type IotDomainConfiguration#service_type}.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#tags IotDomainConfiguration#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#tls_config IotDomainConfiguration#tls_config}.
	TlsConfig *IotDomainConfigurationTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_domain_configuration#validation_certificate_arn IotDomainConfiguration#validation_certificate_arn}.
	ValidationCertificateArn *string `field:"optional" json:"validationCertificateArn" yaml:"validationCertificateArn"`
}

