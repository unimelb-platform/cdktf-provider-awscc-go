package quicksightanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksightanalysis/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis awscc_quicksight_analysis}.
type QuicksightAnalysis interface {
	cdktf.TerraformResource
	AnalysisId() *string
	SetAnalysisId(val *string)
	AnalysisIdInput() *string
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
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
	CreatedTime() *string
	DataSetArns() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Errors() QuicksightAnalysisErrorsList
	ErrorsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() QuicksightAnalysisParametersOutputReference
	ParametersInput() interface{}
	Permissions() QuicksightAnalysisPermissionsList
	PermissionsInput() interface{}
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
	Sheets() QuicksightAnalysisSheetsList
	SourceEntity() QuicksightAnalysisSourceEntityOutputReference
	SourceEntityInput() interface{}
	Status() *string
	Tags() QuicksightAnalysisTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThemeArn() *string
	SetThemeArn(val *string)
	ThemeArnInput() *string
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
	PutErrors(value interface{})
	PutParameters(value *QuicksightAnalysisParameters)
	PutPermissions(value interface{})
	PutSourceEntity(value *QuicksightAnalysisSourceEntity)
	PutTags(value interface{})
	ResetErrors()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPermissions()
	ResetTags()
	ResetThemeArn()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for QuicksightAnalysis
type jsiiProxy_QuicksightAnalysis struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightAnalysis) AnalysisId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) AnalysisIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) DataSetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Errors() QuicksightAnalysisErrorsList {
	var returns QuicksightAnalysisErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Parameters() QuicksightAnalysisParametersOutputReference {
	var returns QuicksightAnalysisParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Permissions() QuicksightAnalysisPermissionsList {
	var returns QuicksightAnalysisPermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) PermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Sheets() QuicksightAnalysisSheetsList {
	var returns QuicksightAnalysisSheetsList
	_jsii_.Get(
		j,
		"sheets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) SourceEntity() QuicksightAnalysisSourceEntityOutputReference {
	var returns QuicksightAnalysisSourceEntityOutputReference
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) SourceEntityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) Tags() QuicksightAnalysisTagsList {
	var returns QuicksightAnalysisTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ThemeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightAnalysis) ThemeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis awscc_quicksight_analysis} Resource.
func NewQuicksightAnalysis(scope constructs.Construct, id *string, config *QuicksightAnalysisConfig) QuicksightAnalysis {
	_init_.Initialize()

	if err := validateNewQuicksightAnalysisParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightAnalysis{}

	_jsii_.Create(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis awscc_quicksight_analysis} Resource.
func NewQuicksightAnalysis_Override(q QuicksightAnalysis, scope constructs.Construct, id *string, config *QuicksightAnalysisConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetAnalysisId(val *string) {
	if err := j.validateSetAnalysisIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisId",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QuicksightAnalysis)SetThemeArn(val *string) {
	if err := j.validateSetThemeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"themeArn",
		val,
	)
}

// Generates CDKTF code for importing a QuicksightAnalysis resource upon running "cdktf plan <stack-name>".
func QuicksightAnalysis_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQuicksightAnalysis_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
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
func QuicksightAnalysis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAnalysis_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightAnalysis_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAnalysis_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightAnalysis_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightAnalysis_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightAnalysis_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.quicksightAnalysis.QuicksightAnalysis",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) PutErrors(value interface{}) {
	if err := q.validatePutErrorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putErrors",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) PutParameters(value *QuicksightAnalysisParameters) {
	if err := q.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) PutPermissions(value interface{}) {
	if err := q.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPermissions",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) PutSourceEntity(value *QuicksightAnalysisSourceEntity) {
	if err := q.validatePutSourceEntityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSourceEntity",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetErrors() {
	_jsii_.InvokeVoid(
		q,
		"resetErrors",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetName() {
	_jsii_.InvokeVoid(
		q,
		"resetName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetPermissions() {
	_jsii_.InvokeVoid(
		q,
		"resetPermissions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) ResetThemeArn() {
	_jsii_.InvokeVoid(
		q,
		"resetThemeArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightAnalysis) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightAnalysis) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

