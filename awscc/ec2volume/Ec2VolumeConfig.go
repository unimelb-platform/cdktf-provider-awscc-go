package ec2volume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VolumeConfig struct {
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
	// The Availability Zone in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#availability_zone Ec2Volume#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The Availability Zone in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#auto_enable_io Ec2Volume#auto_enable_io}
	AutoEnableIo interface{} `field:"optional" json:"autoEnableIo" yaml:"autoEnableIo"`
	// Specifies whether the volume should be encrypted.
	//
	// The effect of setting the encryption state to true depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see Encryption by default in the Amazon Elastic Compute Cloud User Guide. Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see Supported instance types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#encrypted Ec2Volume#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS) to provision for an io1 or io2 volume, with a maximum ratio of 50 IOPS/GiB for io1, and 500 IOPS/GiB for io2.
	//
	// Range is 100 to 64,000 IOPS for volumes in most Regions. Maximum IOPS of 64,000 is guaranteed only on Nitro-based instances. Other instance families guarantee performance up to 32,000 IOPS. For more information, see Amazon EBS volume types in the Amazon Elastic Compute Cloud User Guide. This parameter is valid only for Provisioned IOPS SSD (io1 and io2) volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#iops Ec2Volume#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption.
	//
	// If KmsKeyId is specified, the encrypted state must be true. If you omit this property and your account is enabled for encryption by default, or Encrypted is set to true, then the volume is encrypted using the default CMK specified for your account. If your account does not have a default CMK, then the volume is encrypted using the AWS managed CMK.  Alternatively, if you want to specify a different CMK, you can specify one of the following:  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab. Key alias. Specify the alias for the CMK, prefixed with alias/. For example, for a CMK with the alias my_cmk, use alias/my_cmk. Or to specify the AWS managed CMK, use alias/aws/ebs. Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab. Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#kms_key_id Ec2Volume#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Indicates whether Amazon EBS Multi-Attach is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#multi_attach_enabled Ec2Volume#multi_attach_enabled}
	MultiAttachEnabled interface{} `field:"optional" json:"multiAttachEnabled" yaml:"multiAttachEnabled"`
	// The Amazon Resource Name (ARN) of the Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#outpost_arn Ec2Volume#outpost_arn}
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size.  Constraints: 1-16,384 for gp2, 4-16,384 for io1 and io2, 500-16,384 for st1, 500-16,384 for sc1, and 1-1,024 for standard. If you specify a snapshot, the volume size must be equal to or larger than the snapshot size. Default: If you're creating the volume from a snapshot and don't specify a volume size, the default is the snapshot size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#size Ec2Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which to create the volume.
	//
	// You must specify either a snapshot ID or a volume size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#snapshot_id Ec2Volume#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The tags to apply to the volume during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#tags Ec2Volume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The throughput that the volume supports, in MiB/s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#throughput Ec2Volume#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume type.
	//
	// This parameter can be one of the following values: General Purpose SSD: gp2 | gp3, Provisioned IOPS SSD: io1 | io2, Throughput Optimized HDD: st1, Cold HDD: sc1, Magnetic: standard
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_volume#volume_type Ec2Volume#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

