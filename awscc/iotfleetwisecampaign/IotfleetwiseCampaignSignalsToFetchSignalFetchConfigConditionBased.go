package iotfleetwisecampaign


type IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBased struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_campaign#condition_expression IotfleetwiseCampaign#condition_expression}.
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_campaign#trigger_mode IotfleetwiseCampaign#trigger_mode}.
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

