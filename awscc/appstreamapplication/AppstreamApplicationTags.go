package appstreamapplication


type AppstreamApplicationTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#key AppstreamApplication#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#tag_key AppstreamApplication#tag_key}.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#tag_value AppstreamApplication#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#value AppstreamApplication#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

