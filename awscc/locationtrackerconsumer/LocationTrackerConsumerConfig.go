package locationtrackerconsumer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LocationTrackerConsumerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker_consumer#consumer_arn LocationTrackerConsumer#consumer_arn}.
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker_consumer#tracker_name LocationTrackerConsumer#tracker_name}.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
}

