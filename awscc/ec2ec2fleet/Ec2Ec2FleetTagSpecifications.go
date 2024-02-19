package ec2ec2fleet


type Ec2Ec2FleetTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#resource_type Ec2Ec2Fleet#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#tags Ec2Ec2Fleet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

