package paymentcryptographykey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaymentcryptographyKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#exportable PaymentcryptographyKey#exportable}.
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#key_attributes PaymentcryptographyKey#key_attributes}.
	KeyAttributes *PaymentcryptographyKeyKeyAttributes `field:"required" json:"keyAttributes" yaml:"keyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#derive_key_usage PaymentcryptographyKey#derive_key_usage}.
	DeriveKeyUsage *string `field:"optional" json:"deriveKeyUsage" yaml:"deriveKeyUsage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#enabled PaymentcryptographyKey#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#key_check_value_algorithm PaymentcryptographyKey#key_check_value_algorithm}.
	KeyCheckValueAlgorithm *string `field:"optional" json:"keyCheckValueAlgorithm" yaml:"keyCheckValueAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/paymentcryptography_key#tags PaymentcryptographyKey#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

