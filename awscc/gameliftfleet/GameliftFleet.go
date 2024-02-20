package gameliftfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftfleet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet awscc_gamelift_fleet}.
type GameliftFleet interface {
	cdktf.TerraformResource
	AnywhereConfiguration() GameliftFleetAnywhereConfigurationOutputReference
	AnywhereConfigurationInput() interface{}
	ApplyCapacity() *string
	SetApplyCapacity(val *string)
	ApplyCapacityInput() *string
	BuildId() *string
	SetBuildId(val *string)
	BuildIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference
	CertificateConfigurationInput() interface{}
	ComputeType() *string
	SetComputeType(val *string)
	ComputeTypeInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredEc2Instances() *float64
	SetDesiredEc2Instances(val *float64)
	DesiredEc2InstancesInput() *float64
	Ec2InboundPermissions() GameliftFleetEc2InboundPermissionsList
	Ec2InboundPermissionsInput() interface{}
	Ec2InstanceType() *string
	SetEc2InstanceType(val *string)
	Ec2InstanceTypeInput() *string
	FleetId() *string
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	InstanceRoleArnInput() *string
	InstanceRoleCredentialsProvider() *string
	SetInstanceRoleCredentialsProvider(val *string)
	InstanceRoleCredentialsProviderInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locations() GameliftFleetLocationsList
	LocationsInput() interface{}
	LogPaths() *[]*string
	SetLogPaths(val *[]*string)
	LogPathsInput() *[]*string
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MetricGroupsInput() *[]*string
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	NewGameSessionProtectionPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	PeerVpcAwsAccountId() *string
	SetPeerVpcAwsAccountId(val *string)
	PeerVpcAwsAccountIdInput() *string
	PeerVpcId() *string
	SetPeerVpcId(val *string)
	PeerVpcIdInput() *string
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
	ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference
	ResourceCreationLimitPolicyInput() interface{}
	RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference
	RuntimeConfigurationInput() interface{}
	ScalingPolicies() GameliftFleetScalingPoliciesList
	ScalingPoliciesInput() interface{}
	ScriptId() *string
	SetScriptId(val *string)
	ScriptIdInput() *string
	ServerLaunchParameters() *string
	SetServerLaunchParameters(val *string)
	ServerLaunchParametersInput() *string
	ServerLaunchPath() *string
	SetServerLaunchPath(val *string)
	ServerLaunchPathInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAnywhereConfiguration(value *GameliftFleetAnywhereConfiguration)
	PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration)
	PutEc2InboundPermissions(value interface{})
	PutLocations(value interface{})
	PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy)
	PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration)
	PutScalingPolicies(value interface{})
	ResetAnywhereConfiguration()
	ResetApplyCapacity()
	ResetBuildId()
	ResetCertificateConfiguration()
	ResetComputeType()
	ResetDescription()
	ResetDesiredEc2Instances()
	ResetEc2InboundPermissions()
	ResetEc2InstanceType()
	ResetFleetType()
	ResetInstanceRoleArn()
	ResetInstanceRoleCredentialsProvider()
	ResetLocations()
	ResetLogPaths()
	ResetMaxSize()
	ResetMetricGroups()
	ResetMinSize()
	ResetNewGameSessionProtectionPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerVpcAwsAccountId()
	ResetPeerVpcId()
	ResetResourceCreationLimitPolicy()
	ResetRuntimeConfiguration()
	ResetScalingPolicies()
	ResetScriptId()
	ResetServerLaunchParameters()
	ResetServerLaunchPath()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GameliftFleet
