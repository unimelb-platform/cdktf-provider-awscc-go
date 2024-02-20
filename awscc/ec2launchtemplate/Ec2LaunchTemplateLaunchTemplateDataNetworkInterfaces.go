package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfaces struct {
	// Indicates whether to associate a Carrier IP address with eth0 for a new network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#associate_carrier_ip_address Ec2LaunchTemplate#associate_carrier_ip_address}
	AssociateCarrierIpAddress interface{} `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Associates a public IPv4 address with eth0 for a new network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#associate_public_ip_address Ec2LaunchTemplate#associate_public_ip_address}
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Allows customer to specify Connection Tracking options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#connection_tracking_specification Ec2LaunchTemplate#connection_tracking_specification}
	ConnectionTrackingSpecification *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification `field:"optional" json:"connectionTrackingSpecification" yaml:"connectionTrackingSpecification"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#delete_on_termination Ec2LaunchTemplate#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// A description for the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#description Ec2LaunchTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The device index for the network interface attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#device_index Ec2LaunchTemplate#device_index}
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Allows customer to specify ENA-SRD options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ena_srd_specification Ec2LaunchTemplate#ena_srd_specification}
	EnaSrdSpecification *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
	// The IDs of one or more security groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#groups Ec2LaunchTemplate#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The type of network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#interface_type Ec2LaunchTemplate#interface_type}
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv4 prefixes to be automatically assigned to the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_4_prefix_count Ec2LaunchTemplate#ipv_4_prefix_count}
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// One or more IPv4 prefixes to be assigned to the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_4_prefixes Ec2LaunchTemplate#ipv_4_prefixes}
	Ipv4Prefixes interface{} `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_6_address_count Ec2LaunchTemplate#ipv_6_address_count}
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_6_addresses Ec2LaunchTemplate#ipv_6_addresses}
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The number of IPv6 prefixes to be automatically assigned to the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_6_prefix_count Ec2LaunchTemplate#ipv_6_prefix_count}
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// One or more IPv6 prefixes to be assigned to the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ipv_6_prefixes Ec2LaunchTemplate#ipv_6_prefixes}
	Ipv6Prefixes interface{} `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// The index of the network card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#network_card_index Ec2LaunchTemplate#network_card_index}
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// The ID of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#network_interface_id Ec2LaunchTemplate#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Enables the first IPv6 global unique address (GUA) on a dual stack or IPv6-only ENI immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#primary_ipv_6 Ec2LaunchTemplate#primary_ipv_6}
	PrimaryIpv6 interface{} `field:"optional" json:"primaryIpv6" yaml:"primaryIpv6"`
	// The primary private IPv4 address of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#private_ip_address Ec2LaunchTemplate#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// One or more private IPv4 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#private_ip_addresses Ec2LaunchTemplate#private_ip_addresses}
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#secondary_private_ip_address_count Ec2LaunchTemplate#secondary_private_ip_address_count}
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet for the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#subnet_id Ec2LaunchTemplate#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

