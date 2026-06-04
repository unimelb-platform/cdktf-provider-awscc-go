package guarddutypublishingdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyPublishingDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_publishing_destination#destination_properties GuarddutyPublishingDestination#destination_properties}.
	DestinationProperties *GuarddutyPublishingDestinationDestinationProperties `field:"required" json:"destinationProperties" yaml:"destinationProperties"`
	// The type of resource for the publishing destination. Currently only Amazon S3 buckets are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_publishing_destination#destination_type GuarddutyPublishingDestination#destination_type}
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// The ID of the GuardDuty detector associated with the publishing destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_publishing_destination#detector_id GuarddutyPublishingDestination#detector_id}
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_publishing_destination#tags GuarddutyPublishingDestination#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

