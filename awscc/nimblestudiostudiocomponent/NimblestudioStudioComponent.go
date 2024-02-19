package nimblestudiostudiocomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/nimblestudiostudiocomponent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component awscc_nimblestudio_studio_component}.
type NimblestudioStudioComponent interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Configuration() NimblestudioStudioComponentConfigurationOutputReference
	ConfigurationInput() interface{}
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
	Ec2SecurityGroupIds() *[]*string
	SetEc2SecurityGroupIds(val *[]*string)
	Ec2SecurityGroupIdsInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InitializationScripts() NimblestudioStudioComponentInitializationScriptsList
	InitializationScriptsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RuntimeRoleArn() *string
	SetRuntimeRoleArn(val *string)
	RuntimeRoleArnInput() *string
	ScriptParameters() NimblestudioStudioComponentScriptParametersList
	ScriptParametersInput() interface{}
	SecureInitializationRoleArn() *string
	SetSecureInitializationRoleArn(val *string)
	SecureInitializationRoleArnInput() *string
	StudioComponentId() *string
	StudioId() *string
	SetStudioId(val *string)
	StudioIdInput() *string
	Subtype() *string
	SetSubtype(val *string)
	SubtypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutConfiguration(value *NimblestudioStudioComponentConfiguration)
	PutInitializationScripts(value interface{})
	PutScriptParameters(value interface{})
	ResetConfiguration()
	ResetDescription()
	ResetEc2SecurityGroupIds()
	ResetInitializationScripts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuntimeRoleArn()
	ResetScriptParameters()
	ResetSecureInitializationRoleArn()
	ResetSubtype()
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

// The jsii proxy struct for NimblestudioStudioComponent
type jsiiProxy_NimblestudioStudioComponent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NimblestudioStudioComponent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Configuration() NimblestudioStudioComponentConfigurationOutputReference {
	var returns NimblestudioStudioComponentConfigurationOutputReference
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Ec2SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2SecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Ec2SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2SecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) InitializationScripts() NimblestudioStudioComponentInitializationScriptsList {
	var returns NimblestudioStudioComponentInitializationScriptsList
	_jsii_.Get(
		j,
		"initializationScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) InitializationScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) RuntimeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) RuntimeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) ScriptParameters() NimblestudioStudioComponentScriptParametersList {
	var returns NimblestudioStudioComponentScriptParametersList
	_jsii_.Get(
		j,
		"scriptParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) ScriptParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) SecureInitializationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureInitializationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) SecureInitializationRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureInitializationRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) StudioComponentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioComponentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) StudioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Subtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) SubtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioStudioComponent) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component awscc_nimblestudio_studio_component} Resource.
func NewNimblestudioStudioComponent(scope constructs.Construct, id *string, config *NimblestudioStudioComponentConfig) NimblestudioStudioComponent {
	_init_.Initialize()

	if err := validateNewNimblestudioStudioComponentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NimblestudioStudioComponent{}

	_jsii_.Create(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component awscc_nimblestudio_studio_component} Resource.
func NewNimblestudioStudioComponent_Override(n NimblestudioStudioComponent, scope constructs.Construct, id *string, config *NimblestudioStudioComponentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetEc2SecurityGroupIds(val *[]*string) {
	if err := j.validateSetEc2SecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2SecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetRuntimeRoleArn(val *string) {
	if err := j.validateSetRuntimeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeRoleArn",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetSecureInitializationRoleArn(val *string) {
	if err := j.validateSetSecureInitializationRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureInitializationRoleArn",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetStudioId(val *string) {
	if err := j.validateSetStudioIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetSubtype(val *string) {
	if err := j.validateSetSubtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtype",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NimblestudioStudioComponent)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a NimblestudioStudioComponent resource upon running "cdktf plan <stack-name>".
func NimblestudioStudioComponent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNimblestudioStudioComponent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
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
func NimblestudioStudioComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNimblestudioStudioComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NimblestudioStudioComponent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNimblestudioStudioComponent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NimblestudioStudioComponent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNimblestudioStudioComponent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NimblestudioStudioComponent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.nimblestudioStudioComponent.NimblestudioStudioComponent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) PutConfiguration(value *NimblestudioStudioComponentConfiguration) {
	if err := n.validatePutConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) PutInitializationScripts(value interface{}) {
	if err := n.validatePutInitializationScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putInitializationScripts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) PutScriptParameters(value interface{}) {
	if err := n.validatePutScriptParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putScriptParameters",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetEc2SecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetEc2SecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetInitializationScripts() {
	_jsii_.InvokeVoid(
		n,
		"resetInitializationScripts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetRuntimeRoleArn() {
	_jsii_.InvokeVoid(
		n,
		"resetRuntimeRoleArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetScriptParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetScriptParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetSecureInitializationRoleArn() {
	_jsii_.InvokeVoid(
		n,
		"resetSecureInitializationRoleArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetSubtype() {
	_jsii_.InvokeVoid(
		n,
		"resetSubtype",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioStudioComponent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioStudioComponent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

