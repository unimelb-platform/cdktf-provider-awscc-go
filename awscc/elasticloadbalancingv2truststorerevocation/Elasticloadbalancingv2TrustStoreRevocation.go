package elasticloadbalancingv2truststorerevocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2truststorerevocation/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store_revocation awscc_elasticloadbalancingv2_trust_store_revocation}.
type Elasticloadbalancingv2TrustStoreRevocation interface {
	cdktf.TerraformResource
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
	RevocationContents() Elasticloadbalancingv2TrustStoreRevocationRevocationContentsList
	RevocationContentsInput() interface{}
	RevocationId() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustStoreArn() *string
	SetTrustStoreArn(val *string)
	TrustStoreArnInput() *string
	TrustStoreRevocations() Elasticloadbalancingv2TrustStoreRevocationTrustStoreRevocationsList
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
	PutRevocationContents(value interface{})
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRevocationContents()
	ResetTrustStoreArn()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Elasticloadbalancingv2TrustStoreRevocation
type jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) RevocationContents() Elasticloadbalancingv2TrustStoreRevocationRevocationContentsList {
	var returns Elasticloadbalancingv2TrustStoreRevocationRevocationContentsList
	_jsii_.Get(
		j,
		"revocationContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) RevocationContentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revocationContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) RevocationId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TrustStoreArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) TrustStoreRevocations() Elasticloadbalancingv2TrustStoreRevocationTrustStoreRevocationsList {
	var returns Elasticloadbalancingv2TrustStoreRevocationTrustStoreRevocationsList
	_jsii_.Get(
		j,
		"trustStoreRevocations",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store_revocation awscc_elasticloadbalancingv2_trust_store_revocation} Resource.
func NewElasticloadbalancingv2TrustStoreRevocation(scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreRevocationConfig) Elasticloadbalancingv2TrustStoreRevocation {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2TrustStoreRevocationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store_revocation awscc_elasticloadbalancingv2_trust_store_revocation} Resource.
func NewElasticloadbalancingv2TrustStoreRevocation_Override(e Elasticloadbalancingv2TrustStoreRevocation, scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreRevocationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation)SetTrustStoreArn(val *string) {
	if err := j.validateSetTrustStoreArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreArn",
		val,
	)
}

// Generates CDKTF code for importing a Elasticloadbalancingv2TrustStoreRevocation resource upon running "cdktf plan <stack-name>".
func Elasticloadbalancingv2TrustStoreRevocation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStoreRevocation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
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
func Elasticloadbalancingv2TrustStoreRevocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStoreRevocation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStoreRevocation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStoreRevocation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStoreRevocation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStoreRevocation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Elasticloadbalancingv2TrustStoreRevocation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) PutRevocationContents(value interface{}) {
	if err := e.validatePutRevocationContentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRevocationContents",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ResetRevocationContents() {
	_jsii_.InvokeVoid(
		e,
		"resetRevocationContents",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ResetTrustStoreArn() {
	_jsii_.InvokeVoid(
		e,
		"resetTrustStoreArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

