package applicationsignalsservicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationsignalsservicelevelobjective/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective awscc_applicationsignals_service_level_objective}.
type ApplicationsignalsServiceLevelObjective interface {
	cdktf.TerraformResource
	Arn() *string
	BurnRateConfigurations() ApplicationsignalsServiceLevelObjectiveBurnRateConfigurationsList
	BurnRateConfigurationsInput() interface{}
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
	CreatedTime() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EvaluationType() *string
	ExclusionWindows() ApplicationsignalsServiceLevelObjectiveExclusionWindowsList
	ExclusionWindowsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Goal() ApplicationsignalsServiceLevelObjectiveGoalOutputReference
	GoalInput() interface{}
	Id() *string
	LastUpdatedTime() *float64
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
	RequestBasedSli() ApplicationsignalsServiceLevelObjectiveRequestBasedSliOutputReference
	RequestBasedSliInput() interface{}
	Sli() ApplicationsignalsServiceLevelObjectiveSliOutputReference
	SliInput() interface{}
	Tags() ApplicationsignalsServiceLevelObjectiveTagsList
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
	PutBurnRateConfigurations(value interface{})
	PutExclusionWindows(value interface{})
	PutGoal(value *ApplicationsignalsServiceLevelObjectiveGoal)
	PutRequestBasedSli(value *ApplicationsignalsServiceLevelObjectiveRequestBasedSli)
	PutSli(value *ApplicationsignalsServiceLevelObjectiveSli)
	PutTags(value interface{})
	ResetBurnRateConfigurations()
	ResetDescription()
	ResetExclusionWindows()
	ResetGoal()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestBasedSli()
	ResetSli()
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

// The jsii proxy struct for ApplicationsignalsServiceLevelObjective
type jsiiProxy_ApplicationsignalsServiceLevelObjective struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) BurnRateConfigurations() ApplicationsignalsServiceLevelObjectiveBurnRateConfigurationsList {
	var returns ApplicationsignalsServiceLevelObjectiveBurnRateConfigurationsList
	_jsii_.Get(
		j,
		"burnRateConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) BurnRateConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"burnRateConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) CreatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) EvaluationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) ExclusionWindows() ApplicationsignalsServiceLevelObjectiveExclusionWindowsList {
	var returns ApplicationsignalsServiceLevelObjectiveExclusionWindowsList
	_jsii_.Get(
		j,
		"exclusionWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) ExclusionWindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Goal() ApplicationsignalsServiceLevelObjectiveGoalOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveGoalOutputReference
	_jsii_.Get(
		j,
		"goal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) GoalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"goalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) LastUpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) RequestBasedSli() ApplicationsignalsServiceLevelObjectiveRequestBasedSliOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveRequestBasedSliOutputReference
	_jsii_.Get(
		j,
		"requestBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) RequestBasedSliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBasedSliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Sli() ApplicationsignalsServiceLevelObjectiveSliOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveSliOutputReference
	_jsii_.Get(
		j,
		"sli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) SliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) Tags() ApplicationsignalsServiceLevelObjectiveTagsList {
	var returns ApplicationsignalsServiceLevelObjectiveTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective awscc_applicationsignals_service_level_objective} Resource.
func NewApplicationsignalsServiceLevelObjective(scope constructs.Construct, id *string, config *ApplicationsignalsServiceLevelObjectiveConfig) ApplicationsignalsServiceLevelObjective {
	_init_.Initialize()

	if err := validateNewApplicationsignalsServiceLevelObjectiveParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsServiceLevelObjective{}

	_jsii_.Create(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective awscc_applicationsignals_service_level_objective} Resource.
func NewApplicationsignalsServiceLevelObjective_Override(a ApplicationsignalsServiceLevelObjective, scope constructs.Construct, id *string, config *ApplicationsignalsServiceLevelObjectiveConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjective)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationsignalsServiceLevelObjective resource upon running "cdktf plan <stack-name>".
func ApplicationsignalsServiceLevelObjective_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationsignalsServiceLevelObjective_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
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
func ApplicationsignalsServiceLevelObjective_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationsignalsServiceLevelObjective_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationsignalsServiceLevelObjective_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationsignalsServiceLevelObjective_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationsignalsServiceLevelObjective_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationsignalsServiceLevelObjective_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationsignalsServiceLevelObjective_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjective",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutBurnRateConfigurations(value interface{}) {
	if err := a.validatePutBurnRateConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBurnRateConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutExclusionWindows(value interface{}) {
	if err := a.validatePutExclusionWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclusionWindows",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutGoal(value *ApplicationsignalsServiceLevelObjectiveGoal) {
	if err := a.validatePutGoalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutRequestBasedSli(value *ApplicationsignalsServiceLevelObjectiveRequestBasedSli) {
	if err := a.validatePutRequestBasedSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestBasedSli",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutSli(value *ApplicationsignalsServiceLevelObjectiveSli) {
	if err := a.validatePutSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSli",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetBurnRateConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetBurnRateConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetExclusionWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusionWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetGoal() {
	_jsii_.InvokeVoid(
		a,
		"resetGoal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetRequestBasedSli() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestBasedSli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetSli() {
	_jsii_.InvokeVoid(
		a,
		"resetSli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjective) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

