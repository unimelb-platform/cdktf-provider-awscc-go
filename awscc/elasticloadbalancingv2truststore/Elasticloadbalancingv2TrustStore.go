package elasticloadbalancingv2truststore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2truststore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store}.
type Elasticloadbalancingv2TrustStore interface {
	cdktf.TerraformResource
	CaCertificatesBundleS3Bucket() *string
	SetCaCertificatesBundleS3Bucket(val *string)
	CaCertificatesBundleS3BucketInput() *string
	CaCertificatesBundleS3Key() *string
	SetCaCertificatesBundleS3Key(val *string)
	CaCertificatesBundleS3KeyInput() *string
	CaCertificatesBundleS3ObjectVersion() *string
	SetCaCertificatesBundleS3ObjectVersion(val *string)
	CaCertificatesBundleS3ObjectVersionInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumberOfCaCertificates() *float64
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
	Status() *string
	Tags() Elasticloadbalancingv2TrustStoreTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustStoreArn() *string
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
	PutTags(value interface{})
	ResetCaCertificatesBundleS3Bucket()
	ResetCaCertificatesBundleS3Key()
	ResetCaCertificatesBundleS3ObjectVersion()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for Elasticloadbalancingv2TrustStore
type jsiiProxy_Elasticloadbalancingv2TrustStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) NumberOfCaCertificates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCaCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Tags() Elasticloadbalancingv2TrustStoreTagsList {
	var returns Elasticloadbalancingv2TrustStoreTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store} Resource.
func NewElasticloadbalancingv2TrustStore(scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreConfig) Elasticloadbalancingv2TrustStore {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2TrustStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2TrustStore{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store} Resource.
func NewElasticloadbalancingv2TrustStore_Override(e Elasticloadbalancingv2TrustStore, scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3Bucket(val *string) {
	if err := j.validateSetCaCertificatesBundleS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3Bucket",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3Key(val *string) {
	if err := j.validateSetCaCertificatesBundleS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3Key",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3ObjectVersion(val *string) {
	if err := j.validateSetCaCertificatesBundleS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Elasticloadbalancingv2TrustStore resource upon running "cdktf plan <stack-name>".
func Elasticloadbalancingv2TrustStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
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
func Elasticloadbalancingv2TrustStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Elasticloadbalancingv2TrustStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3Bucket() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3Bucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3Key() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3Key",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3ObjectVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3ObjectVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

