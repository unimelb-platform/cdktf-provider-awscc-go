package appstreamdirectoryconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppstreamDirectoryConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#directory_name AppstreamDirectoryConfig#directory_name}.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#organizational_unit_distinguished_names AppstreamDirectoryConfig#organizational_unit_distinguished_names}.
	OrganizationalUnitDistinguishedNames *[]*string `field:"required" json:"organizationalUnitDistinguishedNames" yaml:"organizationalUnitDistinguishedNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#service_account_credentials AppstreamDirectoryConfig#service_account_credentials}.
	ServiceAccountCredentials *AppstreamDirectoryConfigServiceAccountCredentials `field:"required" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_directory_config#certificate_based_auth_properties AppstreamDirectoryConfig#certificate_based_auth_properties}.
	CertificateBasedAuthProperties *AppstreamDirectoryConfigCertificateBasedAuthProperties `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
}

