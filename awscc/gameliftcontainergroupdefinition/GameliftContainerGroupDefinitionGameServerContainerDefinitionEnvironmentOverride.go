package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionGameServerContainerDefinitionEnvironmentOverride struct {
	// The environment variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_group_definition#name GameliftContainerGroupDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The environment variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_group_definition#value GameliftContainerGroupDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

