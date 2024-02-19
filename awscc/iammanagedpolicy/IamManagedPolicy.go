package iammanagedpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iammanagedpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy awscc_iam_managed_policy}.
type IamManagedPolicy interface {
	cdktf.TerraformResource
	AttachmentCount() *float64
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
	CreateDate() *string
	DefaultVersionId() *string
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
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	Id() *string
	IsAttachable() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedPolicyName() *string
	SetManagedPolicyName(val *string)
	ManagedPolicyNameInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PermissionsBoundaryUsageCount() *float64
	PolicyArn() *string
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	PolicyDocumentInput() *string
	PolicyId() *string
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
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateDate() *string
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
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
	ResetDescription()
	ResetGroups()
	ResetManagedPolicyName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetRoles()
	ResetUsers()
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

// The jsii proxy struct for IamManagedPolicy
type jsiiProxy_IamManagedPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IamManagedPolicy) AttachmentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) DefaultVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) IsAttachable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAttachable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) ManagedPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) ManagedPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PermissionsBoundaryUsageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"permissionsBoundaryUsageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) UpdateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamManagedPolicy) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy awscc_iam_managed_policy} Resource.
func NewIamManagedPolicy(scope constructs.Construct, id *string, config *IamManagedPolicyConfig) IamManagedPolicy {
	_init_.Initialize()

	if err := validateNewIamManagedPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamManagedPolicy{}

	_jsii_.Create(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy awscc_iam_managed_policy} Resource.
func NewIamManagedPolicy_Override(i IamManagedPolicy, scope constructs.Construct, id *string, config *IamManagedPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetManagedPolicyName(val *string) {
	if err := j.validateSetManagedPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPolicyName",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetRoles(val *[]*string) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_IamManagedPolicy)SetUsers(val *[]*string) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

// Generates CDKTF code for importing a IamManagedPolicy resource upon running "cdktf plan <stack-name>".
func IamManagedPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIamManagedPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iamManagedPolicy.IamManagedPolicy",
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
func IamManagedPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamManagedPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamManagedPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamManagedPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamManagedPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamManagedPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IamManagedPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iamManagedPolicy.IamManagedPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IamManagedPolicy) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IamManagedPolicy) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IamManagedPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamManagedPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamManagedPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamManagedPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamManagedPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamManagedPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamManagedPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamManagedPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamManagedPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamManagedPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IamManagedPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamManagedPolicy) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamManagedPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IamManagedPolicy) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamManagedPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetGroups() {
	_jsii_.InvokeVoid(
		i,
		"resetGroups",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetManagedPolicyName() {
	_jsii_.InvokeVoid(
		i,
		"resetManagedPolicyName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetPath() {
	_jsii_.InvokeVoid(
		i,
		"resetPath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetRoles() {
	_jsii_.InvokeVoid(
		i,
		"resetRoles",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) ResetUsers() {
	_jsii_.InvokeVoid(
		i,
		"resetUsers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamManagedPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamManagedPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

