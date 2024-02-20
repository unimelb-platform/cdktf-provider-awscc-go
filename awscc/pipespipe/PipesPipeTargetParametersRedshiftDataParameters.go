package pipespipe


type PipesPipeTargetParametersRedshiftDataParameters struct {
	// Redshift Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#database PipesPipe#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// A list of SQLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#sqls PipesPipe#sqls}
	Sqls *[]*string `field:"required" json:"sqls" yaml:"sqls"`
	// Database user name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#db_user PipesPipe#db_user}
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#secret_manager_arn PipesPipe#secret_manager_arn}
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// A name for Redshift DataAPI statement which can be used as filter of ListStatement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#statement_name PipesPipe#statement_name}
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#with_event PipesPipe#with_event}.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

