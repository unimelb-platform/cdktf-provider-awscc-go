package appsyncsourceapiassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appsyncsourceapiassociation/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association awscc_appsync_source_api_association}.
type AppsyncSourceApiAssociation interface {
	cdktf.TerraformResource
	AssociationArn() *string
	AssociationId() *string
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
	LastSuccessfulMergeDate() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergedApiArn() *string
	MergedApiId() *string
	MergedApiIdentifier() *string
	SetMergedApiIdentifier(val *string)
	MergedApiIdentifierInput() *string
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
	SourceApiArn() *string
	SourceApiAssociationConfig() AppsyncSourceApiAssociationSourceApiAssociationConfigOutputReference
	SourceApiAssociationConfigInput() interface{}
	SourceApiAssociationStatus() *string
	SourceApiAssociationStatusDetail() *string
	SourceApiId() *string
	SourceApiIdentifier() *string
	SetSourceApiIdentifier(val *string)
	SourceApiIdentifierInput() *string
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
	PutSourceApiAssociationConfig(value *AppsyncSourceApiAssociationSourceApiAssociationConfig)
	ResetDescription()
	ResetMergedApiIdentifier()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceApiAssociationConfig()
	ResetSourceApiIdentifier()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncSourceApiAssociation
type jsiiProxy_AppsyncSourceApiAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) AssociationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) LastSuccessfulMergeDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSuccessfulMergeDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) MergedApiArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) MergedApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) MergedApiIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) MergedApiIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiAssociationConfig() AppsyncSourceApiAssociationSourceApiAssociationConfigOutputReference {
	var returns AppsyncSourceApiAssociationSourceApiAssociationConfigOutputReference
	_jsii_.Get(
		j,
		"sourceApiAssociationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiAssociationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceApiAssociationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiAssociationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiAssociationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiAssociationStatusDetail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiAssociationStatusDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) SourceApiIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncSourceApiAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association awscc_appsync_source_api_association} Resource.
func NewAppsyncSourceApiAssociation(scope constructs.Construct, id *string, config *AppsyncSourceApiAssociationConfig) AppsyncSourceApiAssociation {
	_init_.Initialize()

	if err := validateNewAppsyncSourceApiAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncSourceApiAssociation{}

	_jsii_.Create(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association awscc_appsync_source_api_association} Resource.
func NewAppsyncSourceApiAssociation_Override(a AppsyncSourceApiAssociation, scope constructs.Construct, id *string, config *AppsyncSourceApiAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetMergedApiIdentifier(val *string) {
	if err := j.validateSetMergedApiIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergedApiIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncSourceApiAssociation)SetSourceApiIdentifier(val *string) {
	if err := j.validateSetSourceApiIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceApiIdentifier",
		val,
	)
}

// Generates CDKTF code for importing a AppsyncSourceApiAssociation resource upon running "cdktf plan <stack-name>".
func AppsyncSourceApiAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncSourceApiAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
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
func AppsyncSourceApiAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncSourceApiAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncSourceApiAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncSourceApiAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncSourceApiAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncSourceApiAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncSourceApiAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.appsyncSourceApiAssociation.AppsyncSourceApiAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncSourceApiAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) PutSourceApiAssociationConfig(value *AppsyncSourceApiAssociationSourceApiAssociationConfig) {
	if err := a.validatePutSourceApiAssociationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceApiAssociationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ResetMergedApiIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetMergedApiIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ResetSourceApiAssociationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceApiAssociationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ResetSourceApiIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceApiIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncSourceApiAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

