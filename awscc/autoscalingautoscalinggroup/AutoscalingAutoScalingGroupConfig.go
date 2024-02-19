package autoscalingautoscalinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingAutoScalingGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#max_size AutoscalingAutoScalingGroup#max_size}.
	MaxSize *string `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#min_size AutoscalingAutoScalingGroup#min_size}.
	MinSize *string `field:"required" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#auto_scaling_group_name AutoscalingAutoScalingGroup#auto_scaling_group_name}.
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#availability_zones AutoscalingAutoScalingGroup#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#capacity_rebalance AutoscalingAutoScalingGroup#capacity_rebalance}.
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#context AutoscalingAutoScalingGroup#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#cooldown AutoscalingAutoScalingGroup#cooldown}.
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#default_instance_warmup AutoscalingAutoScalingGroup#default_instance_warmup}.
	DefaultInstanceWarmup *float64 `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#desired_capacity AutoscalingAutoScalingGroup#desired_capacity}.
	DesiredCapacity *string `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#desired_capacity_type AutoscalingAutoScalingGroup#desired_capacity_type}.
	DesiredCapacityType *string `field:"optional" json:"desiredCapacityType" yaml:"desiredCapacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#health_check_grace_period AutoscalingAutoScalingGroup#health_check_grace_period}.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#health_check_type AutoscalingAutoScalingGroup#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#instance_id AutoscalingAutoScalingGroup#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#instance_maintenance_policy AutoscalingAutoScalingGroup#instance_maintenance_policy}.
	InstanceMaintenancePolicy *AutoscalingAutoScalingGroupInstanceMaintenancePolicy `field:"optional" json:"instanceMaintenancePolicy" yaml:"instanceMaintenancePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#launch_configuration_name AutoscalingAutoScalingGroup#launch_configuration_name}.
	LaunchConfigurationName *string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#launch_template AutoscalingAutoScalingGroup#launch_template}.
	LaunchTemplate *AutoscalingAutoScalingGroupLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#lifecycle_hook_specification_list AutoscalingAutoScalingGroup#lifecycle_hook_specification_list}.
	LifecycleHookSpecificationList interface{} `field:"optional" json:"lifecycleHookSpecificationList" yaml:"lifecycleHookSpecificationList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#load_balancer_names AutoscalingAutoScalingGroup#load_balancer_names}.
	LoadBalancerNames *[]*string `field:"optional" json:"loadBalancerNames" yaml:"loadBalancerNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#max_instance_lifetime AutoscalingAutoScalingGroup#max_instance_lifetime}.
	MaxInstanceLifetime *float64 `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#metrics_collection AutoscalingAutoScalingGroup#metrics_collection}.
	MetricsCollection interface{} `field:"optional" json:"metricsCollection" yaml:"metricsCollection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#mixed_instances_policy AutoscalingAutoScalingGroup#mixed_instances_policy}.
	MixedInstancesPolicy *AutoscalingAutoScalingGroupMixedInstancesPolicy `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#new_instances_protected_from_scale_in AutoscalingAutoScalingGroup#new_instances_protected_from_scale_in}.
	NewInstancesProtectedFromScaleIn interface{} `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#notification_configuration AutoscalingAutoScalingGroup#notification_configuration}.
	NotificationConfiguration *AutoscalingAutoScalingGroupNotificationConfiguration `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#notification_configurations AutoscalingAutoScalingGroup#notification_configurations}.
	NotificationConfigurations interface{} `field:"optional" json:"notificationConfigurations" yaml:"notificationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#placement_group AutoscalingAutoScalingGroup#placement_group}.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#service_linked_role_arn AutoscalingAutoScalingGroup#service_linked_role_arn}.
	ServiceLinkedRoleArn *string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#tags AutoscalingAutoScalingGroup#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#target_group_ar_ns AutoscalingAutoScalingGroup#target_group_ar_ns}.
	TargetGroupArNs *[]*string `field:"optional" json:"targetGroupArNs" yaml:"targetGroupArNs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#termination_policies AutoscalingAutoScalingGroup#termination_policies}.
	TerminationPolicies *[]*string `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#vpc_zone_identifier AutoscalingAutoScalingGroup#vpc_zone_identifier}.
	VpcZoneIdentifier *[]*string `field:"optional" json:"vpcZoneIdentifier" yaml:"vpcZoneIdentifier"`
}

