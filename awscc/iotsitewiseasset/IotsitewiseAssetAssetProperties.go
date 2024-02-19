package iotsitewiseasset


type IotsitewiseAssetAssetProperties struct {
	// Customer provided ID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#logical_id IotsitewiseAsset#logical_id}
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The property alias that identifies the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#alias IotsitewiseAsset#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The MQTT notification state (ENABLED or DISABLED) for this asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#notification_state IotsitewiseAsset#notification_state}
	NotificationState *string `field:"optional" json:"notificationState" yaml:"notificationState"`
	// The unit of measure (such as Newtons or RPM) of the asset property.
	//
	// If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#unit IotsitewiseAsset#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

