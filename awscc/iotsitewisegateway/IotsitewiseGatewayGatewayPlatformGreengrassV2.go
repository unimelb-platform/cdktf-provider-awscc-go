package iotsitewisegateway


type IotsitewiseGatewayGatewayPlatformGreengrassV2 struct {
	// The name of the CoreDevice in GreenGrass V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#core_device_thing_name IotsitewiseGateway#core_device_thing_name}
	CoreDeviceThingName *string `field:"required" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

