package quicksightdatasource


type QuicksightDataSourceDataSourceParametersAthenaParameters struct {
	// <p>Use the <code>RoleArn</code> structure to override an account-wide role for a specific Athena data source.
	//
	// For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use <code>RoleArn</code> to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#role_arn QuicksightDataSource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// <p>The workgroup that Amazon Athena uses.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#work_group QuicksightDataSource#work_group}
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

