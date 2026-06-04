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
	// The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#availability_zone Ec2Volume#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Indicates whether the volume is auto-enabled for I/O operations.
	//
	// By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#auto_enable_io Ec2Volume#auto_enable_io}
	AutoEnableIo interface{} `field:"optional" json:"autoEnableIo" yaml:"autoEnableIo"`
	// Indicates whether the volume should be encrypted.
	//
	// The effect of setting the encryption state to ``true`` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*.
	//  Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#encrypted Ec2Volume#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//  The following are the supported values for each volume type:
	//   +  ``gp3``: 3,000 - 16,000 IOPS
	//   +  ``io1``: 100 - 64,000 IOPS
	//   +  ``io2``: 100 - 256,000 IOPS
	//
	//  For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html). On other instances, you can achieve performance up to 32,000 IOPS.
	//  This parameter is required for ``io1`` and ``io2`` volumes. The default for ``gp3`` volumes is 3,000 IOPS. This parameter is not supported for ``gp2``, ``st1``, ``sc1``, or ``standard`` volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#iops Ec2Volume#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The identifier of the kms-key-long to use for Amazon EBS encryption.
	//
	// If ``KmsKeyId`` is specified, the encrypted state must be ``true``.
	//  If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.
	//  Alternatively, if you want to specify a different key, you can specify one of the following:
	//   +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//   +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.
	//   +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//   +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#kms_key_id Ec2Volume#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Indicates whether Amazon EBS Multi-Attach is enabled.
	//
	// CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#multi_attach_enabled Ec2Volume#multi_attach_enabled}
	MultiAttachEnabled interface{} `field:"optional" json:"multiAttachEnabled" yaml:"multiAttachEnabled"`
	// The Amazon Resource Name (ARN) of the Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#outpost_arn Ec2Volume#outpost_arn}
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.
	//  The following are the supported volumes sizes for each volume type:
	//   +  ``gp2`` and ``gp3``: 1 - 16,384 GiB
	//   +  ``io1``: 4 - 16,384 GiB
	//   +  ``io2``: 4 - 65,536 GiB
	//   +  ``st1`` and ``sc1``: 125 - 16,384 GiB
	//   +  ``standard``: 1 - 1024 GiB
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#size Ec2Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#snapshot_id Ec2Volume#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The tags to apply to the volume during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#tags Ec2Volume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The throughput to provision for a volume, with a maximum of 1,000 MiB/s.
	//
	// This parameter is valid only for ``gp3`` volumes. The default value is 125.
	//  Valid Range: Minimum value of 125. Maximum value of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#throughput Ec2Volume#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#volume_initialization_rate Ec2Volume#volume_initialization_rate}.
	VolumeInitializationRate *float64 `field:"optional" json:"volumeInitializationRate" yaml:"volumeInitializationRate"`
	// The volume type.
	//
	// This parameter can be one of the following values:
	//   +  General Purpose SSD: ``gp2`` | ``gp3``
	//   +  Provisioned IOPS SSD: ``io1`` | ``io2``
	//   +  Throughput Optimized HDD: ``st1``
	//   +  Cold HDD: ``sc1``
	//   +  Magnetic: ``standard``
	//
	//  For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html).
	//  Default: ``gp2``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_volume#volume_type Ec2Volume#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

