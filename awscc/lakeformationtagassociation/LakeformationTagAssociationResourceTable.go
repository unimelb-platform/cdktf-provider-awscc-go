package lakeformationtagassociation


type LakeformationTagAssociationResourceTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#catalog_id LakeformationTagAssociation#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#database_name LakeformationTagAssociation#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#name LakeformationTagAssociation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#table_wildcard LakeformationTagAssociation#table_wildcard}.
	TableWildcard *string `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

