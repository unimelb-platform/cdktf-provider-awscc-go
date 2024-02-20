package healthlakefhirdatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/healthlakefhirdatastore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore awscc_healthlake_fhir_datastore}.
type HealthlakeFhirDatastore interface {
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
	CreatedAt() HealthlakeFhirDatastoreCreatedAtOutputReference
	DatastoreArn() *string
	DatastoreEndpoint() *string
	DatastoreId() *string
	DatastoreName() *string
	SetDatastoreName(val *string)
	DatastoreNameInput() *string
	DatastoreStatus() *string
	DatastoreTypeVersion() *string
	SetDatastoreTypeVersion(val *string)
	DatastoreTypeVersionInput() *string
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
	IdentityProviderConfiguration() HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference
	IdentityProviderConfigurationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PreloadDataConfig() HealthlakeFhirDatastorePreloadDataConfigOutputReference
	PreloadDataConfigInput() interface{}
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
	SseConfiguration() HealthlakeFhirDatastoreSseConfigurationOutputReference
	SseConfigurationInput() interface{}
	Tags() HealthlakeFhirDatastoreTagsList
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIdentityProviderConfiguration(value *HealthlakeFhirDatastoreIdentityProviderConfiguration)
	PutPreloadDataConfig(value *HealthlakeFhirDatastorePreloadDataConfig)
	PutSseConfiguration(value *HealthlakeFhirDatastoreSseConfiguration)
	PutTags(value interface{})
	ResetDatastoreName()
	ResetIdentityProviderConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreloadDataConfig()
	ResetSseConfiguration()
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

// The jsii proxy struct for HealthlakeFhirDatastore
type jsiiProxy_HealthlakeFhirDatastore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HealthlakeFhirDatastore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) CreatedAt() HealthlakeFhirDatastoreCreatedAtOutputReference {
	var returns HealthlakeFhirDatastoreCreatedAtOutputReference
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DatastoreTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) IdentityProviderConfiguration() HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference {
	var returns HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference
	_jsii_.Get(
		j,
		"identityProviderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) IdentityProviderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) PreloadDataConfig() HealthlakeFhirDatastorePreloadDataConfigOutputReference {
	var returns HealthlakeFhirDatastorePreloadDataConfigOutputReference
	_jsii_.Get(
		j,
		"preloadDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) PreloadDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) SseConfiguration() HealthlakeFhirDatastoreSseConfigurationOutputReference {
	var returns HealthlakeFhirDatastoreSseConfigurationOutputReference
	_jsii_.Get(
		j,
		"sseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) SseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) Tags() HealthlakeFhirDatastoreTagsList {
	var returns HealthlakeFhirDatastoreTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore awscc_healthlake_fhir_datastore} Resource.
func NewHealthlakeFhirDatastore(scope constructs.Construct, id *string, config *HealthlakeFhirDatastoreConfig) HealthlakeFhirDatastore {
	_init_.Initialize()

	if err := validateNewHealthlakeFhirDatastoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthlakeFhirDatastore{}

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore awscc_healthlake_fhir_datastore} Resource.
func NewHealthlakeFhirDatastore_Override(h HealthlakeFhirDatastore, scope constructs.Construct, id *string, config *HealthlakeFhirDatastoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetDatastoreName(val *string) {
	if err := j.validateSetDatastoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreName",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetDatastoreTypeVersion(val *string) {
	if err := j.validateSetDatastoreTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreTypeVersion",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a HealthlakeFhirDatastore resource upon running "cdktf plan <stack-name>".
func HealthlakeFhirDatastore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHealthlakeFhirDatastore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
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
func HealthlakeFhirDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthlakeFhirDatastore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthlakeFhirDatastore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthlakeFhirDatastore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthlakeFhirDatastore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthlakeFhirDatastore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HealthlakeFhirDatastore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) PutIdentityProviderConfiguration(value *HealthlakeFhirDatastoreIdentityProviderConfiguration) {
	if err := h.validatePutIdentityProviderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putIdentityProviderConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) PutPreloadDataConfig(value *HealthlakeFhirDatastorePreloadDataConfig) {
	if err := h.validatePutPreloadDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putPreloadDataConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) PutSseConfiguration(value *HealthlakeFhirDatastoreSseConfiguration) {
	if err := h.validatePutSseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putSseConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) PutTags(value interface{}) {
	if err := h.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTags",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetDatastoreName() {
	_jsii_.InvokeVoid(
		h,
		"resetDatastoreName",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetIdentityProviderConfiguration() {
	_jsii_.InvokeVoid(
		h,
		"resetIdentityProviderConfiguration",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetPreloadDataConfig() {
	_jsii_.InvokeVoid(
		h,
		"resetPreloadDataConfig",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetSseConfiguration() {
	_jsii_.InvokeVoid(
		h,
		"resetSseConfiguration",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

