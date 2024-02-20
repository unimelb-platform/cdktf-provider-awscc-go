package athenaworkgroup


type AthenaWorkGroupWorkGroupConfiguration struct {
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
	CustomerContentEncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfiguration `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#enforce_work_group_configuration AthenaWorkGroup#enforce_work_group_configuration}
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The Athena engine version for running queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#engine_version AthenaWorkGroup#engine_version}
	EngineVersion *AthenaWorkGroupWorkGroupConfigurationEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Execution Role ARN required to run Athena Spark Calculations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#execution_role AthenaWorkGroup#execution_role}
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#publish_cloudwatch_metrics_enabled AthenaWorkGroup#publish_cloudwatch_metrics_enabled}
	PublishCloudwatchMetricsEnabled interface{} `field:"optional" json:"publishCloudwatchMetricsEnabled" yaml:"publishCloudwatchMetricsEnabled"`
	// If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#requester_pays_enabled AthenaWorkGroup#requester_pays_enabled}
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results.
	//
	// These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#result_configuration AthenaWorkGroup#result_configuration}
	ResultConfiguration *AthenaWorkGroupWorkGroupConfigurationResultConfiguration `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
}

