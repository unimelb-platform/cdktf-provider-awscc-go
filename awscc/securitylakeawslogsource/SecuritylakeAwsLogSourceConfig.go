package securitylakeawslogsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecuritylakeAwsLogSourceConfig struct {
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
	// The ARN for the data lake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_aws_log_source#data_lake_arn SecuritylakeAwsLogSource#data_lake_arn}
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
	// The name for a AWS source. This must be a Regionally unique value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_aws_log_source#source_name SecuritylakeAwsLogSource#source_name}
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The version for a AWS source. This must be a Regionally unique value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_aws_log_source#source_version SecuritylakeAwsLogSource#source_version}
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
	// AWS account where you want to collect logs from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_aws_log_source#accounts SecuritylakeAwsLogSource#accounts}
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
}

