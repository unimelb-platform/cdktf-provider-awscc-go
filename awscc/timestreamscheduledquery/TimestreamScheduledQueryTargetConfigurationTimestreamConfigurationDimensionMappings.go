package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationDimensionMappings struct {
	// Type for the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#dimension_value_type TimestreamScheduledQuery#dimension_value_type}
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Column name from query result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#name TimestreamScheduledQuery#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

