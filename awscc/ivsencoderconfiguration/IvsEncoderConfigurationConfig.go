package ivsencoderconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsEncoderConfigurationConfig struct {
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
	// Encoder configuration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_encoder_configuration#name IvsEncoderConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_encoder_configuration#tags IvsEncoderConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_encoder_configuration#video IvsEncoderConfiguration#video}
	Video *IvsEncoderConfigurationVideo `field:"optional" json:"video" yaml:"video"`
}

