package snstopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnsTopicConfig struct {
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
	// The archive policy determines the number of days SNS retains messages.
	//
	// You can set a retention period from 1 to 365 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#archive_policy SnsTopic#archive_policy}
	ArchivePolicy *string `field:"optional" json:"archivePolicy" yaml:"archivePolicy"`
	// Enables content-based deduplication for FIFO topics.
	//
	// +  By default, ``ContentBasedDeduplication`` is set to ``false``. If you create a FIFO topic and this attribute is ``false``, you must specify a value for the ``MessageDeduplicationId`` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	//   +  When you set ``ContentBasedDeduplication`` to ``true``, SNS uses a SHA-256 hash to generate the ``MessageDeduplicationId`` using the body of the message (but not the attributes of the message).
	//  (Optional) To override the generated value, you can specify a value for the the ``MessageDeduplicationId`` parameter for the ``Publish`` action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#content_based_deduplication SnsTopic#content_based_deduplication}
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//  The policy must be in JSON string format.
	//  Length Constraints: Maximum length of 30,720.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#data_protection_policy SnsTopic#data_protection_policy}
	DataProtectionPolicy *string `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The ``DeliveryStatusLogging`` configuration enables you to log the delivery status of messages sent from your Amazon SNS topic to subscribed endpoints with the following supported delivery protocols:   +  HTTP    +  Amazon Kinesis Data Firehose   +   AWS Lambda   +  Platform application endpoint   +  Amazon Simple Queue Service     Once configured, log entries are sent to Amazon CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#delivery_status_logging SnsTopic#delivery_status_logging}
	DeliveryStatusLogging interface{} `field:"optional" json:"deliveryStatusLogging" yaml:"deliveryStatusLogging"`
	// The display name to use for an SNS topic with SMS subscriptions.
	//
	// The display name must be maximum 100 characters long, including hyphens (-), underscores (_), spaces, and tabs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#display_name SnsTopic#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#fifo_throughput_scope SnsTopic#fifo_throughput_scope}.
	FifoThroughputScope *string `field:"optional" json:"fifoThroughputScope" yaml:"fifoThroughputScope"`
	// Set to true to create a FIFO topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#fifo_topic SnsTopic#fifo_topic}
	FifoTopic interface{} `field:"optional" json:"fifoTopic" yaml:"fifoTopic"`
	// The ID of an AWS managed customer master key (CMK) for SNS or a custom CMK.
	//
	// For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms). For more examples, see ``KeyId`` in the *API Reference*.
	//  This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#kms_master_key_id SnsTopic#kms_master_key_id}
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	//
	// By default, ``SignatureVersion`` is set to ``1``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#signature_version SnsTopic#signature_version}
	SignatureVersion *string `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// The SNS subscriptions (endpoints) for this topic.
	//
	// If you specify the ``Subscription`` property in the ``AWS::SNS::Topic`` resource and it creates an associated subscription resource, the associated subscription is not deleted when the ``AWS::SNS::Topic`` resource is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#subscription SnsTopic#subscription}
	Subscription interface{} `field:"optional" json:"subscription" yaml:"subscription"`
	// The list of tags to add to a new topic.
	//
	// To be able to tag a topic on creation, you must have the ``sns:CreateTopic`` and ``sns:TagResource`` permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#tags SnsTopic#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the topic you want to create.
	//
	// Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with ``.fifo``.
	//  If you don't specify a name, CFN generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#topic_name SnsTopic#topic_name}
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
	// Tracing mode of an SNS topic.
	//
	// By default ``TracingConfig`` is set to ``PassThrough``, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to ``Active``, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic#tracing_config SnsTopic#tracing_config}
	TracingConfig *string `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
}

