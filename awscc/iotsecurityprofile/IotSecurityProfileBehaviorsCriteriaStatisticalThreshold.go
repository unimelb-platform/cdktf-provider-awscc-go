package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteriaStatisticalThreshold struct {
	// The percentile which resolves to a threshold value by which compliance with a behavior is determined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#statistic IotSecurityProfile#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

