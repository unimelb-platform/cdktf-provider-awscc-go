package amplifydomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/amplifydomain/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain awscc_amplify_domain}.
type AmplifyDomain interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	Arn() *string
	AutoSubDomainCreationPatterns() *[]*string
	SetAutoSubDomainCreationPatterns(val *[]*string)
	AutoSubDomainCreationPatternsInput() *[]*string
	AutoSubDomainIamRole() *string
	SetAutoSubDomainIamRole(val *string)
	AutoSubDomainIamRoleInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateRecord() *string
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
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainStatus() *string
	EnableAutoSubDomain() interface{}
	SetEnableAutoSubDomain(val interface{})
	EnableAutoSubDomainInput() interface{}
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
	StatusReason() *string
	SubDomainSettings() AmplifyDomainSubDomainSettingsList
	SubDomainSettingsInput() interface{}
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
	PutSubDomainSettings(value interface{})
	ResetAutoSubDomainCreationPatterns()
	ResetAutoSubDomainIamRole()
	ResetEnableAutoSubDomain()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for AmplifyDomain
type jsiiProxy_AmplifyDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AmplifyDomain) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) AutoSubDomainCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoSubDomainCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) AutoSubDomainCreationPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoSubDomainCreationPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) AutoSubDomainIamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSubDomainIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) AutoSubDomainIamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSubDomainIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) CertificateRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) DomainStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) EnableAutoSubDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoSubDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) EnableAutoSubDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoSubDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) SubDomainSettings() AmplifyDomainSubDomainSettingsList {
	var returns AmplifyDomainSubDomainSettingsList
	_jsii_.Get(
		j,
		"subDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) SubDomainSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subDomainSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmplifyDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain awscc_amplify_domain} Resource.
func NewAmplifyDomain(scope constructs.Construct, id *string, config *AmplifyDomainConfig) AmplifyDomain {
	_init_.Initialize()

	if err := validateNewAmplifyDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmplifyDomain{}

	_jsii_.Create(
		"awscc.amplifyDomain.AmplifyDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_domain awscc_amplify_domain} Resource.
func NewAmplifyDomain_Override(a AmplifyDomain, scope constructs.Construct, id *string, config *AmplifyDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.amplifyDomain.AmplifyDomain",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetAutoSubDomainCreationPatterns(val *[]*string) {
	if err := j.validateSetAutoSubDomainCreationPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubDomainCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetAutoSubDomainIamRole(val *string) {
	if err := j.validateSetAutoSubDomainIamRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubDomainIamRole",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetEnableAutoSubDomain(val interface{}) {
	if err := j.validateSetEnableAutoSubDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoSubDomain",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AmplifyDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a AmplifyDomain resource upon running "cdktf plan <stack-name>".
func AmplifyDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAmplifyDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.amplifyDomain.AmplifyDomain",
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
func AmplifyDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyDomain.AmplifyDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyDomain.AmplifyDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AmplifyDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmplifyDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.amplifyDomain.AmplifyDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AmplifyDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.amplifyDomain.AmplifyDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmplifyDomain) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AmplifyDomain) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AmplifyDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AmplifyDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AmplifyDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AmplifyDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AmplifyDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AmplifyDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AmplifyDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AmplifyDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AmplifyDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AmplifyDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AmplifyDomain) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AmplifyDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AmplifyDomain) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AmplifyDomain) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AmplifyDomain) PutSubDomainSettings(value interface{}) {
	if err := a.validatePutSubDomainSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubDomainSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AmplifyDomain) ResetAutoSubDomainCreationPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubDomainCreationPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomain) ResetAutoSubDomainIamRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubDomainIamRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomain) ResetEnableAutoSubDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoSubDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AmplifyDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmplifyDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

