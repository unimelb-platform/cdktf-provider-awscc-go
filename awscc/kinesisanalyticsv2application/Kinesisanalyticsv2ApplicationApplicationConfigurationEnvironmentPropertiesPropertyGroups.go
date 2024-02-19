package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroups struct {
	// Describes the key of an application execution property key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#property_group_id Kinesisanalyticsv2Application#property_group_id}
	PropertyGroupId *string `field:"optional" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Describes the value of an application execution property key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#property_map Kinesisanalyticsv2Application#property_map}
	PropertyMap *map[string]*string `field:"optional" json:"propertyMap" yaml:"propertyMap"`
}

