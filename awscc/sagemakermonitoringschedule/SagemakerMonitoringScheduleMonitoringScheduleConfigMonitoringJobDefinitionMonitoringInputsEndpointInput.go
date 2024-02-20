package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringInputsEndpointInput struct {
	// The name of the endpoint used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#endpoint_name SagemakerMonitoringSchedule#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#local_path SagemakerMonitoringSchedule#local_path}
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Indexes or names of the features to be excluded from analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#exclude_features_attribute SagemakerMonitoringSchedule#exclude_features_attribute}
	ExcludeFeaturesAttribute *string `field:"optional" json:"excludeFeaturesAttribute" yaml:"excludeFeaturesAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#s3_data_distribution_type SagemakerMonitoringSchedule#s3_data_distribution_type}
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the Pipe or File is used as the input mode for transfering data for the monitoring job.
	//
	// Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#s3_input_mode SagemakerMonitoringSchedule#s3_input_mode}
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

