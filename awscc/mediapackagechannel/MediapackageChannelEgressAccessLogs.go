package mediapackagechannel


type MediapackageChannelEgressAccessLogs struct {
	// Sets a custom AWS CloudWatch log group name for access logs.
	//
	// If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#log_group_name MediapackageChannel#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

