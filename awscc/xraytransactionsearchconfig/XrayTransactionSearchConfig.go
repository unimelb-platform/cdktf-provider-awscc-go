package xraytransactionsearchconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/xraytransactionsearchconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/xray_transaction_search_config awscc_xray_transaction_search_config}.
type XrayTransactionSearchConfig interface {
	cdktf.TerraformResource
	AccountId() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IndexingPercentage() *float64
	SetIndexingPercentage(val *float64)
	IndexingPercentageInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ResetIndexingPercentage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for XrayTransactionSearchConfig
type jsiiProxy_XrayTransactionSearchConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_XrayTransactionSearchConfig) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) IndexingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) IndexingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayTransactionSearchConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/xray_transaction_search_config awscc_xray_transaction_search_config} Resource.
func NewXrayTransactionSearchConfig(scope constructs.Construct, id *string, config *XrayTransactionSearchConfigConfig) XrayTransactionSearchConfig {
	_init_.Initialize()

	if err := validateNewXrayTransactionSearchConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_XrayTransactionSearchConfig{}

	_jsii_.Create(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/xray_transaction_search_config awscc_xray_transaction_search_config} Resource.
func NewXrayTransactionSearchConfig_Override(x XrayTransactionSearchConfig, scope constructs.Construct, id *string, config *XrayTransactionSearchConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		[]interface{}{scope, id, config},
		x,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetIndexingPercentage(val *float64) {
	if err := j.validateSetIndexingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexingPercentage",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_XrayTransactionSearchConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a XrayTransactionSearchConfig resource upon running "cdktf plan <stack-name>".
func XrayTransactionSearchConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateXrayTransactionSearchConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
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
func XrayTransactionSearchConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayTransactionSearchConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XrayTransactionSearchConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayTransactionSearchConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XrayTransactionSearchConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayTransactionSearchConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func XrayTransactionSearchConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.xrayTransactionSearchConfig.XrayTransactionSearchConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) AddMoveTarget(moveTarget *string) {
	if err := x.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) AddOverride(path *string, value interface{}) {
	if err := x.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := x.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := x.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		x,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := x.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		x,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := x.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		x,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := x.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		x,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := x.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		x,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := x.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := x.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		x,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := x.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) MoveFromId(id *string) {
	if err := x.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveFromId",
		[]interface{}{id},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := x.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) MoveToId(id *string) {
	if err := x.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveToId",
		[]interface{}{id},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) OverrideLogicalId(newLogicalId *string) {
	if err := x.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ResetIndexingPercentage() {
	_jsii_.InvokeVoid(
		x,
		"resetIndexingPercentage",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		x,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayTransactionSearchConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayTransactionSearchConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

