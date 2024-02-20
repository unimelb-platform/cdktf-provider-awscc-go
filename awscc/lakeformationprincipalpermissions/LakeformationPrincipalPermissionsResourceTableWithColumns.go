package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceTableWithColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#catalog_id LakeformationPrincipalPermissions#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#database_name LakeformationPrincipalPermissions#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#name LakeformationPrincipalPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#column_names LakeformationPrincipalPermissions#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#column_wildcard LakeformationPrincipalPermissions#column_wildcard}.
	ColumnWildcard *LakeformationPrincipalPermissionsResourceTableWithColumnsColumnWildcard `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
}

