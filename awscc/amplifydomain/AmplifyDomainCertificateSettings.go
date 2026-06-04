package amplifydomain


type AmplifyDomainCertificateSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/amplify_domain#certificate_type AmplifyDomain#certificate_type}.
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/amplify_domain#custom_certificate_arn AmplifyDomain#custom_certificate_arn}.
	CustomCertificateArn *string `field:"optional" json:"customCertificateArn" yaml:"customCertificateArn"`
}

