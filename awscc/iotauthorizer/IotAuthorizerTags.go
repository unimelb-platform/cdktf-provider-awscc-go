package iotauthorizer


type IotAuthorizerTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_authorizer#key IotAuthorizer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_authorizer#value IotAuthorizer#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

