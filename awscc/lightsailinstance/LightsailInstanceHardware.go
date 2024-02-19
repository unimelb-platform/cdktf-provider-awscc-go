package lightsailinstance


type LightsailInstanceHardware struct {
	// Disks attached to the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_instance#disks LightsailInstance#disks}
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
}

