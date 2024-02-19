package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteriaMlDetectionConfig struct {
	// The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#confidence_level IotSecurityProfile#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

