package iotpolicy


type IotPolicyTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_policy#key IotPolicy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_policy#value IotPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

