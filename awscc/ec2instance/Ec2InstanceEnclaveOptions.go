package ec2instance


type Ec2InstanceEnclaveOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#enabled Ec2Instance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

