package mediapackageasset


type MediapackageAssetEgressEndpoints struct {
	// The ID of the PackagingConfiguration being applied to the Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_asset#packaging_configuration_id MediapackageAsset#packaging_configuration_id}
	PackagingConfigurationId *string `field:"required" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The URL of the parent manifest for the repackaged Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_asset#url MediapackageAsset#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

