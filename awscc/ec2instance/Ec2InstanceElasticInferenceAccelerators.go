package ec2instance


type Ec2InstanceElasticInferenceAccelerators struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#type Ec2Instance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#count Ec2Instance#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

