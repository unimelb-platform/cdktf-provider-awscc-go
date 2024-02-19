package iotwirelesspartneraccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotwirelesspartneraccount/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_partner_account awscc_iotwireless_partner_account}.
type IotwirelessPartnerAccount interface {
	cdktf.TerraformResource
	AccountLinked() interface{}
	SetAccountLinked(val interface{})
	AccountLinkedInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Fingerprint() *string
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
	PartnerAccountId() *string
	SetPartnerAccountId(val *string)
	PartnerAccountIdInput() *string
	PartnerType() *string
	SetPartnerType(val *string)
	PartnerTypeInput() *string
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
	Sidewalk() IotwirelessPartnerAccountSidewalkOutputReference
	SidewalkInput() interface{}
	SidewalkResponse() IotwirelessPartnerAccountSidewalkResponseOutputReference
	SidewalkResponseInput() interface{}
	SidewalkUpdate() IotwirelessPartnerAccountSidewalkUpdateOutputReference
	SidewalkUpdateInput() interface{}
	Tags() IotwirelessPartnerAccountTagsList
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
	PutSidewalk(value *IotwirelessPartnerAccountSidewalk)
	PutSidewalkResponse(value *IotwirelessPartnerAccountSidewalkResponse)
	PutSidewalkUpdate(value *IotwirelessPartnerAccountSidewalkUpdate)
	PutTags(value interface{})
	ResetAccountLinked()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerAccountId()
	ResetPartnerType()
	ResetSidewalk()
	ResetSidewalkResponse()
	ResetSidewalkUpdate()
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

// The jsii proxy struct for IotwirelessPartnerAccount
type jsiiProxy_IotwirelessPartnerAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotwirelessPartnerAccount) AccountLinked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountLinked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) AccountLinkedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountLinkedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) PartnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) PartnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) PartnerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) PartnerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Sidewalk() IotwirelessPartnerAccountSidewalkOutputReference {
	var returns IotwirelessPartnerAccountSidewalkOutputReference
	_jsii_.Get(
		j,
		"sidewalk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) SidewalkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sidewalkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) SidewalkResponse() IotwirelessPartnerAccountSidewalkResponseOutputReference {
	var returns IotwirelessPartnerAccountSidewalkResponseOutputReference
	_jsii_.Get(
		j,
		"sidewalkResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) SidewalkResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sidewalkResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) SidewalkUpdate() IotwirelessPartnerAccountSidewalkUpdateOutputReference {
	var returns IotwirelessPartnerAccountSidewalkUpdateOutputReference
	_jsii_.Get(
		j,
		"sidewalkUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) SidewalkUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sidewalkUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) Tags() IotwirelessPartnerAccountTagsList {
	var returns IotwirelessPartnerAccountTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessPartnerAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_partner_account awscc_iotwireless_partner_account} Resource.
func NewIotwirelessPartnerAccount(scope constructs.Construct, id *string, config *IotwirelessPartnerAccountConfig) IotwirelessPartnerAccount {
	_init_.Initialize()

	if err := validateNewIotwirelessPartnerAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessPartnerAccount{}

	_jsii_.Create(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_partner_account awscc_iotwireless_partner_account} Resource.
func NewIotwirelessPartnerAccount_Override(i IotwirelessPartnerAccount, scope constructs.Construct, id *string, config *IotwirelessPartnerAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetAccountLinked(val interface{}) {
	if err := j.validateSetAccountLinkedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinked",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetPartnerAccountId(val *string) {
	if err := j.validateSetPartnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerAccountId",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetPartnerType(val *string) {
	if err := j.validateSetPartnerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerType",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotwirelessPartnerAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IotwirelessPartnerAccount resource upon running "cdktf plan <stack-name>".
func IotwirelessPartnerAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotwirelessPartnerAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
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
func IotwirelessPartnerAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessPartnerAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotwirelessPartnerAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessPartnerAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotwirelessPartnerAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotwirelessPartnerAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotwirelessPartnerAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotwirelessPartnerAccount.IotwirelessPartnerAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotwirelessPartnerAccount) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) PutSidewalk(value *IotwirelessPartnerAccountSidewalk) {
	if err := i.validatePutSidewalkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSidewalk",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) PutSidewalkResponse(value *IotwirelessPartnerAccountSidewalkResponse) {
	if err := i.validatePutSidewalkResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSidewalkResponse",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) PutSidewalkUpdate(value *IotwirelessPartnerAccountSidewalkUpdate) {
	if err := i.validatePutSidewalkUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSidewalkUpdate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetAccountLinked() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinked",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetPartnerAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetPartnerAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetPartnerType() {
	_jsii_.InvokeVoid(
		i,
		"resetPartnerType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetSidewalk() {
	_jsii_.InvokeVoid(
		i,
		"resetSidewalk",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetSidewalkResponse() {
	_jsii_.InvokeVoid(
		i,
		"resetSidewalkResponse",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetSidewalkUpdate() {
	_jsii_.InvokeVoid(
		i,
		"resetSidewalkUpdate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessPartnerAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessPartnerAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

