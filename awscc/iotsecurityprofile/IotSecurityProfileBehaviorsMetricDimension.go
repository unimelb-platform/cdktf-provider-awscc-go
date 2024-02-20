package iotsecurityprofile


type IotSecurityProfileBehaviorsMetricDimension struct {
	// A unique identifier for the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#dimension_name IotSecurityProfile#dimension_name}
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Defines how the dimensionValues of a dimension are interpreted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#operator IotSecurityProfile#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

