package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfiguration struct {
	// The code location and type parameters for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration Kinesisanalyticsv2Application#application_code_configuration}
	ApplicationCodeConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration `field:"optional" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration Kinesisanalyticsv2Application#application_snapshot_configuration}
	ApplicationSnapshotConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// Describes execution properties for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#environment_properties Kinesisanalyticsv2Application#environment_properties}
	EnvironmentProperties *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// The creation and update parameters for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration Kinesisanalyticsv2Application#flink_application_configuration}
	FlinkApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// The creation and update parameters for a SQL-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration Kinesisanalyticsv2Application#sql_application_configuration}
	SqlApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// The array of descriptions of VPC configurations available to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#vpc_configurations Kinesisanalyticsv2Application#vpc_configurations}
	VpcConfigurations interface{} `field:"optional" json:"vpcConfigurations" yaml:"vpcConfigurations"`
	// The configuration parameters for a Kinesis Data Analytics Studio notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#zeppelin_application_configuration Kinesisanalyticsv2Application#zeppelin_application_configuration}
	ZeppelinApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfiguration `field:"optional" json:"zeppelinApplicationConfiguration" yaml:"zeppelinApplicationConfiguration"`
}

