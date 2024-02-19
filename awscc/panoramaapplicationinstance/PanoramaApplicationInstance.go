package panoramaapplicationinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/panoramaapplicationinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance awscc_panorama_application_instance}.
type PanoramaApplicationInstance interface {
	cdktf.TerraformResource
	ApplicationInstanceId() *string
	ApplicationInstanceIdToReplace() *string
	SetApplicationInstanceIdToReplace(val *string)
	ApplicationInstanceIdToReplaceInput() *string
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
	CreatedTime() *float64
	DefaultRuntimeContextDevice() *string
	SetDefaultRuntimeContextDevice(val *string)
	DefaultRuntimeContextDeviceInput() *string
	DefaultRuntimeContextDeviceName() *string
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
	HealthStatus() *string
	Id() *string
	LastUpdatedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManifestOverridesPayload() PanoramaApplicationInstanceManifestOverridesPayloadOutputReference
	ManifestOverridesPayloadInput() interface{}
	ManifestPayload() PanoramaApplicationInstanceManifestPayloadOutputReference
	ManifestPayloadInput() interface{}
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
	RuntimeRoleArn() *string
	SetRuntimeRoleArn(val *string)
	RuntimeRoleArnInput() *string
	Status() *string
	StatusDescription() *string
	Tags() PanoramaApplicationInstanceTagsList
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
	PutManifestOverridesPayload(value *PanoramaApplicationInstanceManifestOverridesPayload)
	PutManifestPayload(value *PanoramaApplicationInstanceManifestPayload)
	PutTags(value interface{})
	ResetApplicationInstanceIdToReplace()
	ResetDescription()
	ResetManifestOverridesPayload()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuntimeRoleArn()
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

// The jsii proxy struct for PanoramaApplicationInstance
type jsiiProxy_PanoramaApplicationInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PanoramaApplicationInstance) ApplicationInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ApplicationInstanceIdToReplace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInstanceIdToReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ApplicationInstanceIdToReplaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInstanceIdToReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) CreatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) DefaultRuntimeContextDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRuntimeContextDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) DefaultRuntimeContextDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRuntimeContextDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) DefaultRuntimeContextDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRuntimeContextDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) HealthStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) LastUpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ManifestOverridesPayload() PanoramaApplicationInstanceManifestOverridesPayloadOutputReference {
	var returns PanoramaApplicationInstanceManifestOverridesPayloadOutputReference
	_jsii_.Get(
		j,
		"manifestOverridesPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ManifestOverridesPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifestOverridesPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ManifestPayload() PanoramaApplicationInstanceManifestPayloadOutputReference {
	var returns PanoramaApplicationInstanceManifestPayloadOutputReference
	_jsii_.Get(
		j,
		"manifestPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) ManifestPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifestPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) RuntimeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) RuntimeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) StatusDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) Tags() PanoramaApplicationInstanceTagsList {
	var returns PanoramaApplicationInstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaApplicationInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance awscc_panorama_application_instance} Resource.
func NewPanoramaApplicationInstance(scope constructs.Construct, id *string, config *PanoramaApplicationInstanceConfig) PanoramaApplicationInstance {
	_init_.Initialize()

	if err := validateNewPanoramaApplicationInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PanoramaApplicationInstance{}

	_jsii_.Create(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance awscc_panorama_application_instance} Resource.
func NewPanoramaApplicationInstance_Override(p PanoramaApplicationInstance, scope constructs.Construct, id *string, config *PanoramaApplicationInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetApplicationInstanceIdToReplace(val *string) {
	if err := j.validateSetApplicationInstanceIdToReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationInstanceIdToReplace",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetDefaultRuntimeContextDevice(val *string) {
	if err := j.validateSetDefaultRuntimeContextDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRuntimeContextDevice",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PanoramaApplicationInstance)SetRuntimeRoleArn(val *string) {
	if err := j.validateSetRuntimeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeRoleArn",
		val,
	)
}

// Generates CDKTF code for importing a PanoramaApplicationInstance resource upon running "cdktf plan <stack-name>".
func PanoramaApplicationInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePanoramaApplicationInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
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
func PanoramaApplicationInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaApplicationInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PanoramaApplicationInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaApplicationInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PanoramaApplicationInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaApplicationInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PanoramaApplicationInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.panoramaApplicationInstance.PanoramaApplicationInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) PutManifestOverridesPayload(value *PanoramaApplicationInstanceManifestOverridesPayload) {
	if err := p.validatePutManifestOverridesPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManifestOverridesPayload",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) PutManifestPayload(value *PanoramaApplicationInstanceManifestPayload) {
	if err := p.validatePutManifestPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManifestPayload",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) PutTags(value interface{}) {
	if err := p.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetApplicationInstanceIdToReplace() {
	_jsii_.InvokeVoid(
		p,
		"resetApplicationInstanceIdToReplace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetManifestOverridesPayload() {
	_jsii_.InvokeVoid(
		p,
		"resetManifestOverridesPayload",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetRuntimeRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetRuntimeRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaApplicationInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaApplicationInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

