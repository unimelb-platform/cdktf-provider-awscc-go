package ivsplaybackkeypair

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsPlaybackKeyPairConfig struct {
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
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource.
	//
	// The value does not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_playback_key_pair#name IvsPlaybackKeyPair#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The public portion of a customer-generated key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_playback_key_pair#public_key_material IvsPlaybackKeyPair#public_key_material}
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_playback_key_pair#tags IvsPlaybackKeyPair#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

