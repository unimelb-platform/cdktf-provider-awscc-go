package ec2ipampool


type Ec2IpamPoolSourceResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool#resource_id Ec2IpamPool#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool#resource_owner Ec2IpamPool#resource_owner}.
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool#resource_region Ec2IpamPool#resource_region}.
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool#resource_type Ec2IpamPool#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

