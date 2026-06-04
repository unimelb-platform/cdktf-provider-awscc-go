package gluedatabase


type GlueDatabaseDatabaseInputCreateTableDefaultPermissionsPrincipal struct {
	// An identifier for the AWS Lake Formation principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_database#data_lake_principal_identifier GlueDatabase#data_lake_principal_identifier}
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

