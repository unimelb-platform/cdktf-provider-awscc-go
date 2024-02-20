package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappings struct {
	// Value type of the measure value column to be read from the query result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#measure_value_type TimestreamScheduledQuery#measure_value_type}
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Source measure value column in the query result where the attribute value is to be read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#source_column TimestreamScheduledQuery#source_column}
	SourceColumn *string `field:"required" json:"sourceColumn" yaml:"sourceColumn"`
	// Custom name to be used for attribute name in derived table.
	//
	// If not provided, source column name would be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#target_multi_measure_attribute_name TimestreamScheduledQuery#target_multi_measure_attribute_name}
	TargetMultiMeasureAttributeName *string `field:"optional" json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}

