package gameliftcontainergroupdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftcontainergroupdefinition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_group_definition awscc_gamelift_container_group_definition}.
type GameliftContainerGroupDefinition interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerGroupDefinitionArn() *string
	ContainerGroupType() *string
	SetContainerGroupType(val *string)
	ContainerGroupTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
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
	GameServerContainerDefinition() GameliftContainerGroupDefinitionGameServerContainerDefinitionOutputReference
	GameServerContainerDefinitionInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
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
	SourceVersionNumber() *float64
	SetSourceVersionNumber(val *float64)
	SourceVersionNumberInput() *float64
	Status() *string
	StatusReason() *string
	SupportContainerDefinitions() GameliftContainerGroupDefinitionSupportContainerDefinitionsList
	SupportContainerDefinitionsInput() interface{}
	Tags() GameliftContainerGroupDefinitionTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TotalMemoryLimitMebibytes() *float64
	SetTotalMemoryLimitMebibytes(val *float64)
	TotalMemoryLimitMebibytesInput() *float64
	TotalVcpuLimit() *float64
	SetTotalVcpuLimit(val *float64)
	TotalVcpuLimitInput() *float64
	VersionDescription() *string
	SetVersionDescription(val *string)
	VersionDescriptionInput() *string
	VersionNumber() *float64
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
	PutGameServerContainerDefinition(value *GameliftContainerGroupDefinitionGameServerContainerDefinition)
	PutSupportContainerDefinitions(value interface{})
	PutTags(value interface{})
	ResetContainerGroupType()
	ResetGameServerContainerDefinition()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceVersionNumber()
	ResetSupportContainerDefinitions()
	ResetTags()
	ResetVersionDescription()
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

// The jsii proxy struct for GameliftContainerGroupDefinition
type jsiiProxy_GameliftContainerGroupDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) ContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) ContainerGroupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerGroupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) ContainerGroupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerGroupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) GameServerContainerDefinition() GameliftContainerGroupDefinitionGameServerContainerDefinitionOutputReference {
	var returns GameliftContainerGroupDefinitionGameServerContainerDefinitionOutputReference
	_jsii_.Get(
		j,
		"gameServerContainerDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) GameServerContainerDefinitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameServerContainerDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) SourceVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) SourceVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) SupportContainerDefinitions() GameliftContainerGroupDefinitionSupportContainerDefinitionsList {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsList
	_jsii_.Get(
		j,
		"supportContainerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) SupportContainerDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportContainerDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) Tags() GameliftContainerGroupDefinitionTagsList {
	var returns GameliftContainerGroupDefinitionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TotalMemoryLimitMebibytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalMemoryLimitMebibytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TotalMemoryLimitMebibytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalMemoryLimitMebibytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TotalVcpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalVcpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) TotalVcpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalVcpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) VersionDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinition) VersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_group_definition awscc_gamelift_container_group_definition} Resource.
func NewGameliftContainerGroupDefinition(scope constructs.Construct, id *string, config *GameliftContainerGroupDefinitionConfig) GameliftContainerGroupDefinition {
	_init_.Initialize()

	if err := validateNewGameliftContainerGroupDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerGroupDefinition{}

	_jsii_.Create(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_group_definition awscc_gamelift_container_group_definition} Resource.
func NewGameliftContainerGroupDefinition_Override(g GameliftContainerGroupDefinition, scope constructs.Construct, id *string, config *GameliftContainerGroupDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetContainerGroupType(val *string) {
	if err := j.validateSetContainerGroupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerGroupType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetSourceVersionNumber(val *float64) {
	if err := j.validateSetSourceVersionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVersionNumber",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetTotalMemoryLimitMebibytes(val *float64) {
	if err := j.validateSetTotalMemoryLimitMebibytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalMemoryLimitMebibytes",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetTotalVcpuLimit(val *float64) {
	if err := j.validateSetTotalVcpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalVcpuLimit",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinition)SetVersionDescription(val *string) {
	if err := j.validateSetVersionDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

// Generates CDKTF code for importing a GameliftContainerGroupDefinition resource upon running "cdktf plan <stack-name>".
func GameliftContainerGroupDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftContainerGroupDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
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
func GameliftContainerGroupDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerGroupDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerGroupDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerGroupDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerGroupDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerGroupDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftContainerGroupDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GameliftContainerGroupDefinition) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) PutGameServerContainerDefinition(value *GameliftContainerGroupDefinitionGameServerContainerDefinition) {
	if err := g.validatePutGameServerContainerDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGameServerContainerDefinition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) PutSupportContainerDefinitions(value interface{}) {
	if err := g.validatePutSupportContainerDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSupportContainerDefinitions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetContainerGroupType() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerGroupType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetGameServerContainerDefinition() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerContainerDefinition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetSourceVersionNumber() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceVersionNumber",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetSupportContainerDefinitions() {
	_jsii_.InvokeVoid(
		g,
		"resetSupportContainerDefinitions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ResetVersionDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetVersionDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

