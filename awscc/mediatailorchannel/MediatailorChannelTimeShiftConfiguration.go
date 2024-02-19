package mediatailorchannel


type MediatailorChannelTimeShiftConfiguration struct {
	// <p>The maximum time delay for time-shifted viewing.
	//
	// The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel#max_time_delay_seconds MediatailorChannel#max_time_delay_seconds}
	MaxTimeDelaySeconds *float64 `field:"required" json:"maxTimeDelaySeconds" yaml:"maxTimeDelaySeconds"`
}

