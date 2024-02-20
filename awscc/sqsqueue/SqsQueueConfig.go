package sqsqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqsQueueConfig struct {
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
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication.
	//
	// During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#content_based_deduplication SqsQueue#content_based_deduplication}
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#deduplication_scope SqsQueue#deduplication_scope}
	DeduplicationScope *string `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// The time in seconds for which the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#delay_seconds SqsQueue#delay_seconds}
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// If set to true, creates a FIFO queue.
	//
	// If you don't specify this property, Amazon SQS creates a standard queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#fifo_queue SqsQueue#fifo_queue}
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#fifo_throughput_limit SqsQueue#fifo_throughput_limit}
	FifoThroughputLimit *string `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#kms_data_key_reuse_period_seconds SqsQueue#kms_data_key_reuse_period_seconds}
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK.
	//
	// To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#kms_master_key_id SqsQueue#kms_master_key_id}
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#maximum_message_size SqsQueue#maximum_message_size}
	MaximumMessageSize *float64 `field:"optional" json:"maximumMessageSize" yaml:"maximumMessageSize"`
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#message_retention_period SqsQueue#message_retention_period}
	MessageRetentionPeriod *float64 `field:"optional" json:"messageRetentionPeriod" yaml:"messageRetentionPeriod"`
	// A name for the queue.
	//
	// To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#queue_name SqsQueue#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available.
	//
	// You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#receive_message_wait_time_seconds SqsQueue#receive_message_wait_time_seconds}
	ReceiveMessageWaitTimeSeconds *float64 `field:"optional" json:"receiveMessageWaitTimeSeconds" yaml:"receiveMessageWaitTimeSeconds"`
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#redrive_allow_policy SqsQueue#redrive_allow_policy}
	RedriveAllowPolicy *string `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#redrive_policy SqsQueue#redrive_policy}
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Enables server-side queue encryption using SQS owned encryption keys.
	//
	// Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#sqs_managed_sse_enabled SqsQueue#sqs_managed_sse_enabled}
	SqsManagedSseEnabled interface{} `field:"optional" json:"sqsManagedSseEnabled" yaml:"sqsManagedSseEnabled"`
	// The tags that you attach to this queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#tags SqsQueue#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The length of time during which a message will be unavailable after a message is delivered from the queue.
	//
	// This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sqs_queue#visibility_timeout SqsQueue#visibility_timeout}
	VisibilityTimeout *float64 `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

