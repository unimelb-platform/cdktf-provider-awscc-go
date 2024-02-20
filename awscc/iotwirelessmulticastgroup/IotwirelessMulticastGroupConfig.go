package iotwirelessmulticastgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessMulticastGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Multicast group LoRaWAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#lo_ra_wan IotwirelessMulticastGroup#lo_ra_wan}
	LoRaWan *IotwirelessMulticastGroupLoRaWan `field:"required" json:"loRaWan" yaml:"loRaWan"`
	// Wireless device to associate. Only for update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#associate_wireless_device IotwirelessMulticastGroup#associate_wireless_device}
	AssociateWirelessDevice *string `field:"optional" json:"associateWirelessDevice" yaml:"associateWirelessDevice"`
	// Multicast group description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#description IotwirelessMulticastGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Wireless device to disassociate. Only for update request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#disassociate_wireless_device IotwirelessMulticastGroup#disassociate_wireless_device}
	DisassociateWirelessDevice *string `field:"optional" json:"disassociateWirelessDevice" yaml:"disassociateWirelessDevice"`
	// Name of Multicast group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#name IotwirelessMulticastGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the Multicast group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#tags IotwirelessMulticastGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

