package datasynctask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncTaskConfig struct {
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
	// The ARN of an AWS storage resource's location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#destination_location_arn DatasyncTask#destination_location_arn}
	DestinationLocationArn *string `field:"required" json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// The ARN of the source location for the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#source_location_arn DatasyncTask#source_location_arn}
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#cloudwatch_log_group_arn DatasyncTask#cloudwatch_log_group_arn}
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#excludes DatasyncTask#excludes}.
	Excludes interface{} `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#includes DatasyncTask#includes}.
	Includes interface{} `field:"optional" json:"includes" yaml:"includes"`
	// The name of a task.
	//
	// This value is a text reference that is used to identify the task in the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#name DatasyncTask#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents the options that are available to control the behavior of a StartTaskExecution operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#options DatasyncTask#options}
	Options *DatasyncTaskOptions `field:"optional" json:"options" yaml:"options"`
	// Specifies the schedule you want your task to use for repeated executions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#schedule DatasyncTask#schedule}
	Schedule *DatasyncTaskSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#tags DatasyncTask#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies how you want to configure a task report, which provides detailed information about for your Datasync transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#task_report_config DatasyncTask#task_report_config}
	TaskReportConfig *DatasyncTaskTaskReportConfig `field:"optional" json:"taskReportConfig" yaml:"taskReportConfig"`
}

