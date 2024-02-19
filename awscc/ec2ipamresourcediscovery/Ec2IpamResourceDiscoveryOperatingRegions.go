package ec2ipamresourcediscovery


type Ec2IpamResourceDiscoveryOperatingRegions struct {
	// The name of the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_resource_discovery#region_name Ec2IpamResourceDiscovery#region_name}
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

