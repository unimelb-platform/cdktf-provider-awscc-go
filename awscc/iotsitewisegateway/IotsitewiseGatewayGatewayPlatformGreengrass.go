package iotsitewisegateway


type IotsitewiseGatewayGatewayPlatformGreengrass struct {
	// The ARN of the Greengrass group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#group_arn IotsitewiseGateway#group_arn}
	GroupArn *string `field:"required" json:"groupArn" yaml:"groupArn"`
}

