package ec2instance


type Ec2InstanceHibernationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#configured Ec2Instance#configured}.
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

