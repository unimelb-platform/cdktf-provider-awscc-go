package gameliftcontainerfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftcontainerfleet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet}.
type GameliftContainerFleet interface {
	cdktf.TerraformResource
	BillingType() *string
	SetBillingType(val *string)
	BillingTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentConfiguration() GameliftContainerFleetDeploymentConfigurationOutputReference
	DeploymentConfigurationInput() interface{}
	DeploymentDetails() GameliftContainerFleetDeploymentDetailsOutputReference
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FleetArn() *string
	FleetId() *string
	FleetRoleArn() *string
	SetFleetRoleArn(val *string)
	FleetRoleArnInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameServerContainerGroupDefinitionArn() *string
	GameServerContainerGroupDefinitionName() *string
	SetGameServerContainerGroupDefinitionName(val *string)
	GameServerContainerGroupDefinitionNameInput() *string
	GameServerContainerGroupsPerInstance() *float64
	SetGameServerContainerGroupsPerInstance(val *float64)
	GameServerContainerGroupsPerInstanceInput() *float64
	GameSessionCreationLimitPolicy() GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference
	GameSessionCreationLimitPolicyInput() interface{}
	Id() *string
	InstanceConnectionPortRange() GameliftContainerFleetInstanceConnectionPortRangeOutputReference
	InstanceConnectionPortRangeInput() interface{}
	InstanceInboundPermissions() GameliftContainerFleetInstanceInboundPermissionsList
	InstanceInboundPermissionsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locations() GameliftContainerFleetLocationsList
	LocationsInput() interface{}
	LogConfiguration() GameliftContainerFleetLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	MaximumGameServerContainerGroupsPerInstance() *float64
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MetricGroupsInput() *[]*string
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	NewGameSessionProtectionPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	PerInstanceContainerGroupDefinitionArn() *string
	PerInstanceContainerGroupDefinitionName() *string
	SetPerInstanceContainerGroupDefinitionName(val *string)
	PerInstanceContainerGroupDefinitionNameInput() *string
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
	ScalingPolicies() GameliftContainerFleetScalingPoliciesList
	ScalingPoliciesInput() interface{}
	Status() *string
	Tags() GameliftContainerFleetTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDeploymentConfiguration(value *GameliftContainerFleetDeploymentConfiguration)
	PutGameSessionCreationLimitPolicy(value *GameliftContainerFleetGameSessionCreationLimitPolicy)
	PutInstanceConnectionPortRange(value *GameliftContainerFleetInstanceConnectionPortRange)
	PutInstanceInboundPermissions(value interface{})
	PutLocations(value interface{})
	PutLogConfiguration(value *GameliftContainerFleetLogConfiguration)
	PutScalingPolicies(value interface{})
	PutTags(value interface{})
	ResetBillingType()
	ResetDeploymentConfiguration()
	ResetDescription()
	ResetGameServerContainerGroupDefinitionName()
	ResetGameServerContainerGroupsPerInstance()
	ResetGameSessionCreationLimitPolicy()
	ResetInstanceConnectionPortRange()
	ResetInstanceInboundPermissions()
	ResetInstanceType()
	ResetLocations()
	ResetLogConfiguration()
	ResetMetricGroups()
	ResetNewGameSessionProtectionPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerInstanceContainerGroupDefinitionName()
	ResetScalingPolicies()
	ResetTags()
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

// The jsii proxy struct for GameliftContainerFleet
type jsiiProxy_GameliftContainerFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftContainerFleet) BillingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) BillingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentConfiguration() GameliftContainerFleetDeploymentConfigurationOutputReference {
	var returns GameliftContainerFleetDeploymentConfigurationOutputReference
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentDetails() GameliftContainerFleetDeploymentDetailsOutputReference {
	var returns GameliftContainerFleetDeploymentDetailsOutputReference
	_jsii_.Get(
		j,
		"deploymentDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupsPerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameServerContainerGroupsPerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameSessionCreationLimitPolicy() GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference {
	var returns GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference
	_jsii_.Get(
		j,
		"gameSessionCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameSessionCreationLimitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameSessionCreationLimitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceConnectionPortRange() GameliftContainerFleetInstanceConnectionPortRangeOutputReference {
	var returns GameliftContainerFleetInstanceConnectionPortRangeOutputReference
	_jsii_.Get(
		j,
		"instanceConnectionPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceConnectionPortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConnectionPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceInboundPermissions() GameliftContainerFleetInstanceInboundPermissionsList {
	var returns GameliftContainerFleetInstanceInboundPermissionsList
	_jsii_.Get(
		j,
		"instanceInboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceInboundPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceInboundPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Locations() GameliftContainerFleetLocationsList {
	var returns GameliftContainerFleetLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LogConfiguration() GameliftContainerFleetLogConfigurationOutputReference {
	var returns GameliftContainerFleetLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MaximumGameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumGameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MetricGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) NewGameSessionProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ScalingPolicies() GameliftContainerFleetScalingPoliciesList {
	var returns GameliftContainerFleetScalingPoliciesList
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ScalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Tags() GameliftContainerFleetTagsList {
	var returns GameliftContainerFleetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet} Resource.
func NewGameliftContainerFleet(scope constructs.Construct, id *string, config *GameliftContainerFleetConfig) GameliftContainerFleet {
	_init_.Initialize()

	if err := validateNewGameliftContainerFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerFleet{}

	_jsii_.Create(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet} Resource.
func NewGameliftContainerFleet_Override(g GameliftContainerFleet, scope constructs.Construct, id *string, config *GameliftContainerFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetBillingType(val *string) {
	if err := j.validateSetBillingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetFleetRoleArn(val *string) {
	if err := j.validateSetFleetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetRoleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetGameServerContainerGroupDefinitionName(val *string) {
	if err := j.validateSetGameServerContainerGroupDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetGameServerContainerGroupsPerInstance(val *float64) {
	if err := j.validateSetGameServerContainerGroupsPerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerContainerGroupsPerInstance",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetMetricGroups(val *[]*string) {
	if err := j.validateSetMetricGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetNewGameSessionProtectionPolicy(val *string) {
	if err := j.validateSetNewGameSessionProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetPerInstanceContainerGroupDefinitionName(val *string) {
	if err := j.validateSetPerInstanceContainerGroupDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perInstanceContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GameliftContainerFleet resource upon running "cdktf plan <stack-name>".
func GameliftContainerFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
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
func GameliftContainerFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftContainerFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftContainerFleet.GameliftContainerFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutDeploymentConfiguration(value *GameliftContainerFleetDeploymentConfiguration) {
	if err := g.validatePutDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutGameSessionCreationLimitPolicy(value *GameliftContainerFleetGameSessionCreationLimitPolicy) {
	if err := g.validatePutGameSessionCreationLimitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGameSessionCreationLimitPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutInstanceConnectionPortRange(value *GameliftContainerFleetInstanceConnectionPortRange) {
	if err := g.validatePutInstanceConnectionPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceConnectionPortRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutInstanceInboundPermissions(value interface{}) {
	if err := g.validatePutInstanceInboundPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceInboundPermissions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutLocations(value interface{}) {
	if err := g.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutLogConfiguration(value *GameliftContainerFleetLogConfiguration) {
	if err := g.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutScalingPolicies(value interface{}) {
	if err := g.validatePutScalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetBillingType() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameServerContainerGroupDefinitionName() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerContainerGroupDefinitionName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameServerContainerGroupsPerInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerContainerGroupsPerInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameSessionCreationLimitPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetGameSessionCreationLimitPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceConnectionPortRange() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceConnectionPortRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceInboundPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceInboundPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetMetricGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetNewGameSessionProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNewGameSessionProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetPerInstanceContainerGroupDefinitionName() {
	_jsii_.InvokeVoid(
		g,
		"resetPerInstanceContainerGroupDefinitionName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetScalingPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

