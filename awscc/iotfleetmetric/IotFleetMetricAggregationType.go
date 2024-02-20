package iotfleetmetric


type IotFleetMetricAggregationType struct {
	// Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_fleet_metric#name IotFleetMetric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fleet Indexing aggregation type values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_fleet_metric#values IotFleetMetric#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

