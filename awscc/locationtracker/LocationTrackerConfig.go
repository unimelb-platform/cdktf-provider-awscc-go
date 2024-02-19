package locationtracker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LocationTrackerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#tracker_name LocationTracker#tracker_name}.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#description LocationTracker#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#event_bridge_enabled LocationTracker#event_bridge_enabled}.
	EventBridgeEnabled interface{} `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#kms_key_enable_geospatial_queries LocationTracker#kms_key_enable_geospatial_queries}.
	KmsKeyEnableGeospatialQueries interface{} `field:"optional" json:"kmsKeyEnableGeospatialQueries" yaml:"kmsKeyEnableGeospatialQueries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#kms_key_id LocationTracker#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#position_filtering LocationTracker#position_filtering}.
	PositionFiltering *string `field:"optional" json:"positionFiltering" yaml:"positionFiltering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#pricing_plan LocationTracker#pricing_plan}.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#pricing_plan_data_source LocationTracker#pricing_plan_data_source}.
	PricingPlanDataSource *string `field:"optional" json:"pricingPlanDataSource" yaml:"pricingPlanDataSource"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_tracker#tags LocationTracker#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

