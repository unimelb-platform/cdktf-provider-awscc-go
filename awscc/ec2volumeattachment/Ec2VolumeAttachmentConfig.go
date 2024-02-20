package ec2volumeattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VolumeAttachmentConfig struct {
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
	// The ID of the instance to which the volume attaches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume_attachment#instance_id Ec2VolumeAttachment#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ID of the Amazon EBS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume_attachment#volume_id Ec2VolumeAttachment#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// The device name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume_attachment#device Ec2VolumeAttachment#device}
	Device *string `field:"optional" json:"device" yaml:"device"`
}

