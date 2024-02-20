package s3bucket


type S3BucketNotificationConfigurationEventBridgeConfiguration struct {
	// Specifies whether to send notifications to Amazon EventBridge when events occur in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event_bridge_enabled S3Bucket#event_bridge_enabled}
	EventBridgeEnabled interface{} `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
}

