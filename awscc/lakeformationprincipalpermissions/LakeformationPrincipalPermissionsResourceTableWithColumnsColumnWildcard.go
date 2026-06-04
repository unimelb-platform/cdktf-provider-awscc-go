package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceTableWithColumnsColumnWildcard struct {
	// Excludes column names. Any column with this name will be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lakeformation_principal_permissions#excluded_column_names LakeformationPrincipalPermissions#excluded_column_names}
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

