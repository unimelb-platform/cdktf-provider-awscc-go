package ec2instance


type Ec2InstanceElasticGpuSpecifications struct {
	// The type of Elastic Graphics accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_instance#type Ec2Instance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

