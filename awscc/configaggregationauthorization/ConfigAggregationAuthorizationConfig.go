package configaggregationauthorization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigAggregationAuthorizationConfig struct {
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
	// The 12-digit account ID of the account authorized to aggregate data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_aggregation_authorization#authorized_account_id ConfigAggregationAuthorization#authorized_account_id}
	AuthorizedAccountId *string `field:"required" json:"authorizedAccountId" yaml:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_aggregation_authorization#authorized_aws_region ConfigAggregationAuthorization#authorized_aws_region}
	AuthorizedAwsRegion *string `field:"required" json:"authorizedAwsRegion" yaml:"authorizedAwsRegion"`
	// The tags for the AggregationAuthorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_aggregation_authorization#tags ConfigAggregationAuthorization#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

