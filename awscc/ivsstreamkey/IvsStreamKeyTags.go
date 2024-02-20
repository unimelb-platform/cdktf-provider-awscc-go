package ivsstreamkey


type IvsStreamKeyTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_stream_key#key IvsStreamKey#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_stream_key#value IvsStreamKey#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

