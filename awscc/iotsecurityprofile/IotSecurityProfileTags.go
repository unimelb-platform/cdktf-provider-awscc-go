package iotsecurityprofile


type IotSecurityProfileTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_security_profile#key IotSecurityProfile#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_security_profile#value IotSecurityProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

