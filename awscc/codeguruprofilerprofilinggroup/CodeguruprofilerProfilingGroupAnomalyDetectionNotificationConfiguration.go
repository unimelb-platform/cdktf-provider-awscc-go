package codeguruprofilerprofilinggroup


type CodeguruprofilerProfilingGroupAnomalyDetectionNotificationConfiguration struct {
	// Unique arn of the resource to be used for notifications.
	//
	// We support a valid SNS topic arn as a channel uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#channel_uri CodeguruprofilerProfilingGroup#channel_uri}
	ChannelUri *string `field:"required" json:"channelUri" yaml:"channelUri"`
	// Unique identifier for each Channel in the notification configuration of a Profiling Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#channel_id CodeguruprofilerProfilingGroup#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
}

