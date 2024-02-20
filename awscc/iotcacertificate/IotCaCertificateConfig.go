package iotcacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotCaCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#ca_certificate_pem IotCaCertificate#ca_certificate_pem}.
	CaCertificatePem *string `field:"required" json:"caCertificatePem" yaml:"caCertificatePem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#status IotCaCertificate#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#auto_registration_status IotCaCertificate#auto_registration_status}.
	AutoRegistrationStatus *string `field:"optional" json:"autoRegistrationStatus" yaml:"autoRegistrationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#certificate_mode IotCaCertificate#certificate_mode}.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#registration_config IotCaCertificate#registration_config}.
	RegistrationConfig *IotCaCertificateRegistrationConfig `field:"optional" json:"registrationConfig" yaml:"registrationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#remove_auto_registration IotCaCertificate#remove_auto_registration}.
	RemoveAutoRegistration interface{} `field:"optional" json:"removeAutoRegistration" yaml:"removeAutoRegistration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#tags IotCaCertificate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The private key verification certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate#verification_certificate_pem IotCaCertificate#verification_certificate_pem}
	VerificationCertificatePem *string `field:"optional" json:"verificationCertificatePem" yaml:"verificationCertificatePem"`
}

