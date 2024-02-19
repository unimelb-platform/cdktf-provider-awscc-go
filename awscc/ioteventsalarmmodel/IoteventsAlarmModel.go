package ioteventsalarmmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsalarmmodel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model awscc_iotevents_alarm_model}.
type IoteventsAlarmModel interface {
	cdktf.TerraformResource
	AlarmCapabilities() IoteventsAlarmModelAlarmCapabilitiesOutputReference
	AlarmCapabilitiesInput() interface{}
	AlarmEventActions() IoteventsAlarmModelAlarmEventActionsOutputReference
	AlarmEventActionsInput() interface{}
	AlarmModelDescription() *string
	SetAlarmModelDescription(val *string)
	AlarmModelDescriptionInput() *string
	AlarmModelName() *string
	SetAlarmModelName(val *string)
	AlarmModelNameInput() *string
	AlarmRule() IoteventsAlarmModelAlarmRuleOutputReference
	AlarmRuleInput() interface{}
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
	Key() *string
	SetKey(val *string)
	KeyInput() *string
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Severity() *float64
	SetSeverity(val *float64)
	SeverityInput() *float64
	Tags() IoteventsAlarmModelTagsList
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
	PutAlarmCapabilities(value *IoteventsAlarmModelAlarmCapabilities)
	PutAlarmEventActions(value *IoteventsAlarmModelAlarmEventActions)
	PutAlarmRule(value *IoteventsAlarmModelAlarmRule)
	PutTags(value interface{})
	ResetAlarmCapabilities()
	ResetAlarmEventActions()
	ResetAlarmModelDescription()
	ResetAlarmModelName()
	ResetKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSeverity()
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

// The jsii proxy struct for IoteventsAlarmModel
type jsiiProxy_IoteventsAlarmModel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmCapabilities() IoteventsAlarmModelAlarmCapabilitiesOutputReference {
	var returns IoteventsAlarmModelAlarmCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"alarmCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmEventActions() IoteventsAlarmModelAlarmEventActionsOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsOutputReference
	_jsii_.Get(
		j,
		"alarmEventActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmEventActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmEventActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmRule() IoteventsAlarmModelAlarmRuleOutputReference {
	var returns IoteventsAlarmModelAlarmRuleOutputReference
	_jsii_.Get(
		j,
		"alarmRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) AlarmRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) SeverityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) Tags() IoteventsAlarmModelTagsList {
	var returns IoteventsAlarmModelTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model awscc_iotevents_alarm_model} Resource.
func NewIoteventsAlarmModel(scope constructs.Construct, id *string, config *IoteventsAlarmModelConfig) IoteventsAlarmModel {
	_init_.Initialize()

	if err := validateNewIoteventsAlarmModelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsAlarmModel{}

	_jsii_.Create(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model awscc_iotevents_alarm_model} Resource.
func NewIoteventsAlarmModel_Override(i IoteventsAlarmModel, scope constructs.Construct, id *string, config *IoteventsAlarmModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetAlarmModelDescription(val *string) {
	if err := j.validateSetAlarmModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmModelDescription",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetAlarmModelName(val *string) {
	if err := j.validateSetAlarmModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmModelName",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModel)SetSeverity(val *float64) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

// Generates CDKTF code for importing a IoteventsAlarmModel resource upon running "cdktf plan <stack-name>".
func IoteventsAlarmModel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIoteventsAlarmModel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
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
func IoteventsAlarmModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIoteventsAlarmModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IoteventsAlarmModel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIoteventsAlarmModel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IoteventsAlarmModel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIoteventsAlarmModel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IoteventsAlarmModel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) PutAlarmCapabilities(value *IoteventsAlarmModelAlarmCapabilities) {
	if err := i.validatePutAlarmCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAlarmCapabilities",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) PutAlarmEventActions(value *IoteventsAlarmModelAlarmEventActions) {
	if err := i.validatePutAlarmEventActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAlarmEventActions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) PutAlarmRule(value *IoteventsAlarmModelAlarmRule) {
	if err := i.validatePutAlarmRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAlarmRule",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetAlarmCapabilities() {
	_jsii_.InvokeVoid(
		i,
		"resetAlarmCapabilities",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetAlarmEventActions() {
	_jsii_.InvokeVoid(
		i,
		"resetAlarmEventActions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetAlarmModelDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetAlarmModelDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetAlarmModelName() {
	_jsii_.InvokeVoid(
		i,
		"resetAlarmModelName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetKey() {
	_jsii_.InvokeVoid(
		i,
		"resetKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetSeverity() {
	_jsii_.InvokeVoid(
		i,
		"resetSeverity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

