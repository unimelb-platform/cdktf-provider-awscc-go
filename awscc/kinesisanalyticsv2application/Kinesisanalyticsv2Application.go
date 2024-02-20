package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisanalyticsv2application/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application awscc_kinesisanalyticsv2_application}.
type Kinesisanalyticsv2Application interface {
	cdktf.TerraformResource
	ApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
	ApplicationConfigurationInput() interface{}
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	ApplicationDescriptionInput() *string
	ApplicationMaintenanceConfiguration() Kinesisanalyticsv2ApplicationApplicationMaintenanceConfigurationOutputReference
	ApplicationMaintenanceConfigurationInput() interface{}
	ApplicationMode() *string
	SetApplicationMode(val *string)
	ApplicationModeInput() *string
	ApplicationName() *string
	SetApplicationName(val *string)
	ApplicationNameInput() *string
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
	RunConfiguration() Kinesisanalyticsv2ApplicationRunConfigurationOutputReference
	RunConfigurationInput() interface{}
	RuntimeEnvironment() *string
	SetRuntimeEnvironment(val *string)
	RuntimeEnvironmentInput() *string
	ServiceExecutionRole() *string
	SetServiceExecutionRole(val *string)
	ServiceExecutionRoleInput() *string
	Tags() Kinesisanalyticsv2ApplicationTagsList
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
	PutApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfiguration)
	PutApplicationMaintenanceConfiguration(value *Kinesisanalyticsv2ApplicationApplicationMaintenanceConfiguration)
	PutRunConfiguration(value *Kinesisanalyticsv2ApplicationRunConfiguration)
	PutTags(value interface{})
	ResetApplicationConfiguration()
	ResetApplicationDescription()
	ResetApplicationMaintenanceConfiguration()
	ResetApplicationMode()
	ResetApplicationName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunConfiguration()
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

// The jsii proxy struct for Kinesisanalyticsv2Application
type jsiiProxy_Kinesisanalyticsv2Application struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationMaintenanceConfiguration() Kinesisanalyticsv2ApplicationApplicationMaintenanceConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationMaintenanceConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationMaintenanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationMaintenanceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationMaintenanceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RunConfiguration() Kinesisanalyticsv2ApplicationRunConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationRunConfigurationOutputReference
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RunConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RuntimeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RuntimeEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ServiceExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ServiceExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Tags() Kinesisanalyticsv2ApplicationTagsList {
	var returns Kinesisanalyticsv2ApplicationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application awscc_kinesisanalyticsv2_application} Resource.
func NewKinesisanalyticsv2Application(scope constructs.Construct, id *string, config *Kinesisanalyticsv2ApplicationConfig) Kinesisanalyticsv2Application {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2Application{}

	_jsii_.Create(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application awscc_kinesisanalyticsv2_application} Resource.
func NewKinesisanalyticsv2Application_Override(k Kinesisanalyticsv2Application, scope constructs.Construct, id *string, config *Kinesisanalyticsv2ApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetApplicationDescription(val *string) {
	if err := j.validateSetApplicationDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationDescription",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetApplicationMode(val *string) {
	if err := j.validateSetApplicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationMode",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetApplicationName(val *string) {
	if err := j.validateSetApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetRuntimeEnvironment(val *string) {
	if err := j.validateSetRuntimeEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironment",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application)SetServiceExecutionRole(val *string) {
	if err := j.validateSetServiceExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRole",
		val,
	)
}

// Generates CDKTF code for importing a Kinesisanalyticsv2Application resource upon running "cdktf plan <stack-name>".
func Kinesisanalyticsv2Application_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
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
func Kinesisanalyticsv2Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Kinesisanalyticsv2Application_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Kinesisanalyticsv2Application_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Kinesisanalyticsv2Application_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfiguration) {
	if err := k.validatePutApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutApplicationMaintenanceConfiguration(value *Kinesisanalyticsv2ApplicationApplicationMaintenanceConfiguration) {
	if err := k.validatePutApplicationMaintenanceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationMaintenanceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutRunConfiguration(value *Kinesisanalyticsv2ApplicationRunConfiguration) {
	if err := k.validatePutRunConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRunConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutTags(value interface{}) {
	if err := k.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTags",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationMaintenanceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationMaintenanceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationMode() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationName() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetRunConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetRunConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

