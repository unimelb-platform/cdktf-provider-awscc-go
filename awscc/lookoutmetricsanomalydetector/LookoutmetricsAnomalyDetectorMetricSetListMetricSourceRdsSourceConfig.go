package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#database_host LookoutmetricsAnomalyDetector#database_host}.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#database_name LookoutmetricsAnomalyDetector#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#database_port LookoutmetricsAnomalyDetector#database_port}.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#db_instance_identifier LookoutmetricsAnomalyDetector#db_instance_identifier}.
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#role_arn LookoutmetricsAnomalyDetector#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#secret_manager_arn LookoutmetricsAnomalyDetector#secret_manager_arn}.
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#table_name LookoutmetricsAnomalyDetector#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#vpc_configuration LookoutmetricsAnomalyDetector#vpc_configuration}.
	VpcConfiguration *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

