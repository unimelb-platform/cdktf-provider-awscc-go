package codeguruprofilerprofilinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodeguruprofilerProfilingGroupConfig struct {
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
	// The name of the profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#profiling_group_name CodeguruprofilerProfilingGroup#profiling_group_name}
	ProfilingGroupName *string `field:"required" json:"profilingGroupName" yaml:"profilingGroupName"`
	// The agent permissions attached to this profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#agent_permissions CodeguruprofilerProfilingGroup#agent_permissions}
	AgentPermissions *CodeguruprofilerProfilingGroupAgentPermissions `field:"optional" json:"agentPermissions" yaml:"agentPermissions"`
	// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#anomaly_detection_notification_configuration CodeguruprofilerProfilingGroup#anomaly_detection_notification_configuration}
	AnomalyDetectionNotificationConfiguration interface{} `field:"optional" json:"anomalyDetectionNotificationConfiguration" yaml:"anomalyDetectionNotificationConfiguration"`
	// The compute platform of the profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#compute_platform CodeguruprofilerProfilingGroup#compute_platform}
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// The tags associated with a profiling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeguruprofiler_profiling_group#tags CodeguruprofilerProfilingGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

