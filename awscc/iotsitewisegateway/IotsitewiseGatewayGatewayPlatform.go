package iotsitewisegateway


type IotsitewiseGatewayGatewayPlatform struct {
	// A gateway that runs on AWS IoT Greengrass V1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#greengrass IotsitewiseGateway#greengrass}
	Greengrass *IotsitewiseGatewayGatewayPlatformGreengrass `field:"optional" json:"greengrass" yaml:"greengrass"`
	// A gateway that runs on AWS IoT Greengrass V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#greengrass_v2 IotsitewiseGateway#greengrass_v2}
	GreengrassV2 *IotsitewiseGatewayGatewayPlatformGreengrassV2 `field:"optional" json:"greengrassV2" yaml:"greengrassV2"`
}

