package kinesisvideosignalingchannel


type KinesisvideoSignalingChannelTags struct {
	// The key name of the tag.
	//
	// Specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#key KinesisvideoSignalingChannel#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.  The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#value KinesisvideoSignalingChannel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

