package athenapreparedstatement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaPreparedStatementConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The query string for the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_prepared_statement#query_statement AthenaPreparedStatement#query_statement}
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The name of the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_prepared_statement#statement_name AthenaPreparedStatement#statement_name}
	StatementName *string `field:"required" json:"statementName" yaml:"statementName"`
	// The name of the workgroup to which the prepared statement belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_prepared_statement#work_group AthenaPreparedStatement#work_group}
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
	// The description of the prepared statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_prepared_statement#description AthenaPreparedStatement#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

