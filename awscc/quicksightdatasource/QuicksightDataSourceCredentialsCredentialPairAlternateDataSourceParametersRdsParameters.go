package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRdsParameters struct {
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// <p>Instance ID.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#instance_id QuicksightDataSource#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

