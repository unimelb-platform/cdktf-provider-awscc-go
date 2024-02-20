package lightsailcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LightsailCertificateConfig struct {
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
	// The name for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_certificate#certificate_name LightsailCertificate#certificate_name}
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The domain name (e.g., example.com ) for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_certificate#domain_name LightsailCertificate#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_certificate#subject_alternative_names LightsailCertificate#subject_alternative_names}
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_certificate#tags LightsailCertificate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

