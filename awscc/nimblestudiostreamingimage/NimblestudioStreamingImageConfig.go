package nimblestudiostreamingimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NimblestudioStreamingImageConfig struct {
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
	// <p>The ID of an EC2 machine image with which to create this streaming image.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_streaming_image#ec_2_image_id NimblestudioStreamingImage#ec_2_image_id}
	Ec2ImageId *string `field:"required" json:"ec2ImageId" yaml:"ec2ImageId"`
	// <p>A friendly name for a streaming image resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_streaming_image#name NimblestudioStreamingImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The studioId. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_streaming_image#studio_id NimblestudioStreamingImage#studio_id}
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// <p>A human-readable description of the streaming image.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_streaming_image#description NimblestudioStreamingImage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_streaming_image#tags NimblestudioStreamingImage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

