package cloudtraileventdatastore


type CloudtrailEventDataStoreInsightSelectors struct {
	// The type of Insights to log on an event data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_event_data_store#insight_type CloudtrailEventDataStore#insight_type}
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

