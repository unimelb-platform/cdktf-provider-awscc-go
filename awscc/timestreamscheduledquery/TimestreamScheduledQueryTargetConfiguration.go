package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfiguration struct {
	// Configuration needed to write data into the Timestream database and table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/timestream_scheduled_query#timestream_configuration TimestreamScheduledQuery#timestream_configuration}
	TimestreamConfiguration *TimestreamScheduledQueryTargetConfigurationTimestreamConfiguration `field:"optional" json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

