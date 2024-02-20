package lakeformationtagassociation


type LakeformationTagAssociationResourceDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#catalog_id LakeformationTagAssociation#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_tag_association#name LakeformationTagAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

