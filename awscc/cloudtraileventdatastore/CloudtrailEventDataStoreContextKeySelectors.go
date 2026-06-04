package cloudtraileventdatastore


type CloudtrailEventDataStoreContextKeySelectors struct {
	// An operator that includes events that match the exact value of the event record field specified in Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudtrail_event_data_store#equals CloudtrailEventDataStore#equals}
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Specifies the type of the event record field in ContextKeySelector. Valid values include RequestContext, TagContext.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudtrail_event_data_store#type CloudtrailEventDataStore#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

