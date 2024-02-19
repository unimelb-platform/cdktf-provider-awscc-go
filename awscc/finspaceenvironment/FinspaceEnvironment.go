package finspaceenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/finspaceenvironment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/finspace_environment awscc_finspace_environment}.
type FinspaceEnvironment interface {
	cdktf.TerraformResource
	AwsAccountId() *string
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
	DataBundles() *[]*string
	SetDataBundles(val *[]*string)
	DataBundlesInput() *[]*string
	DedicatedServiceAccountId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentArn() *string
	EnvironmentId() *string
	EnvironmentUrl() *string
	FederationMode() *string
	SetFederationMode(val *string)
	FederationModeInput() *string
	FederationParameters() FinspaceEnvironmentFederationParametersOutputReference
	FederationParametersInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
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
	SageMakerStudioDomainUrl() *string
	Status() *string
	SuperuserParameters() FinspaceEnvironmentSuperuserParametersOutputReference
	SuperuserParametersInput() interface{}
	Tags() FinspaceEnvironmentTagsList
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
	PutFederationParameters(value *FinspaceEnvironmentFederationParameters)
	PutSuperuserParameters(value *FinspaceEnvironmentSuperuserParameters)
	PutTags(value interface{})
	ResetDataBundles()
	ResetDescription()
	ResetFederationMode()
	ResetFederationParameters()
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSuperuserParameters()
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

// The jsii proxy struct for FinspaceEnvironment
type jsiiProxy_FinspaceEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FinspaceEnvironment) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) DataBundles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataBundles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) DataBundlesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataBundlesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) DedicatedServiceAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedServiceAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) EnvironmentUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) FederationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) FederationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) FederationParameters() FinspaceEnvironmentFederationParametersOutputReference {
	var returns FinspaceEnvironmentFederationParametersOutputReference
	_jsii_.Get(
		j,
		"federationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) FederationParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"federationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) SageMakerStudioDomainUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sageMakerStudioDomainUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) SuperuserParameters() FinspaceEnvironmentSuperuserParametersOutputReference {
	var returns FinspaceEnvironmentSuperuserParametersOutputReference
	_jsii_.Get(
		j,
		"superuserParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) SuperuserParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"superuserParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) Tags() FinspaceEnvironmentTagsList {
	var returns FinspaceEnvironmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/finspace_environment awscc_finspace_environment} Resource.
func NewFinspaceEnvironment(scope constructs.Construct, id *string, config *FinspaceEnvironmentConfig) FinspaceEnvironment {
	_init_.Initialize()

	if err := validateNewFinspaceEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FinspaceEnvironment{}

	_jsii_.Create(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/finspace_environment awscc_finspace_environment} Resource.
func NewFinspaceEnvironment_Override(f FinspaceEnvironment, scope constructs.Construct, id *string, config *FinspaceEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetDataBundles(val *[]*string) {
	if err := j.validateSetDataBundlesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataBundles",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetFederationMode(val *string) {
	if err := j.validateSetFederationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federationMode",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a FinspaceEnvironment resource upon running "cdktf plan <stack-name>".
func FinspaceEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFinspaceEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
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
func FinspaceEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FinspaceEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FinspaceEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFinspaceEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FinspaceEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.finspaceEnvironment.FinspaceEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) PutFederationParameters(value *FinspaceEnvironmentFederationParameters) {
	if err := f.validatePutFederationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFederationParameters",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) PutSuperuserParameters(value *FinspaceEnvironmentSuperuserParameters) {
	if err := f.validatePutSuperuserParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSuperuserParameters",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) PutTags(value interface{}) {
	if err := f.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetDataBundles() {
	_jsii_.InvokeVoid(
		f,
		"resetDataBundles",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetFederationMode() {
	_jsii_.InvokeVoid(
		f,
		"resetFederationMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetFederationParameters() {
	_jsii_.InvokeVoid(
		f,
		"resetFederationParameters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetSuperuserParameters() {
	_jsii_.InvokeVoid(
		f,
		"resetSuperuserParameters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

