package syntheticscanary


type SyntheticsCanaryScheduleRetryConfig struct {
	// maximum times the canary will be retried upon the scheduled run failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/synthetics_canary#max_retries SyntheticsCanary#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

