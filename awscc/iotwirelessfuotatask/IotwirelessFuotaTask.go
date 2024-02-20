package iotwirelessfuotatask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotwirelessfuotatask/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_fuota_task awscc_iotwireless_fuota_task}.
type IotwirelessFuotaTask interface {
	cdktf.TerraformResource
	Arn() *string
	AssociateMulticastGroup() *string
	SetAssociateMulticastGroup(val *string)
	AssociateMulticastGroupInput() *string
	AssociateWirelessDevice() *string
	SetAssociateWirelessDevice(val *string)
	AssociateWirelessDeviceInput() *string
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
	DisassociateMulticastGroup() *string
	SetDisassociateMulticastGroup(val *string)
	DisassociateMulticastGroupInput() *string
	DisassociateWirelessDevice() *string
	SetDisassociateWirelessDevice(val *string)
	DisassociateWirelessDeviceInput() *string
	FirmwareUpdateImage() *string
	SetFirmwareUpdateImage(val *string)
	FirmwareUpdateImageInput() *string
	FirmwareUpdateRole() *string
	SetFirmwareUpdateRole(val *string)
	FirmwareUpdateRoleInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FuotaTaskStatus() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoRaWan() IotwirelessFuotaTaskLoRaWanOutputReference
	LoRaWanInput() interface{}
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
	Tags() IotwirelessFuotaTaskTagsList
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
	PutLoRaWan(value *IotwirelessFuotaTaskLoRaWan)
	PutTags(value interface{})
	ResetAssociateMulticastGroup()
	ResetAssociateWirelessDevice()
	ResetDescription()
	ResetDisassociateMulticastGroup()
	ResetDisassociateWirelessDevice()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for IotwirelessFuotaTask
type jsiiProxy_IotwirelessFuotaTask struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotwirelessFuotaTask) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) AssociateMulticastGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateMulticastGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) AssociateMulticastGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateMulticastGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) AssociateWirelessDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateWirelessDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) AssociateWirelessDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateWirelessDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DisassociateMulticastGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disassociateMulticastGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DisassociateMulticastGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disassociateMulticastGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DisassociateWirelessDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disassociateWirelessDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) DisassociateWirelessDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disassociateWirelessDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FirmwareUpdateImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareUpdateImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FirmwareUpdateImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareUpdateImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FirmwareUpdateRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareUpdateRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FirmwareUpdateRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareUpdateRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) FuotaTaskStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fuotaTaskStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) LoRaWan() IotwirelessFuotaTaskLoRaWanOutputReference {
	var returns IotwirelessFuotaTaskLoRaWanOutputReference
	_jsii_.Get(
		j,
		"loRaWan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) LoRaWanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loRaWanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) Tags() IotwirelessFuotaTaskTagsList {
	var returns IotwirelessFuotaTaskTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessFuotaTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_fuota_task awscc_iotwireless_fuota_task} Resource.
func NewIotwirelessFuotaTask(scope constructs.Construct, id *string, config *IotwirelessFuotaTaskConfig) IotwirelessFuotaTask {
	_init_.Initialize()

	if err := validateNewIotwirelessFuotaTaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessFuotaTask{}

	_jsii_.Create(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_fuota_task awscc_iotwireless_fuota_task} Resource.
func NewIotwirelessFuotaTask_Override(i IotwirelessFuotaTask, scope constructs.Construct, id *string, config *IotwirelessFuotaTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetAssociateMulticastGroup(val *string) {
	if err := j.validateSetAssociateMulticastGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateMulticastGroup",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetAssociateWirelessDevice(val *string) {
	if err := j.validateSetAssociateWirelessDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateWirelessDevice",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetDisassociateMulticastGroup(val *string) {
	if err := j.validateSetDisassociateMulticastGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disassociateMulticastGroup",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetDisassociateWirelessDevice(val *string) {
	if err := j.validateSetDisassociateWirelessDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disassociateWirelessDevice",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetFirmwareUpdateImage(val *string) {
	if err := j.validateSetFirmwareUpdateImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firmwareUpdateImage",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetFirmwareUpdateRole(val *string) {
	if err := j.validateSetFirmwareUpdateRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firmwareUpdateRole",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotwirelessFuotaTask)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IotwirelessFuotaTask resource upon running "cdktf plan <stack-name>".
func IotwirelessFuotaTask_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotwirelessFuotaTask_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
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
func IotwirelessFuotaTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessFuotaTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotwirelessFuotaTask_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessFuotaTask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotwirelessFuotaTask_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessFuotaTask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotwirelessFuotaTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotwirelessFuotaTask.IotwirelessFuotaTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotwirelessFuotaTask) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessFuotaTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessFuotaTask) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessFuotaTask) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) PutLoRaWan(value *IotwirelessFuotaTaskLoRaWan) {
	if err := i.validatePutLoRaWanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLoRaWan",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetAssociateMulticastGroup() {
	_jsii_.InvokeVoid(
		i,
		"resetAssociateMulticastGroup",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetAssociateWirelessDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetAssociateWirelessDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetDisassociateMulticastGroup() {
	_jsii_.InvokeVoid(
		i,
		"resetDisassociateMulticastGroup",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetDisassociateWirelessDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetDisassociateWirelessDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessFuotaTask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessFuotaTask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessFuotaTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessFuotaTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

