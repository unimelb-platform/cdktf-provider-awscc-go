package ec2instance


type Ec2InstanceVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#device Ec2Instance#device}.
	Device *string `field:"required" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#volume_id Ec2Instance#volume_id}.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

