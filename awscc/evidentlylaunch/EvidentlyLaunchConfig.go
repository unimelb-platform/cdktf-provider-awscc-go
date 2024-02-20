package evidentlylaunch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EvidentlyLaunchConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#groups EvidentlyLaunch#groups}.
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#name EvidentlyLaunch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#project EvidentlyLaunch#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#scheduled_splits_config EvidentlyLaunch#scheduled_splits_config}.
	ScheduledSplitsConfig interface{} `field:"required" json:"scheduledSplitsConfig" yaml:"scheduledSplitsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#description EvidentlyLaunch#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Start or Stop Launch Launch. Default is not started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#execution_status EvidentlyLaunch#execution_status}
	ExecutionStatus *EvidentlyLaunchExecutionStatus `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#metric_monitors EvidentlyLaunch#metric_monitors}.
	MetricMonitors interface{} `field:"optional" json:"metricMonitors" yaml:"metricMonitors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#randomization_salt EvidentlyLaunch#randomization_salt}.
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#tags EvidentlyLaunch#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

