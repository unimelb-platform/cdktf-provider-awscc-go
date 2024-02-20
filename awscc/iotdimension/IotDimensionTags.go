package iotdimension


type IotDimensionTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_dimension#key IotDimension#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_dimension#value IotDimension#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