type jsiiProxy_GameliftFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftFleet) AnywhereConfiguration() GameliftFleetAnywhereConfigurationOutputReference {
	var returns GameliftFleetAnywhereConfigurationOutputReference
	_jsii_.Get(
		j,
		"anywhereConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) AnywhereConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anywhereConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ApplyCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ApplyCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference {
	var returns GameliftFleetCertificateConfigurationOutputReference
	_jsii_.Get(
		j,
		"certificateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DesiredEc2Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredEc2Instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DesiredEc2InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredEc2InstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermissions() GameliftFleetEc2InboundPermissionsList {
	var returns GameliftFleetEc2InboundPermissionsList
	_jsii_.Get(
		j,
		"ec2InboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InboundPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleCredentialsProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleCredentialsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleCredentialsProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleCredentialsProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Locations() GameliftFleetLocationsList {
	var returns GameliftFleetLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) LocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) LogPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) LogPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) PeerVpcAwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcAwsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) PeerVpcAwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcAwsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) PeerVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) PeerVpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference {
	var returns GameliftFleetResourceCreationLimitPolicyOutputReference
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference {
	var returns GameliftFleetRuntimeConfigurationOutputReference
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScalingPolicies() GameliftFleetScalingPoliciesList {
	var returns GameliftFleetScalingPoliciesList
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ServerLaunchParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ServerLaunchParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ServerLaunchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ServerLaunchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet awscc_gamelift_fleet} Resource.
func NewGameliftFleet(scope constructs.Construct, id *string, config *GameliftFleetConfig) GameliftFleet {
	_init_.Initialize()

	if err := validateNewGameliftFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftFleet{}

	_jsii_.Create(
		"awscc.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet awscc_gamelift_fleet} Resource.
func NewGameliftFleet_Override(g GameliftFleet, scope constructs.Construct, id *string, config *GameliftFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftFleet)SetApplyCapacity(val *string) {
	if err := j.validateSetApplyCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyCapacity",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetBuildId(val *string) {
	if err := j.validateSetBuildIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetComputeType(val *string) {
	if err := j.validateSetComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDesiredEc2Instances(val *float64) {
	if err := j.validateSetDesiredEc2InstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredEc2Instances",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetEc2InstanceType(val *string) {
	if err := j.validateSetEc2InstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetInstanceRoleArn(val *string) {
	if err := j.validateSetInstanceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetInstanceRoleCredentialsProvider(val *string) {
	if err := j.validateSetInstanceRoleCredentialsProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRoleCredentialsProvider",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetLogPaths(val *[]*string) {
	if err := j.validateSetLogPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logPaths",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetMetricGroups(val *[]*string) {
	if err := j.validateSetMetricGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetNewGameSessionProtectionPolicy(val *string) {
	if err := j.validateSetNewGameSessionProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetPeerVpcAwsAccountId(val *string) {
	if err := j.validateSetPeerVpcAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVpcAwsAccountId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetPeerVpcId(val *string) {
	if err := j.validateSetPeerVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVpcId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetScriptId(val *string) {
	if err := j.validateSetScriptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetServerLaunchParameters(val *string) {
	if err := j.validateSetServerLaunchParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverLaunchParameters",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetServerLaunchPath(val *string) {
	if err := j.validateSetServerLaunchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverLaunchPath",
		val,
	)
}

// Generates CDKTF code for importing a GameliftFleet resource upon running "cdktf plan <stack-name>".
func GameliftFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftFleet.GameliftFleet",
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
func GameliftFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftFleet.GameliftFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftFleet.GameliftFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftFleet.GameliftFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftFleet.GameliftFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftFleet) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftFleet) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftFleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftFleet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftFleet) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftFleet) PutAnywhereConfiguration(value *GameliftFleetAnywhereConfiguration) {
	if err := g.validatePutAnywhereConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnywhereConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration) {
	if err := g.validatePutCertificateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutEc2InboundPermissions(value interface{}) {
	if err := g.validatePutEc2InboundPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEc2InboundPermissions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutLocations(value interface{}) {
	if err := g.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy) {
	if err := g.validatePutResourceCreationLimitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceCreationLimitPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration) {
	if err := g.validatePutRuntimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutScalingPolicies(value interface{}) {
	if err := g.validatePutScalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) ResetAnywhereConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetAnywhereConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetApplyCapacity() {
	_jsii_.InvokeVoid(
		g,
		"resetApplyCapacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetBuildId() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetCertificateConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetComputeType() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetDesiredEc2Instances() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredEc2Instances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetEc2InboundPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetEc2InboundPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetEc2InstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetEc2InstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		g,
		"resetFleetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetInstanceRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetInstanceRoleCredentialsProvider() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRoleCredentialsProvider",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetLogPaths() {
	_jsii_.InvokeVoid(
		g,
		"resetLogPaths",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetMaxSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetMetricGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetMinSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMinSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetNewGameSessionProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNewGameSessionProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetPeerVpcAwsAccountId() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerVpcAwsAccountId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetPeerVpcId() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerVpcId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetResourceCreationLimitPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceCreationLimitPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetRuntimeConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetScalingPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetScriptId() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetServerLaunchParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetServerLaunchParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetServerLaunchPath() {
	_jsii_.InvokeVoid(
		g,
		"resetServerLaunchPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

