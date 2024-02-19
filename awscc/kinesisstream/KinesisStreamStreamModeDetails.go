package kinesisstream


type KinesisStreamStreamModeDetails struct {
	// The mode of the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#stream_mode KinesisStream#stream_mode}
	StreamMode *string `field:"required" json:"streamMode" yaml:"streamMode"`
}

