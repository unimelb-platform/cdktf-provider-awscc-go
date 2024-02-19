package ec2launchtemplate


type Ec2LaunchTemplateTagSpecifications struct {
	// The type of resource to tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#resource_type Ec2LaunchTemplate#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#tags Ec2LaunchTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

