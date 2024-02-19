package iotfleetwisecampaign


type IotfleetwiseCampaignCollectionScheme struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_campaign#condition_based_collection_scheme IotfleetwiseCampaign#condition_based_collection_scheme}.
	ConditionBasedCollectionScheme *IotfleetwiseCampaignCollectionSchemeConditionBasedCollectionScheme `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_campaign#time_based_collection_scheme IotfleetwiseCampaign#time_based_collection_scheme}.
	TimeBasedCollectionScheme *IotfleetwiseCampaignCollectionSchemeTimeBasedCollectionScheme `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

