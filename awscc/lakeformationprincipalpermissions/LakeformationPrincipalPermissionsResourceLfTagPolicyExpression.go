package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceLfTagPolicyExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#tag_key LakeformationPrincipalPermissions#tag_key}.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#tag_values LakeformationPrincipalPermissions#tag_values}.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

