package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs struct {
	// Indicates whether the EBS volume is deleted on instance termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#delete_on_termination Ec2LaunchTemplate#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Indicates whether the EBS volume is encrypted.
	//
	// Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#encrypted Ec2LaunchTemplate#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#iops Ec2LaunchTemplate#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The ARN of the symmetric AWS Key Management Service (AWS KMS) CMK used for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#kms_key_id Ec2LaunchTemplate#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ID of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#snapshot_id Ec2LaunchTemplate#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput to provision for a gp3 volume, with a maximum of 1,000 MiB/s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#throughput Ec2LaunchTemplate#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#volume_size Ec2LaunchTemplate#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#volume_type Ec2LaunchTemplate#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

