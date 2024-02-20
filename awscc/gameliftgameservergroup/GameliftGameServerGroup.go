package gameliftgameservergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftgameservergroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_server_group awscc_gamelift_game_server_group}.
type GameliftGameServerGroup interface {
	cdktf.TerraformResource
	AutoScalingGroupArn() *string
	AutoScalingPolicy() GameliftGameServerGroupAutoScalingPolicyOutputReference
	AutoScalingPolicyInput() interface{}
	BalancingStrategy() *string
	SetBalancingStrategy(val *string)
	BalancingStrategyInput() *string
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
	DeleteOption() *string
	SetDeleteOption(val *string)
	DeleteOptionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameServerGroupArn() *string
	GameServerGroupName() *string
	SetGameServerGroupName(val *string)
	GameServerGroupNameInput() *string
	GameServerProtectionPolicy() *string
	SetGameServerProtectionPolicy(val *string)
	GameServerProtectionPolicyInput() *string
	Id() *string
	InstanceDefinitions() GameliftGameServerGroupInstanceDefinitionsList
	InstanceDefinitionsInput() interface{}
	LaunchTemplate() GameliftGameServerGroupLaunchTemplateOutputReference
	LaunchTemplateInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	// The tree node.
	Node() constructs.Node
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() GameliftGameServerGroupTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcSubnets() *[]*string
	SetVpcSubnets(val *[]*string)
	VpcSubnetsInput() *[]*string
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
	PutAutoScalingPolicy(value *GameliftGameServerGroupAutoScalingPolicy)
	PutInstanceDefinitions(value interface{})
	PutLaunchTemplate(value *GameliftGameServerGroupLaunchTemplate)
	PutTags(value interface{})
	ResetAutoScalingPolicy()
	ResetBalancingStrategy()
	ResetDeleteOption()
	ResetGameServerProtectionPolicy()
	ResetLaunchTemplate()
	ResetMaxSize()
	ResetMinSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetVpcSubnets()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GameliftGameServerGroup
type jsiiProxy_GameliftGameServerGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftGameServerGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) AutoScalingPolicy() GameliftGameServerGroupAutoScalingPolicyOutputReference {
	var returns GameliftGameServerGroupAutoScalingPolicyOutputReference
	_jsii_.Get(
		j,
		"autoScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) AutoScalingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) BalancingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) BalancingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) DeleteOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) DeleteOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) GameServerGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) GameServerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) GameServerGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) GameServerProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) GameServerProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) InstanceDefinitions() GameliftGameServerGroupInstanceDefinitionsList {
	var returns GameliftGameServerGroupInstanceDefinitionsList
	_jsii_.Get(
		j,
		"instanceDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) InstanceDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) LaunchTemplate() GameliftGameServerGroupLaunchTemplateOutputReference {
	var returns GameliftGameServerGroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) LaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) Tags() GameliftGameServerGroupTagsList {
	var returns GameliftGameServerGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) VpcSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftGameServerGroup) VpcSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_server_group awscc_gamelift_game_server_group} Resource.
func NewGameliftGameServerGroup(scope constructs.Construct, id *string, config *GameliftGameServerGroupConfig) GameliftGameServerGroup {
	_init_.Initialize()

	if err := validateNewGameliftGameServerGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftGameServerGroup{}

	_jsii_.Create(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_server_group awscc_gamelift_game_server_group} Resource.
func NewGameliftGameServerGroup_Override(g GameliftGameServerGroup, scope constructs.Construct, id *string, config *GameliftGameServerGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetBalancingStrategy(val *string) {
	if err := j.validateSetBalancingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balancingStrategy",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetDeleteOption(val *string) {
	if err := j.validateSetDeleteOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOption",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetGameServerGroupName(val *string) {
	if err := j.validateSetGameServerGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerGroupName",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetGameServerProtectionPolicy(val *string) {
	if err := j.validateSetGameServerProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftGameServerGroup)SetVpcSubnets(val *[]*string) {
	if err := j.validateSetVpcSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSubnets",
		val,
	)
}

// Generates CDKTF code for importing a GameliftGameServerGroup resource upon running "cdktf plan <stack-name>".
func GameliftGameServerGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftGameServerGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
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
func GameliftGameServerGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameServerGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftGameServerGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameServerGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftGameServerGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftGameServerGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftGameServerGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftGameServerGroup.GameliftGameServerGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftGameServerGroup) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftGameServerGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftGameServerGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftGameServerGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) PutAutoScalingPolicy(value *GameliftGameServerGroupAutoScalingPolicy) {
	if err := g.validatePutAutoScalingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoScalingPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) PutInstanceDefinitions(value interface{}) {
	if err := g.validatePutInstanceDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceDefinitions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) PutLaunchTemplate(value *GameliftGameServerGroupLaunchTemplate) {
	if err := g.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetAutoScalingPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoScalingPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetBalancingStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetBalancingStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetDeleteOption() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetGameServerProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetMaxSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetMinSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMinSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) ResetVpcSubnets() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcSubnets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftGameServerGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameServerGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameServerGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftGameServerGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

