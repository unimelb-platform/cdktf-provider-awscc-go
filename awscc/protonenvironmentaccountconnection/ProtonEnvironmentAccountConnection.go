package protonenvironmentaccountconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/protonenvironmentaccountconnection/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/proton_environment_account_connection awscc_proton_environment_account_connection}.
type ProtonEnvironmentAccountConnection interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CodebuildRoleArn() *string
	SetCodebuildRoleArn(val *string)
	CodebuildRoleArnInput() *string
	ComponentRoleArn() *string
	SetComponentRoleArn(val *string)
	ComponentRoleArnInput() *string
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
	EnvironmentAccountConnectionId() *string
	EnvironmentAccountId() *string
	SetEnvironmentAccountId(val *string)
	EnvironmentAccountIdInput() *string
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	EnvironmentNameInput() *string
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
	ManagementAccountId() *string
	SetManagementAccountId(val *string)
	ManagementAccountIdInput() *string
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
	Status() *string
	Tags() ProtonEnvironmentAccountConnectionTagsList
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
	PutTags(value interface{})
	ResetCodebuildRoleArn()
	ResetComponentRoleArn()
	ResetEnvironmentAccountId()
	ResetEnvironmentName()
	ResetManagementAccountId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleArn()
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

// The jsii proxy struct for ProtonEnvironmentAccountConnection
type jsiiProxy_ProtonEnvironmentAccountConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) CodebuildRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codebuildRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) CodebuildRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codebuildRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ComponentRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ComponentRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) EnvironmentAccountConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) EnvironmentAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) EnvironmentAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) EnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ManagementAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) ManagementAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) Tags() ProtonEnvironmentAccountConnectionTagsList {
	var returns ProtonEnvironmentAccountConnectionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/proton_environment_account_connection awscc_proton_environment_account_connection} Resource.
func NewProtonEnvironmentAccountConnection(scope constructs.Construct, id *string, config *ProtonEnvironmentAccountConnectionConfig) ProtonEnvironmentAccountConnection {
	_init_.Initialize()

	if err := validateNewProtonEnvironmentAccountConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProtonEnvironmentAccountConnection{}

	_jsii_.Create(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/proton_environment_account_connection awscc_proton_environment_account_connection} Resource.
func NewProtonEnvironmentAccountConnection_Override(p ProtonEnvironmentAccountConnection, scope constructs.Construct, id *string, config *ProtonEnvironmentAccountConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetCodebuildRoleArn(val *string) {
	if err := j.validateSetCodebuildRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codebuildRoleArn",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetComponentRoleArn(val *string) {
	if err := j.validateSetComponentRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentRoleArn",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetEnvironmentAccountId(val *string) {
	if err := j.validateSetEnvironmentAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentAccountId",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetEnvironmentName(val *string) {
	if err := j.validateSetEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetManagementAccountId(val *string) {
	if err := j.validateSetManagementAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementAccountId",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProtonEnvironmentAccountConnection)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Generates CDKTF code for importing a ProtonEnvironmentAccountConnection resource upon running "cdktf plan <stack-name>".
func ProtonEnvironmentAccountConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProtonEnvironmentAccountConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
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
func ProtonEnvironmentAccountConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProtonEnvironmentAccountConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProtonEnvironmentAccountConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProtonEnvironmentAccountConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProtonEnvironmentAccountConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProtonEnvironmentAccountConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProtonEnvironmentAccountConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.protonEnvironmentAccountConnection.ProtonEnvironmentAccountConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) PutTags(value interface{}) {
	if err := p.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetCodebuildRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetCodebuildRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetComponentRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetComponentRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetEnvironmentAccountId() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironmentAccountId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetEnvironmentName() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironmentName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetManagementAccountId() {
	_jsii_.InvokeVoid(
		p,
		"resetManagementAccountId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProtonEnvironmentAccountConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

