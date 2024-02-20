package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#catalog_id LakeformationPrincipalPermissions#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#name LakeformationPrincipalPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

