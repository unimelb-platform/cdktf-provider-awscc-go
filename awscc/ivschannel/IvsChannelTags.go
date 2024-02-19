package ivschannel


type IvsChannelTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#key IvsChannel#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#value IvsChannel#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

