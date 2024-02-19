package autoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/autoscalingautoscalinggroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group}.
type AutoscalingAutoScalingGroup interface {
	cdktf.TerraformResource
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	AutoScalingGroupNameInput() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	CapacityRebalance() interface{}
	SetCapacityRebalance(val interface{})
	CapacityRebalanceInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	Cooldown() *string
	SetCooldown(val *string)
	CooldownInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultInstanceWarmup() *float64
	SetDefaultInstanceWarmup(val *float64)
	DefaultInstanceWarmupInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCapacity() *string
	SetDesiredCapacity(val *string)
	DesiredCapacityInput() *string
	DesiredCapacityType() *string
	SetDesiredCapacityType(val *string)
	DesiredCapacityTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	HealthCheckGracePeriodInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	Id() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InstanceMaintenancePolicy() AutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference
	InstanceMaintenancePolicyInput() interface{}
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
	LaunchConfigurationNameInput() *string
	LaunchTemplate() AutoscalingAutoScalingGroupLaunchTemplateOutputReference
	LaunchTemplateInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleHookSpecificationList() AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList
	LifecycleHookSpecificationListInput() interface{}
	LoadBalancerNames() *[]*string
	SetLoadBalancerNames(val *[]*string)
	LoadBalancerNamesInput() *[]*string
	MaxInstanceLifetime() *float64
	SetMaxInstanceLifetime(val *float64)
	MaxInstanceLifetimeInput() *float64
	MaxSize() *string
	SetMaxSize(val *string)
	MaxSizeInput() *string
	MetricsCollection() AutoscalingAutoScalingGroupMetricsCollectionList
	MetricsCollectionInput() interface{}
	MinSize() *string
	SetMinSize(val *string)
	MinSizeInput() *string
	MixedInstancesPolicy() AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference
	MixedInstancesPolicyInput() interface{}
	NewInstancesProtectedFromScaleIn() interface{}
	SetNewInstancesProtectedFromScaleIn(val interface{})
	NewInstancesProtectedFromScaleInInput() interface{}
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() AutoscalingAutoScalingGroupNotificationConfigurationOutputReference
	NotificationConfigurationInput() interface{}
	NotificationConfigurations() AutoscalingAutoScalingGroupNotificationConfigurationsList
	NotificationConfigurationsInput() interface{}
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ServiceLinkedRoleArn() *string
	SetServiceLinkedRoleArn(val *string)
	ServiceLinkedRoleArnInput() *string
	Tags() AutoscalingAutoScalingGroupTagsList
	TagsInput() interface{}
	TargetGroupArNs() *[]*string
	SetTargetGroupArNs(val *[]*string)
	TargetGroupArNsInput() *[]*string
	TerminationPolicies() *[]*string
	SetTerminationPolicies(val *[]*string)
	TerminationPoliciesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcZoneIdentifier() *[]*string
	SetVpcZoneIdentifier(val *[]*string)
	VpcZoneIdentifierInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutInstanceMaintenancePolicy(value *AutoscalingAutoScalingGroupInstanceMaintenancePolicy)
	PutLaunchTemplate(value *AutoscalingAutoScalingGroupLaunchTemplate)
	PutLifecycleHookSpecificationList(value interface{})
	PutMetricsCollection(value interface{})
	PutMixedInstancesPolicy(value *AutoscalingAutoScalingGroupMixedInstancesPolicy)
	PutNotificationConfiguration(value *AutoscalingAutoScalingGroupNotificationConfiguration)
	PutNotificationConfigurations(value interface{})
	PutTags(value interface{})
	ResetAutoScalingGroupName()
	ResetAvailabilityZones()
	ResetCapacityRebalance()
	ResetContext()
	ResetCooldown()
	ResetDefaultInstanceWarmup()
	ResetDesiredCapacity()
	ResetDesiredCapacityType()
	ResetHealthCheckGracePeriod()
	ResetHealthCheckType()
	ResetInstanceId()
	ResetInstanceMaintenancePolicy()
	ResetLaunchConfigurationName()
	ResetLaunchTemplate()
	ResetLifecycleHookSpecificationList()
	ResetLoadBalancerNames()
	ResetMaxInstanceLifetime()
	ResetMetricsCollection()
	ResetMixedInstancesPolicy()
	ResetNewInstancesProtectedFromScaleIn()
	ResetNotificationConfiguration()
	ResetNotificationConfigurations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroup()
	ResetServiceLinkedRoleArn()
	ResetTags()
	ResetTargetGroupArNs()
	ResetTerminationPolicies()
	ResetVpcZoneIdentifier()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutoscalingAutoScalingGroup
type jsiiProxy_AutoscalingAutoScalingGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) AutoScalingGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) CapacityRebalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) CooldownInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DefaultInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DesiredCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DesiredCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) DesiredCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) InstanceMaintenancePolicy() AutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference {
	var returns AutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) InstanceMaintenancePolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LaunchConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LaunchTemplate() AutoscalingAutoScalingGroupLaunchTemplateOutputReference {
	var returns AutoscalingAutoScalingGroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LifecycleHookSpecificationList() AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList {
	var returns AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList
	_jsii_.Get(
		j,
		"lifecycleHookSpecificationList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LifecycleHookSpecificationListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHookSpecificationListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LoadBalancerNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) LoadBalancerNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MaxInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MaxSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MetricsCollection() AutoscalingAutoScalingGroupMetricsCollectionList {
	var returns AutoscalingAutoScalingGroupMetricsCollectionList
	_jsii_.Get(
		j,
		"metricsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MetricsCollectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MinSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MinSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MixedInstancesPolicy() AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference {
	var returns AutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) MixedInstancesPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mixedInstancesPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NewInstancesProtectedFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NewInstancesProtectedFromScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NotificationConfiguration() AutoscalingAutoScalingGroupNotificationConfigurationOutputReference {
	var returns AutoscalingAutoScalingGroupNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NotificationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NotificationConfigurations() AutoscalingAutoScalingGroupNotificationConfigurationsList {
	var returns AutoscalingAutoScalingGroupNotificationConfigurationsList
	_jsii_.Get(
		j,
		"notificationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) NotificationConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) ServiceLinkedRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) Tags() AutoscalingAutoScalingGroupTagsList {
	var returns AutoscalingAutoScalingGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TargetGroupArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TargetGroupArNsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArNsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TerminationPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup) VpcZoneIdentifierInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifierInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group} Resource.
func NewAutoscalingAutoScalingGroup(scope constructs.Construct, id *string, config *AutoscalingAutoScalingGroupConfig) AutoscalingAutoScalingGroup {
	_init_.Initialize()

	if err := validateNewAutoscalingAutoScalingGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingAutoScalingGroup{}

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group} Resource.
func NewAutoscalingAutoScalingGroup_Override(a AutoscalingAutoScalingGroup, scope constructs.Construct, id *string, config *AutoscalingAutoScalingGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetAutoScalingGroupName(val *string) {
	if err := j.validateSetAutoScalingGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetCapacityRebalance(val interface{}) {
	if err := j.validateSetCapacityRebalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetCooldown(val *string) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetDefaultInstanceWarmup(val *float64) {
	if err := j.validateSetDefaultInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetDesiredCapacity(val *string) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetDesiredCapacityType(val *string) {
	if err := j.validateSetDesiredCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetLaunchConfigurationName(val *string) {
	if err := j.validateSetLaunchConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetLoadBalancerNames(val *[]*string) {
	if err := j.validateSetLoadBalancerNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerNames",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetMaxInstanceLifetime(val *float64) {
	if err := j.validateSetMaxInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetMaxSize(val *string) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetMinSize(val *string) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetNewInstancesProtectedFromScaleIn(val interface{}) {
	if err := j.validateSetNewInstancesProtectedFromScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetServiceLinkedRoleArn(val *string) {
	if err := j.validateSetServiceLinkedRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetTargetGroupArNs(val *[]*string) {
	if err := j.validateSetTargetGroupArNsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArNs",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetTerminationPolicies(val *[]*string) {
	if err := j.validateSetTerminationPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroup)SetVpcZoneIdentifier(val *[]*string) {
	if err := j.validateSetVpcZoneIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

// Generates CDKTF code for importing a AutoscalingAutoScalingGroup resource upon running "cdktf plan <stack-name>".
func AutoscalingAutoScalingGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAutoscalingAutoScalingGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AutoscalingAutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingAutoScalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingAutoScalingGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingAutoScalingGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingAutoScalingGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingAutoScalingGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutoscalingAutoScalingGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutInstanceMaintenancePolicy(value *AutoscalingAutoScalingGroupInstanceMaintenancePolicy) {
	if err := a.validatePutInstanceMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceMaintenancePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutLaunchTemplate(value *AutoscalingAutoScalingGroupLaunchTemplate) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutLifecycleHookSpecificationList(value interface{}) {
	if err := a.validatePutLifecycleHookSpecificationListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycleHookSpecificationList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutMetricsCollection(value interface{}) {
	if err := a.validatePutMetricsCollectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricsCollection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutMixedInstancesPolicy(value *AutoscalingAutoScalingGroupMixedInstancesPolicy) {
	if err := a.validatePutMixedInstancesPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMixedInstancesPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutNotificationConfiguration(value *AutoscalingAutoScalingGroupNotificationConfiguration) {
	if err := a.validatePutNotificationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotificationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutNotificationConfigurations(value interface{}) {
	if err := a.validatePutNotificationConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotificationConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetAutoScalingGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoScalingGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetCapacityRebalance() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityRebalance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetContext() {
	_jsii_.InvokeVoid(
		a,
		"resetContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetDefaultInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetDesiredCapacityType() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetInstanceId() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetInstanceMaintenancePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceMaintenancePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetLaunchConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetLifecycleHookSpecificationList() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleHookSpecificationList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetLoadBalancerNames() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetMaxInstanceLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInstanceLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetMetricsCollection() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsCollection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetMixedInstancesPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMixedInstancesPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetNewInstancesProtectedFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetNewInstancesProtectedFromScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetNotificationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetNotificationConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetServiceLinkedRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceLinkedRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetTargetGroupArNs() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArNs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetTerminationPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ResetVpcZoneIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcZoneIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

