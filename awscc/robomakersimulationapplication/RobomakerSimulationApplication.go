package robomakersimulationapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/robomakersimulationapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application awscc_robomaker_simulation_application}.
type RobomakerSimulationApplication interface {
	cdktf.TerraformResource
	Arn() *string
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
	CurrentRevisionId() *string
	SetCurrentRevisionId(val *string)
	CurrentRevisionIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
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
	RenderingEngine() RobomakerSimulationApplicationRenderingEngineOutputReference
	RenderingEngineInput() interface{}
	RobotSoftwareSuite() RobomakerSimulationApplicationRobotSoftwareSuiteOutputReference
	RobotSoftwareSuiteInput() interface{}
	SimulationSoftwareSuite() RobomakerSimulationApplicationSimulationSoftwareSuiteOutputReference
	SimulationSoftwareSuiteInput() interface{}
	Sources() RobomakerSimulationApplicationSourcesList
	SourcesInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	PutRenderingEngine(value *RobomakerSimulationApplicationRenderingEngine)
	PutRobotSoftwareSuite(value *RobomakerSimulationApplicationRobotSoftwareSuite)
	PutSimulationSoftwareSuite(value *RobomakerSimulationApplicationSimulationSoftwareSuite)
	PutSources(value interface{})
	ResetCurrentRevisionId()
	ResetEnvironment()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRenderingEngine()
	ResetSources()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RobomakerSimulationApplication
type jsiiProxy_RobomakerSimulationApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RobomakerSimulationApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) CurrentRevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentRevisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) CurrentRevisionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentRevisionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) RenderingEngine() RobomakerSimulationApplicationRenderingEngineOutputReference {
	var returns RobomakerSimulationApplicationRenderingEngineOutputReference
	_jsii_.Get(
		j,
		"renderingEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) RenderingEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderingEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) RobotSoftwareSuite() RobomakerSimulationApplicationRobotSoftwareSuiteOutputReference {
	var returns RobomakerSimulationApplicationRobotSoftwareSuiteOutputReference
	_jsii_.Get(
		j,
		"robotSoftwareSuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) RobotSoftwareSuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"robotSoftwareSuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) SimulationSoftwareSuite() RobomakerSimulationApplicationSimulationSoftwareSuiteOutputReference {
	var returns RobomakerSimulationApplicationSimulationSoftwareSuiteOutputReference
	_jsii_.Get(
		j,
		"simulationSoftwareSuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) SimulationSoftwareSuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"simulationSoftwareSuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Sources() RobomakerSimulationApplicationSourcesList {
	var returns RobomakerSimulationApplicationSourcesList
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) SourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RobomakerSimulationApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application awscc_robomaker_simulation_application} Resource.
func NewRobomakerSimulationApplication(scope constructs.Construct, id *string, config *RobomakerSimulationApplicationConfig) RobomakerSimulationApplication {
	_init_.Initialize()

	if err := validateNewRobomakerSimulationApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RobomakerSimulationApplication{}

	_jsii_.Create(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/robomaker_simulation_application awscc_robomaker_simulation_application} Resource.
func NewRobomakerSimulationApplication_Override(r RobomakerSimulationApplication, scope constructs.Construct, id *string, config *RobomakerSimulationApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetCurrentRevisionId(val *string) {
	if err := j.validateSetCurrentRevisionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentRevisionId",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RobomakerSimulationApplication)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a RobomakerSimulationApplication resource upon running "cdktf plan <stack-name>".
func RobomakerSimulationApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRobomakerSimulationApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
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
func RobomakerSimulationApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRobomakerSimulationApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RobomakerSimulationApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRobomakerSimulationApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RobomakerSimulationApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRobomakerSimulationApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RobomakerSimulationApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.robomakerSimulationApplication.RobomakerSimulationApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) PutRenderingEngine(value *RobomakerSimulationApplicationRenderingEngine) {
	if err := r.validatePutRenderingEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRenderingEngine",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) PutRobotSoftwareSuite(value *RobomakerSimulationApplicationRobotSoftwareSuite) {
	if err := r.validatePutRobotSoftwareSuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRobotSoftwareSuite",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) PutSimulationSoftwareSuite(value *RobomakerSimulationApplicationSimulationSoftwareSuite) {
	if err := r.validatePutSimulationSoftwareSuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSimulationSoftwareSuite",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) PutSources(value interface{}) {
	if err := r.validatePutSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSources",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetCurrentRevisionId() {
	_jsii_.InvokeVoid(
		r,
		"resetCurrentRevisionId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetEnvironment() {
	_jsii_.InvokeVoid(
		r,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetRenderingEngine() {
	_jsii_.InvokeVoid(
		r,
		"resetRenderingEngine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetSources() {
	_jsii_.InvokeVoid(
		r,
		"resetSources",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RobomakerSimulationApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RobomakerSimulationApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

