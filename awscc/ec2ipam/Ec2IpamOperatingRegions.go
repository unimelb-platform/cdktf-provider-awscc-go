package ec2ipam


type Ec2IpamOperatingRegions struct {
	// The name of the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam#region_name Ec2Ipam#region_name}
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

