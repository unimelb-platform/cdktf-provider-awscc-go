package connectuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectuser/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user awscc_connect_user}.
type ConnectUser interface {
	cdktf.TerraformResource
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
	DirectoryUserId() *string
	SetDirectoryUserId(val *string)
	DirectoryUserIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HierarchyGroupArn() *string
	SetHierarchyGroupArn(val *string)
	HierarchyGroupArnInput() *string
	Id() *string
	IdentityInfo() ConnectUserIdentityInfoOutputReference
	IdentityInfoInput() interface{}
	InstanceArn() *string
	SetInstanceArn(val *string)
	InstanceArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PhoneConfig() ConnectUserPhoneConfigOutputReference
	PhoneConfigInput() interface{}
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
	RoutingProfileArn() *string
	SetRoutingProfileArn(val *string)
	RoutingProfileArnInput() *string
	SecurityProfileArns() *[]*string
	SetSecurityProfileArns(val *[]*string)
	SecurityProfileArnsInput() *[]*string
	Tags() ConnectUserTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserArn() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UserProficiencies() ConnectUserUserProficienciesList
	UserProficienciesInput() interface{}
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
	PutIdentityInfo(value *ConnectUserIdentityInfo)
	PutPhoneConfig(value *ConnectUserPhoneConfig)
	PutTags(value interface{})
	PutUserProficiencies(value interface{})
	ResetDirectoryUserId()
	ResetHierarchyGroupArn()
	ResetIdentityInfo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetTags()
	ResetUserProficiencies()
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

// The jsii proxy struct for ConnectUser
type jsiiProxy_ConnectUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConnectUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) DirectoryUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) DirectoryUserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryUserIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) HierarchyGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) HierarchyGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) IdentityInfo() ConnectUserIdentityInfoOutputReference {
	var returns ConnectUserIdentityInfoOutputReference
	_jsii_.Get(
		j,
		"identityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) IdentityInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) InstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) PhoneConfig() ConnectUserPhoneConfigOutputReference {
	var returns ConnectUserPhoneConfigOutputReference
	_jsii_.Get(
		j,
		"phoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) PhoneConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) RoutingProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) RoutingProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) SecurityProfileArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityProfileArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) SecurityProfileArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityProfileArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Tags() ConnectUserTagsList {
	var returns ConnectUserTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) UserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) UserProficiencies() ConnectUserUserProficienciesList {
	var returns ConnectUserUserProficienciesList
	_jsii_.Get(
		j,
		"userProficiencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUser) UserProficienciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProficienciesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user awscc_connect_user} Resource.
func NewConnectUser(scope constructs.Construct, id *string, config *ConnectUserConfig) ConnectUser {
	_init_.Initialize()

	if err := validateNewConnectUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUser{}

	_jsii_.Create(
		"awscc.connectUser.ConnectUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user awscc_connect_user} Resource.
func NewConnectUser_Override(c ConnectUser, scope constructs.Construct, id *string, config *ConnectUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectUser.ConnectUser",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConnectUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetDirectoryUserId(val *string) {
	if err := j.validateSetDirectoryUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryUserId",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetHierarchyGroupArn(val *string) {
	if err := j.validateSetHierarchyGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyGroupArn",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetInstanceArn(val *string) {
	if err := j.validateSetInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceArn",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetRoutingProfileArn(val *string) {
	if err := j.validateSetRoutingProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingProfileArn",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetSecurityProfileArns(val *[]*string) {
	if err := j.validateSetSecurityProfileArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileArns",
		val,
	)
}

func (j *jsiiProxy_ConnectUser)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a ConnectUser resource upon running "cdktf plan <stack-name>".
func ConnectUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConnectUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.connectUser.ConnectUser",
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
func ConnectUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectUser.ConnectUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConnectUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectUser.ConnectUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConnectUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectUser.ConnectUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConnectUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.connectUser.ConnectUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectUser) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConnectUser) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConnectUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConnectUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConnectUser) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConnectUser) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConnectUser) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConnectUser) PutIdentityInfo(value *ConnectUserIdentityInfo) {
	if err := c.validatePutIdentityInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentityInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUser) PutPhoneConfig(value *ConnectUserPhoneConfig) {
	if err := c.validatePutPhoneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhoneConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUser) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUser) PutUserProficiencies(value interface{}) {
	if err := c.validatePutUserProficienciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUserProficiencies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUser) ResetDirectoryUserId() {
	_jsii_.InvokeVoid(
		c,
		"resetDirectoryUserId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetHierarchyGroupArn() {
	_jsii_.InvokeVoid(
		c,
		"resetHierarchyGroupArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetIdentityInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) ResetUserProficiencies() {
	_jsii_.InvokeVoid(
		c,
		"resetUserProficiencies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

