package rolesanywhereprofile


type RolesanywhereProfileAttributeMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rolesanywhere_profile#certificate_field RolesanywhereProfile#certificate_field}.
	CertificateField *string `field:"optional" json:"certificateField" yaml:"certificateField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rolesanywhere_profile#mapping_rules RolesanywhereProfile#mapping_rules}.
	MappingRules interface{} `field:"optional" json:"mappingRules" yaml:"mappingRules"`
}

