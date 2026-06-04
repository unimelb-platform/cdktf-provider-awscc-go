package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#availability_zone RedshiftserverlessWorkgroup#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#network_interface_id RedshiftserverlessWorkgroup#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#private_ip_address RedshiftserverlessWorkgroup#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#subnet_id RedshiftserverlessWorkgroup#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

