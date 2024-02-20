package iotcertificateprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotCertificateProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_certificate_provider#account_default_for_operations IotCertificateProvider#account_default_for_operations}.
	AccountDefaultForOperations *[]*string `field:"required" json:"accountDefaultForOperations" yaml:"accountDefaultForOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_certificate_provider#lambda_function_arn IotCertificateProvider#lambda_function_arn}.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_certificate_provider#certificate_provider_name IotCertificateProvider#certificate_provider_name}.
	CertificateProviderName *string `field:"optional" json:"certificateProviderName" yaml:"certificateProviderName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_certificate_provider#tags IotCertificateProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

