package connectuserhierarchystructure


type ConnectUserHierarchyStructureUserHierarchyStructureLevelThree struct {
	// The Amazon Resource Name (ARN) of the hierarchy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_user_hierarchy_structure#hierarchy_level_arn ConnectUserHierarchyStructure#hierarchy_level_arn}
	HierarchyLevelArn *string `field:"optional" json:"hierarchyLevelArn" yaml:"hierarchyLevelArn"`
	// The identifier of the hierarchy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_user_hierarchy_structure#hierarchy_level_id ConnectUserHierarchyStructure#hierarchy_level_id}
	HierarchyLevelId *string `field:"optional" json:"hierarchyLevelId" yaml:"hierarchyLevelId"`
	// The name of the hierarchy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_user_hierarchy_structure#name ConnectUserHierarchyStructure#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

