package quicksighttopic


type QuicksightTopicDataSetsNamedEntitiesSemanticEntityType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#sub_type_name QuicksightTopic#sub_type_name}.
	SubTypeName *string `field:"optional" json:"subTypeName" yaml:"subTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#type_name QuicksightTopic#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#type_parameters QuicksightTopic#type_parameters}.
	TypeParameters *map[string]*string `field:"optional" json:"typeParameters" yaml:"typeParameters"`
}

