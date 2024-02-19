package evidentlyexperiment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/evidentlyexperiment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment awscc_evidently_experiment}.
type EvidentlyExperiment interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	MetricGoals() EvidentlyExperimentMetricGoalsList
	MetricGoalsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnlineAbConfig() EvidentlyExperimentOnlineAbConfigOutputReference
	OnlineAbConfigInput() interface{}
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RandomizationSalt() *string
	SetRandomizationSalt(val *string)
	RandomizationSaltInput() *string
	// Experimental.
	RawOverrides() interface{}
	RemoveSegment() interface{}
	SetRemoveSegment(val interface{})
	RemoveSegmentInput() interface{}
	RunningStatus() EvidentlyExperimentRunningStatusOutputReference
	RunningStatusInput() interface{}
	SamplingRate() *float64
	SetSamplingRate(val *float64)
	SamplingRateInput() *float64
	Segment() *string
	SetSegment(val *string)
	SegmentInput() *string
	Tags() EvidentlyExperimentTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Treatments() EvidentlyExperimentTreatmentsList
	TreatmentsInput() interface{}
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
	PutMetricGoals(value interface{})
	PutOnlineAbConfig(value *EvidentlyExperimentOnlineAbConfig)
	PutRunningStatus(value *EvidentlyExperimentRunningStatus)
	PutTags(value interface{})
	PutTreatments(value interface{})
	ResetDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRandomizationSalt()
	ResetRemoveSegment()
	ResetRunningStatus()
	ResetSamplingRate()
	ResetSegment()
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

// The jsii proxy struct for EvidentlyExperiment
type jsiiProxy_EvidentlyExperiment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EvidentlyExperiment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) MetricGoals() EvidentlyExperimentMetricGoalsList {
	var returns EvidentlyExperimentMetricGoalsList
	_jsii_.Get(
		j,
		"metricGoals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) MetricGoalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGoalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) OnlineAbConfig() EvidentlyExperimentOnlineAbConfigOutputReference {
	var returns EvidentlyExperimentOnlineAbConfigOutputReference
	_jsii_.Get(
		j,
		"onlineAbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) OnlineAbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlineAbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RandomizationSalt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomizationSalt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RandomizationSaltInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomizationSaltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RemoveSegment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeSegment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RemoveSegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeSegmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RunningStatus() EvidentlyExperimentRunningStatusOutputReference {
	var returns EvidentlyExperimentRunningStatusOutputReference
	_jsii_.Get(
		j,
		"runningStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) RunningStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) SamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) SamplingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Segment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) SegmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Tags() EvidentlyExperimentTagsList {
	var returns EvidentlyExperimentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) Treatments() EvidentlyExperimentTreatmentsList {
	var returns EvidentlyExperimentTreatmentsList
	_jsii_.Get(
		j,
		"treatments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperiment) TreatmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatmentsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment awscc_evidently_experiment} Resource.
func NewEvidentlyExperiment(scope constructs.Construct, id *string, config *EvidentlyExperimentConfig) EvidentlyExperiment {
	_init_.Initialize()

	if err := validateNewEvidentlyExperimentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvidentlyExperiment{}

	_jsii_.Create(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment awscc_evidently_experiment} Resource.
func NewEvidentlyExperiment_Override(e EvidentlyExperiment, scope constructs.Construct, id *string, config *EvidentlyExperimentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetRandomizationSalt(val *string) {
	if err := j.validateSetRandomizationSaltParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomizationSalt",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetRemoveSegment(val interface{}) {
	if err := j.validateSetRemoveSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeSegment",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetSamplingRate(val *float64) {
	if err := j.validateSetSamplingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperiment)SetSegment(val *string) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

// Generates CDKTF code for importing a EvidentlyExperiment resource upon running "cdktf plan <stack-name>".
func EvidentlyExperiment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEvidentlyExperiment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
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
func EvidentlyExperiment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvidentlyExperiment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EvidentlyExperiment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvidentlyExperiment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EvidentlyExperiment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvidentlyExperiment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EvidentlyExperiment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.evidentlyExperiment.EvidentlyExperiment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) PutMetricGoals(value interface{}) {
	if err := e.validatePutMetricGoalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetricGoals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) PutOnlineAbConfig(value *EvidentlyExperimentOnlineAbConfig) {
	if err := e.validatePutOnlineAbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnlineAbConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) PutRunningStatus(value *EvidentlyExperimentRunningStatus) {
	if err := e.validatePutRunningStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRunningStatus",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) PutTreatments(value interface{}) {
	if err := e.validatePutTreatmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTreatments",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetRandomizationSalt() {
	_jsii_.InvokeVoid(
		e,
		"resetRandomizationSalt",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetRemoveSegment() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoveSegment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetRunningStatus() {
	_jsii_.InvokeVoid(
		e,
		"resetRunningStatus",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetSamplingRate() {
	_jsii_.InvokeVoid(
		e,
		"resetSamplingRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetSegment() {
	_jsii_.InvokeVoid(
		e,
		"resetSegment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperiment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperiment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

