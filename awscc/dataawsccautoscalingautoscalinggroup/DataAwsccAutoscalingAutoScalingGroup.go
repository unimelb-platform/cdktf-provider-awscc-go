package dataawsccautoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccautoscalingautoscalinggroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group}.
type DataAwsccAutoscalingAutoScalingGroup interface {
	cdktf.TerraformDataSource
	AutoScalingGroupName() *string
	AvailabilityZones() *[]*string
	CapacityRebalance() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	Cooldown() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultInstanceWarmup() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCapacity() *string
	DesiredCapacityType() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriod() *float64
	HealthCheckType() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	InstanceMaintenancePolicy() DataAwsccAutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference
	LaunchConfigurationName() *string
	LaunchTemplate() DataAwsccAutoscalingAutoScalingGroupLaunchTemplateOutputReference
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleHookSpecificationList() DataAwsccAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList
	LoadBalancerNames() *[]*string
	MaxInstanceLifetime() *float64
	MaxSize() *string
	MetricsCollection() DataAwsccAutoscalingAutoScalingGroupMetricsCollectionList
	MinSize() *string
	MixedInstancesPolicy() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference
	NewInstancesProtectedFromScaleIn() cdktf.IResolvable
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationOutputReference
	NotificationConfigurations() DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationsList
	PlacementGroup() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceLinkedRoleArn() *string
	Tags() DataAwsccAutoscalingAutoScalingGroupTagsList
	TargetGroupArNs() *[]*string
	TerminationPolicies() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcZoneIdentifier() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsccAutoscalingAutoScalingGroup
type jsiiProxy_DataAwsccAutoscalingAutoScalingGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) CapacityRebalance() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) DesiredCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) InstanceMaintenancePolicy() DataAwsccAutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) LaunchTemplate() DataAwsccAutoscalingAutoScalingGroupLaunchTemplateOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) LifecycleHookSpecificationList() DataAwsccAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList {
	var returns DataAwsccAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructList
	_jsii_.Get(
		j,
		"lifecycleHookSpecificationList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) LoadBalancerNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) MaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) MetricsCollection() DataAwsccAutoscalingAutoScalingGroupMetricsCollectionList {
	var returns DataAwsccAutoscalingAutoScalingGroupMetricsCollectionList
	_jsii_.Get(
		j,
		"metricsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) MinSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) MixedInstancesPolicy() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyOutputReference
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) NewInstancesProtectedFromScaleIn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) NotificationConfiguration() DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) NotificationConfigurations() DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationsList {
	var returns DataAwsccAutoscalingAutoScalingGroupNotificationConfigurationsList
	_jsii_.Get(
		j,
		"notificationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) Tags() DataAwsccAutoscalingAutoScalingGroupTagsList {
	var returns DataAwsccAutoscalingAutoScalingGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) TargetGroupArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group} Data Source.
func NewDataAwsccAutoscalingAutoScalingGroup(scope constructs.Construct, id *string, config *DataAwsccAutoscalingAutoScalingGroupConfig) DataAwsccAutoscalingAutoScalingGroup {
	_init_.Initialize()

	if err := validateNewDataAwsccAutoscalingAutoScalingGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAutoscalingAutoScalingGroup{}

	_jsii_.Create(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/autoscaling_auto_scaling_group awscc_autoscaling_auto_scaling_group} Data Source.
func NewDataAwsccAutoscalingAutoScalingGroup_Override(d DataAwsccAutoscalingAutoScalingGroup, scope constructs.Construct, id *string, config *DataAwsccAutoscalingAutoScalingGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccAutoscalingAutoScalingGroup resource upon running "cdktf plan <stack-name>".
func DataAwsccAutoscalingAutoScalingGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccAutoscalingAutoScalingGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
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
func DataAwsccAutoscalingAutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAutoscalingAutoScalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccAutoscalingAutoScalingGroup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAutoscalingAutoScalingGroup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccAutoscalingAutoScalingGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAutoscalingAutoScalingGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccAutoscalingAutoScalingGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

