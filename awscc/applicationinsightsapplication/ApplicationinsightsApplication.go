package applicationinsightsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationinsightsapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application awscc_applicationinsights_application}.
type ApplicationinsightsApplication interface {
	cdktf.TerraformResource
	ApplicationArn() *string
	AutoConfigurationEnabled() interface{}
	SetAutoConfigurationEnabled(val interface{})
	AutoConfigurationEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComponentMonitoringSettings() ApplicationinsightsApplicationComponentMonitoringSettingsList
	ComponentMonitoringSettingsInput() interface{}
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
	CustomComponents() ApplicationinsightsApplicationCustomComponentsList
	CustomComponentsInput() interface{}
	CweMonitorEnabled() interface{}
	SetCweMonitorEnabled(val interface{})
	CweMonitorEnabledInput() interface{}
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
	GroupingType() *string
	SetGroupingType(val *string)
	GroupingTypeInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPatternSets() ApplicationinsightsApplicationLogPatternSetsList
	LogPatternSetsInput() interface{}
	// The tree node.
	Node() constructs.Node
	OpsCenterEnabled() interface{}
	SetOpsCenterEnabled(val interface{})
	OpsCenterEnabledInput() interface{}
	OpsItemSnsTopicArn() *string
	SetOpsItemSnsTopicArn(val *string)
	OpsItemSnsTopicArnInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() ApplicationinsightsApplicationTagsList
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
	PutComponentMonitoringSettings(value interface{})
	PutCustomComponents(value interface{})
	PutLogPatternSets(value interface{})
	PutTags(value interface{})
	ResetAutoConfigurationEnabled()
	ResetComponentMonitoringSettings()
	ResetCustomComponents()
	ResetCweMonitorEnabled()
	ResetGroupingType()
	ResetLogPatternSets()
	ResetOpsCenterEnabled()
	ResetOpsItemSnsTopicArn()
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

// The jsii proxy struct for ApplicationinsightsApplication
type jsiiProxy_ApplicationinsightsApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationinsightsApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) AutoConfigurationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoConfigurationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) AutoConfigurationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoConfigurationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ComponentMonitoringSettings() ApplicationinsightsApplicationComponentMonitoringSettingsList {
	var returns ApplicationinsightsApplicationComponentMonitoringSettingsList
	_jsii_.Get(
		j,
		"componentMonitoringSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ComponentMonitoringSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentMonitoringSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) CustomComponents() ApplicationinsightsApplicationCustomComponentsList {
	var returns ApplicationinsightsApplicationCustomComponentsList
	_jsii_.Get(
		j,
		"customComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) CustomComponentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customComponentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) CweMonitorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cweMonitorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) CweMonitorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cweMonitorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) GroupingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) GroupingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) LogPatternSets() ApplicationinsightsApplicationLogPatternSetsList {
	var returns ApplicationinsightsApplicationLogPatternSetsList
	_jsii_.Get(
		j,
		"logPatternSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) LogPatternSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPatternSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) OpsCenterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsCenterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) OpsCenterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsCenterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) OpsItemSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsItemSnsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) OpsItemSnsTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsItemSnsTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) Tags() ApplicationinsightsApplicationTagsList {
	var returns ApplicationinsightsApplicationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application awscc_applicationinsights_application} Resource.
func NewApplicationinsightsApplication(scope constructs.Construct, id *string, config *ApplicationinsightsApplicationConfig) ApplicationinsightsApplication {
	_init_.Initialize()

	if err := validateNewApplicationinsightsApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationinsightsApplication{}

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application awscc_applicationinsights_application} Resource.
func NewApplicationinsightsApplication_Override(a ApplicationinsightsApplication, scope constructs.Construct, id *string, config *ApplicationinsightsApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetAutoConfigurationEnabled(val interface{}) {
	if err := j.validateSetAutoConfigurationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetCweMonitorEnabled(val interface{}) {
	if err := j.validateSetCweMonitorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cweMonitorEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetGroupingType(val *string) {
	if err := j.validateSetGroupingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingType",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetOpsCenterEnabled(val interface{}) {
	if err := j.validateSetOpsCenterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsCenterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetOpsItemSnsTopicArn(val *string) {
	if err := j.validateSetOpsItemSnsTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsItemSnsTopicArn",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplication)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationinsightsApplication resource upon running "cdktf plan <stack-name>".
func ApplicationinsightsApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationinsightsApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
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
func ApplicationinsightsApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationinsightsApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationinsightsApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationinsightsApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationinsightsApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationinsightsApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationinsightsApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationinsightsApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplication) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) PutComponentMonitoringSettings(value interface{}) {
	if err := a.validatePutComponentMonitoringSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComponentMonitoringSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) PutCustomComponents(value interface{}) {
	if err := a.validatePutCustomComponentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomComponents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) PutLogPatternSets(value interface{}) {
	if err := a.validatePutLogPatternSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogPatternSets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetAutoConfigurationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoConfigurationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetComponentMonitoringSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetComponentMonitoringSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetCustomComponents() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomComponents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetCweMonitorEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCweMonitorEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetGroupingType() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetLogPatternSets() {
	_jsii_.InvokeVoid(
		a,
		"resetLogPatternSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetOpsCenterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsCenterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetOpsItemSnsTopicArn() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsItemSnsTopicArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

