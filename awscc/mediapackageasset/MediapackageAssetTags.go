package mediapackageasset


type MediapackageAssetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_asset#key MediapackageAsset#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_asset#value MediapackageAsset#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

