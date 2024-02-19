package cloudwatchmetricstream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchMetricStreamConfig struct {
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
	// The ARN of the Kinesis Firehose where to stream the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#firehose_arn CloudwatchMetricStream#firehose_arn}
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
	// The output format of the data streamed to the Kinesis Firehose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#output_format CloudwatchMetricStream#output_format}
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The ARN of the role that provides access to the Kinesis Firehose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#role_arn CloudwatchMetricStream#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Define which metrics will be not streamed.
	//
	// Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#exclude_filters CloudwatchMetricStream#exclude_filters}
	ExcludeFilters interface{} `field:"optional" json:"excludeFilters" yaml:"excludeFilters"`
	// Define which metrics will be streamed.
	//
	// Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#include_filters CloudwatchMetricStream#include_filters}
	IncludeFilters interface{} `field:"optional" json:"includeFilters" yaml:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream.
	//
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#include_linked_accounts_metrics CloudwatchMetricStream#include_linked_accounts_metrics}
	IncludeLinkedAccountsMetrics interface{} `field:"optional" json:"includeLinkedAccountsMetrics" yaml:"includeLinkedAccountsMetrics"`
	// Name of the metric stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#name CloudwatchMetricStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed.
	//
	// You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#statistics_configurations CloudwatchMetricStream#statistics_configurations}
	StatisticsConfigurations interface{} `field:"optional" json:"statisticsConfigurations" yaml:"statisticsConfigurations"`
	// A set of tags to assign to the delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#tags CloudwatchMetricStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

