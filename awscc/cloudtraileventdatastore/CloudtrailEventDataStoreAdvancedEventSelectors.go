package cloudtraileventdatastore


type CloudtrailEventDataStoreAdvancedEventSelectors struct {
	// Contains all selector statements in an advanced event selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_event_data_store#field_selectors CloudtrailEventDataStore#field_selectors}
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_event_data_store#name CloudtrailEventDataStore#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

