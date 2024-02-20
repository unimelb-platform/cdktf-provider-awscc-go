package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration struct {
	// Describes whether the application uses the default parallelism for the Kinesis Data Analytics service.
	//
	// You must set this property to `CUSTOM` in order to change your application's `AutoScalingEnabled`, `Parallelism`, or `ParallelismPerKPU` properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#configuration_type Kinesisanalyticsv2Application#configuration_type}
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes whether the Kinesis Data Analytics service can increase the parallelism of the application in response to increased throughput.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#auto_scaling_enabled Kinesisanalyticsv2Application#auto_scaling_enabled}
	AutoScalingEnabled interface{} `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Describes the initial number of parallel tasks that a Java-based Kinesis Data Analytics application can perform.
	//
	// The Kinesis Data Analytics service can increase this number automatically if ParallelismConfiguration:AutoScalingEnabled is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#parallelism Kinesisanalyticsv2Application#parallelism}
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Describes the number of parallel tasks that a Java-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application.
	//
	// For more information about KPUs, see Amazon Kinesis Data Analytics Pricing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#parallelism_per_kpu Kinesisanalyticsv2Application#parallelism_per_kpu}
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
}

