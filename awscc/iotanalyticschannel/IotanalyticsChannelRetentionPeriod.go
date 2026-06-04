package iotanalyticschannel


type IotanalyticsChannelRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_channel#number_of_days IotanalyticsChannel#number_of_days}.
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_channel#unlimited IotanalyticsChannel#unlimited}.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

