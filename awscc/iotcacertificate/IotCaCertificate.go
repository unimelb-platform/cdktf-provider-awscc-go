package iotcacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotcacertificate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate awscc_iot_ca_certificate}.
type IotCaCertificate interface {
	cdktf.TerraformResource
	Arn() *string
	AutoRegistrationStatus() *string
	SetAutoRegistrationStatus(val *string)
	AutoRegistrationStatusInput() *string
	CaCertificatePem() *string
	SetCaCertificatePem(val *string)
	CaCertificatePemInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateMode() *string
	SetCertificateMode(val *string)
	CertificateModeInput() *string
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
	RegistrationConfig() IotCaCertificateRegistrationConfigOutputReference
	RegistrationConfigInput() interface{}
	RemoveAutoRegistration() interface{}
	SetRemoveAutoRegistration(val interface{})
	RemoveAutoRegistrationInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() IotCaCertificateTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VerificationCertificatePem() *string
	SetVerificationCertificatePem(val *string)
	VerificationCertificatePemInput() *string
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
	PutRegistrationConfig(value *IotCaCertificateRegistrationConfig)
	PutTags(value interface{})
	ResetAutoRegistrationStatus()
	ResetCertificateMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistrationConfig()
	ResetRemoveAutoRegistration()
	ResetTags()
	ResetVerificationCertificatePem()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IotCaCertificate
type jsiiProxy_IotCaCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotCaCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) AutoRegistrationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoRegistrationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) AutoRegistrationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoRegistrationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) CaCertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) CaCertificatePemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatePemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) CertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) CertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) RegistrationConfig() IotCaCertificateRegistrationConfigOutputReference {
	var returns IotCaCertificateRegistrationConfigOutputReference
	_jsii_.Get(
		j,
		"registrationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) RegistrationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) RemoveAutoRegistration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeAutoRegistration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) RemoveAutoRegistrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeAutoRegistrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) Tags() IotCaCertificateTagsList {
	var returns IotCaCertificateTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) VerificationCertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationCertificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCaCertificate) VerificationCertificatePemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationCertificatePemInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate awscc_iot_ca_certificate} Resource.
func NewIotCaCertificate(scope constructs.Construct, id *string, config *IotCaCertificateConfig) IotCaCertificate {
	_init_.Initialize()

	if err := validateNewIotCaCertificateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotCaCertificate{}

	_jsii_.Create(
		"awscc.iotCaCertificate.IotCaCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_ca_certificate awscc_iot_ca_certificate} Resource.
func NewIotCaCertificate_Override(i IotCaCertificate, scope constructs.Construct, id *string, config *IotCaCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotCaCertificate.IotCaCertificate",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetAutoRegistrationStatus(val *string) {
	if err := j.validateSetAutoRegistrationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRegistrationStatus",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetCaCertificatePem(val *string) {
	if err := j.validateSetCaCertificatePemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatePem",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetCertificateMode(val *string) {
	if err := j.validateSetCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateMode",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetRemoveAutoRegistration(val interface{}) {
	if err := j.validateSetRemoveAutoRegistrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeAutoRegistration",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_IotCaCertificate)SetVerificationCertificatePem(val *string) {
	if err := j.validateSetVerificationCertificatePemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationCertificatePem",
		val,
	)
}

// Generates CDKTF code for importing a IotCaCertificate resource upon running "cdktf plan <stack-name>".
func IotCaCertificate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotCaCertificate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotCaCertificate.IotCaCertificate",
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
func IotCaCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotCaCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotCaCertificate.IotCaCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotCaCertificate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotCaCertificate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotCaCertificate.IotCaCertificate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotCaCertificate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotCaCertificate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotCaCertificate.IotCaCertificate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotCaCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotCaCertificate.IotCaCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotCaCertificate) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotCaCertificate) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotCaCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotCaCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotCaCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotCaCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotCaCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotCaCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotCaCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotCaCertificate) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotCaCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotCaCertificate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotCaCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotCaCertificate) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotCaCertificate) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotCaCertificate) PutRegistrationConfig(value *IotCaCertificateRegistrationConfig) {
	if err := i.validatePutRegistrationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRegistrationConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotCaCertificate) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetAutoRegistrationStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetAutoRegistrationStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetCertificateMode() {
	_jsii_.InvokeVoid(
		i,
		"resetCertificateMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetRegistrationConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetRegistrationConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetRemoveAutoRegistration() {
	_jsii_.InvokeVoid(
		i,
		"resetRemoveAutoRegistration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) ResetVerificationCertificatePem() {
	_jsii_.InvokeVoid(
		i,
		"resetVerificationCertificatePem",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCaCertificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCaCertificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCaCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCaCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

