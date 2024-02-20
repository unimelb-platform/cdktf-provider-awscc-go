package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdates struct {
	// Additional Configuration that are passed to Athena Spark Calculations running in this workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#additional_configuration AthenaWorkGroup#additional_configuration}
	AdditionalConfiguration *string `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#bytes_scanned_cutoff_per_query AthenaWorkGroup#bytes_scanned_cutoff_per_query}
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// Indicates the KMS key for encrypting notebook content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#customer_content_encryption_configuration AthenaWorkGroup#customer_content_encryption_configuration}
	CustomerContentEncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationUpdatesCustomerContentEncryptionConfiguration `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#enforce_work_group_configuration AthenaWorkGroup#enforce_work_group_configuration}
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The Athena engine version for running queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#engine_version AthenaWorkGroup#engine_version}
	EngineVersion *AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Execution Role ARN required to run Athena Spark Calculations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#execution_role AthenaWorkGroup#execution_role}
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#publish_cloudwatch_metrics_enabled AthenaWorkGroup#publish_cloudwatch_metrics_enabled}
	PublishCloudwatchMetricsEnabled interface{} `field:"optional" json:"publishCloudwatchMetricsEnabled" yaml:"publishCloudwatchMetricsEnabled"`
	// Indicates that the data usage control limit per query is removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#remove_bytes_scanned_cutoff_per_query AthenaWorkGroup#remove_bytes_scanned_cutoff_per_query}
	RemoveBytesScannedCutoffPerQuery interface{} `field:"optional" json:"removeBytesScannedCutoffPerQuery" yaml:"removeBytesScannedCutoffPerQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#remove_customer_content_encryption_configuration AthenaWorkGroup#remove_customer_content_encryption_configuration}.
	RemoveCustomerContentEncryptionConfiguration interface{} `field:"optional" json:"removeCustomerContentEncryptionConfiguration" yaml:"removeCustomerContentEncryptionConfiguration"`
	// If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#requester_pays_enabled AthenaWorkGroup#requester_pays_enabled}
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// The result configuration information about the queries in this workgroup that will be updated.
	//
	// Includes the updated results location and an updated option for encrypting query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#result_configuration_updates AthenaWorkGroup#result_configuration_updates}
	ResultConfigurationUpdates *AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdates `field:"optional" json:"resultConfigurationUpdates" yaml:"resultConfigurationUpdates"`
}

