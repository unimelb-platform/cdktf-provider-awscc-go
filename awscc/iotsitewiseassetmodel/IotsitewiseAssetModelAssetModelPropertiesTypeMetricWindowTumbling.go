package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelPropertiesTypeMetricWindowTumbling struct {
	// The time interval for the tumbling window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#interval IotsitewiseAssetModel#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The shift or reference point on timeline for the contiguous time intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#offset IotsitewiseAssetModel#offset}
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
}

