package iamsamlprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamSamlProviderConfig struct {
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
	// The private key from your external identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#add_private_key IamSamlProvider#add_private_key}
	AddPrivateKey *string `field:"optional" json:"addPrivateKey" yaml:"addPrivateKey"`
	// The encryption setting for the SAML provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#assertion_encryption_mode IamSamlProvider#assertion_encryption_mode}
	AssertionEncryptionMode *string `field:"optional" json:"assertionEncryptionMode" yaml:"assertionEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#name IamSamlProvider#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#private_key_list IamSamlProvider#private_key_list}.
	PrivateKeyList interface{} `field:"optional" json:"privateKeyList" yaml:"privateKeyList"`
	// The Key ID of the private key to remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#remove_private_key IamSamlProvider#remove_private_key}
	RemovePrivateKey *string `field:"optional" json:"removePrivateKey" yaml:"removePrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#saml_metadata_document IamSamlProvider#saml_metadata_document}.
	SamlMetadataDocument *string `field:"optional" json:"samlMetadataDocument" yaml:"samlMetadataDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_saml_provider#tags IamSamlProvider#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

