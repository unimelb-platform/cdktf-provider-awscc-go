package kinesisstreamconsumer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisStreamConsumerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Kinesis Stream Consumer.
	//
	// For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesis_stream_consumer#consumer_name KinesisStreamConsumer#consumer_name}
	ConsumerName *string `field:"required" json:"consumerName" yaml:"consumerName"`
	// The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesis_stream_consumer#stream_arn KinesisStreamConsumer#stream_arn}
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesis_stream_consumer#tags KinesisStreamConsumer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

