package logsmetricfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsMetricFilterConfig struct {
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
	// Pattern that Logs follows to interpret each entry in a log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#filter_pattern LogsMetricFilter#filter_pattern}
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// Existing log group that you want to associate with this filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#log_group_name LogsMetricFilter#log_group_name}
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// A collection of information that defines how metric data gets emitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#metric_transformations LogsMetricFilter#metric_transformations}
	MetricTransformations interface{} `field:"required" json:"metricTransformations" yaml:"metricTransformations"`
	// A name for the metric filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#filter_name LogsMetricFilter#filter_name}
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

