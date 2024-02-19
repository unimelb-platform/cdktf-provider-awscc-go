package appstreamdirectoryconfig


type AppstreamDirectoryConfigCertificateBasedAuthProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#certificate_authority_arn AppstreamDirectoryConfig#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#status AppstreamDirectoryConfig#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

