package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelPropertiesTypeMetricWindow struct {
	// Contains a tumbling window, which is a repeating fixed-sized, non-overlapping, and contiguous time interval.
	//
	// This window is used in metric and aggregation computations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#tumbling IotsitewiseAssetModel#tumbling}
	Tumbling *IotsitewiseAssetModelAssetModelPropertiesTypeMetricWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

