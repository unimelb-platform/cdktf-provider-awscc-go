package ivschatroom

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvschatRoomConfig struct {
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
	// Array of logging configuration identifiers attached to the room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#logging_configuration_identifiers IvschatRoom#logging_configuration_identifiers}
	LoggingConfigurationIdentifiers *[]*string `field:"optional" json:"loggingConfigurationIdentifiers" yaml:"loggingConfigurationIdentifiers"`
	// The maximum number of characters in a single message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#maximum_message_length IvschatRoom#maximum_message_length}
	MaximumMessageLength *float64 `field:"optional" json:"maximumMessageLength" yaml:"maximumMessageLength"`
	// The maximum number of messages per second that can be sent to the room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#maximum_message_rate_per_second IvschatRoom#maximum_message_rate_per_second}
	MaximumMessageRatePerSecond *float64 `field:"optional" json:"maximumMessageRatePerSecond" yaml:"maximumMessageRatePerSecond"`
	// Configuration information for optional review of messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#message_review_handler IvschatRoom#message_review_handler}
	MessageReviewHandler *IvschatRoomMessageReviewHandler `field:"optional" json:"messageReviewHandler" yaml:"messageReviewHandler"`
	// The name of the room. The value does not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#name IvschatRoom#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_room#tags IvschatRoom#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

