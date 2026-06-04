package ec2keypair

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2KeyPairConfig struct {
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
	// A unique name for the key pair.  Constraints: Up to 255 ASCII characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#key_name Ec2KeyPair#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The format of the key pair.  Default: ``pem``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#key_format Ec2KeyPair#key_format}
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// The type of key pair.
	//
	// Note that ED25519 keys are not supported for Windows instances.
	//  If the ``PublicKeyMaterial`` property is specified, the ``KeyType`` property is ignored, and the key type is inferred from the ``PublicKeyMaterial`` value.
	//  Default: ``rsa``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#key_type Ec2KeyPair#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The public key material.
	//
	// The ``PublicKeyMaterial`` property is used to import a key pair. If this property is not specified, then a new key pair will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#public_key_material Ec2KeyPair#public_key_material}
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// The tags to apply to the key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#tags Ec2KeyPair#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

