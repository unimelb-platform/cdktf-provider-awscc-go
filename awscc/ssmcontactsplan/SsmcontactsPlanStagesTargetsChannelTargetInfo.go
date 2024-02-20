package ssmcontactsplan


type SsmcontactsPlanStagesTargetsChannelTargetInfo struct {
	// The Amazon Resource Name (ARN) of the contact channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#channel_id SsmcontactsPlan#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The number of minutes to wait to retry sending engagement in the case the engagement initially fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#retry_interval_in_minutes SsmcontactsPlan#retry_interval_in_minutes}
	RetryIntervalInMinutes *float64 `field:"required" json:"retryIntervalInMinutes" yaml:"retryIntervalInMinutes"`
}

