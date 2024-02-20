package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#catalog LakeformationPrincipalPermissions#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#database LakeformationPrincipalPermissions#database}.
	Database *LakeformationPrincipalPermissionsResourceDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#data_cells_filter LakeformationPrincipalPermissions#data_cells_filter}.
	DataCellsFilter *LakeformationPrincipalPermissionsResourceDataCellsFilter `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#data_location LakeformationPrincipalPermissions#data_location}.
	DataLocation *LakeformationPrincipalPermissionsResourceDataLocation `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#lf_tag LakeformationPrincipalPermissions#lf_tag}.
	LfTag *LakeformationPrincipalPermissionsResourceLfTag `field:"optional" json:"lfTag" yaml:"lfTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#lf_tag_policy LakeformationPrincipalPermissions#lf_tag_policy}.
	LfTagPolicy *LakeformationPrincipalPermissionsResourceLfTagPolicy `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#table LakeformationPrincipalPermissions#table}.
	Table *LakeformationPrincipalPermissionsResourceTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#table_with_columns LakeformationPrincipalPermissions#table_with_columns}.
	TableWithColumns *LakeformationPrincipalPermissionsResourceTableWithColumns `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

