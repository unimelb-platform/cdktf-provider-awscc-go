package applicationinsightsapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationinsightsApplicationConfig struct {
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
	// The name of the resource group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#resource_group_name ApplicationinsightsApplication#resource_group_name}
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// If set to true, application will be configured with recommended monitoring configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#auto_configuration_enabled ApplicationinsightsApplication#auto_configuration_enabled}
	AutoConfigurationEnabled interface{} `field:"optional" json:"autoConfigurationEnabled" yaml:"autoConfigurationEnabled"`
	// The monitoring settings of the components.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#component_monitoring_settings ApplicationinsightsApplication#component_monitoring_settings}
	ComponentMonitoringSettings interface{} `field:"optional" json:"componentMonitoringSettings" yaml:"componentMonitoringSettings"`
	// The custom grouped components.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#custom_components ApplicationinsightsApplication#custom_components}
	CustomComponents interface{} `field:"optional" json:"customComponents" yaml:"customComponents"`
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#cwe_monitor_enabled ApplicationinsightsApplication#cwe_monitor_enabled}
	CweMonitorEnabled interface{} `field:"optional" json:"cweMonitorEnabled" yaml:"cweMonitorEnabled"`
	// The grouping type of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#grouping_type ApplicationinsightsApplication#grouping_type}
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// The log pattern sets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#log_pattern_sets ApplicationinsightsApplication#log_pattern_sets}
	LogPatternSets interface{} `field:"optional" json:"logPatternSets" yaml:"logPatternSets"`
	// When set to true, creates opsItems for any problems detected on an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#ops_center_enabled ApplicationinsightsApplication#ops_center_enabled}
	OpsCenterEnabled interface{} `field:"optional" json:"opsCenterEnabled" yaml:"opsCenterEnabled"`
	// The SNS topic provided to Application Insights that is associated to the created opsItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#ops_item_sns_topic_arn ApplicationinsightsApplication#ops_item_sns_topic_arn}
	OpsItemSnsTopicArn *string `field:"optional" json:"opsItemSnsTopicArn" yaml:"opsItemSnsTopicArn"`
	// The tags of Application Insights application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#tags ApplicationinsightsApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

