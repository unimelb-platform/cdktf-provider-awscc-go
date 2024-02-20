package configconfigurationaggregator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigurationAggregatorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_configuration_aggregator#account_aggregation_sources ConfigConfigurationAggregator#account_aggregation_sources}.
	AccountAggregationSources interface{} `field:"optional" json:"accountAggregationSources" yaml:"accountAggregationSources"`
	// The name of the aggregator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_configuration_aggregator#configuration_aggregator_name ConfigConfigurationAggregator#configuration_aggregator_name}
	ConfigurationAggregatorName *string `field:"optional" json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_configuration_aggregator#organization_aggregation_source ConfigConfigurationAggregator#organization_aggregation_source}.
	OrganizationAggregationSource *ConfigConfigurationAggregatorOrganizationAggregationSource `field:"optional" json:"organizationAggregationSource" yaml:"organizationAggregationSource"`
	// The tags for the configuration aggregator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_configuration_aggregator#tags ConfigConfigurationAggregator#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

