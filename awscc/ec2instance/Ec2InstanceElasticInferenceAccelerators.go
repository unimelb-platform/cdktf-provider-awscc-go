package ec2instance


type Ec2InstanceElasticInferenceAccelerators struct {
	// The number of elastic inference accelerators to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_instance#count Ec2Instance#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The type of elastic inference accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_instance#type Ec2Instance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

