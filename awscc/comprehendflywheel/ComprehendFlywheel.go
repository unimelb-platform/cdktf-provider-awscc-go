package comprehendflywheel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/comprehendflywheel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel awscc_comprehend_flywheel}.
type ComprehendFlywheel interface {
	cdktf.TerraformResource
	ActiveModelArn() *string
	SetActiveModelArn(val *string)
	ActiveModelArnInput() *string
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
	DataAccessRoleArn() *string
	SetDataAccessRoleArn(val *string)
	DataAccessRoleArnInput() *string
	DataLakeS3Uri() *string
	SetDataLakeS3Uri(val *string)
	DataLakeS3UriInput() *string
	DataSecurityConfig() ComprehendFlywheelDataSecurityConfigOutputReference
	DataSecurityConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FlywheelName() *string
	SetFlywheelName(val *string)
	FlywheelNameInput() *string
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
	ModelType() *string
	SetModelType(val *string)
	ModelTypeInput() *string
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
	Tags() ComprehendFlywheelTagsList
	TagsInput() interface{}
	TaskConfig() ComprehendFlywheelTaskConfigOutputReference
	TaskConfigInput() interface{}
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
	PutDataSecurityConfig(value *ComprehendFlywheelDataSecurityConfig)
	PutTags(value interface{})
	PutTaskConfig(value *ComprehendFlywheelTaskConfig)
	ResetActiveModelArn()
	ResetDataSecurityConfig()
	ResetModelType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTaskConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComprehendFlywheel
type jsiiProxy_ComprehendFlywheel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComprehendFlywheel) ActiveModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) ActiveModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataLakeS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataLakeS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataSecurityConfig() ComprehendFlywheelDataSecurityConfigOutputReference {
	var returns ComprehendFlywheelDataSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"dataSecurityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DataSecurityConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSecurityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) FlywheelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flywheelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) FlywheelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flywheelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) ModelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) ModelTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) Tags() ComprehendFlywheelTagsList {
	var returns ComprehendFlywheelTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TaskConfig() ComprehendFlywheelTaskConfigOutputReference {
	var returns ComprehendFlywheelTaskConfigOutputReference
	_jsii_.Get(
		j,
		"taskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TaskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendFlywheel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel awscc_comprehend_flywheel} Resource.
func NewComprehendFlywheel(scope constructs.Construct, id *string, config *ComprehendFlywheelConfig) ComprehendFlywheel {
	_init_.Initialize()

	if err := validateNewComprehendFlywheelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendFlywheel{}

	_jsii_.Create(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel awscc_comprehend_flywheel} Resource.
func NewComprehendFlywheel_Override(c ComprehendFlywheel, scope constructs.Construct, id *string, config *ComprehendFlywheelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetActiveModelArn(val *string) {
	if err := j.validateSetActiveModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeModelArn",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetDataAccessRoleArn(val *string) {
	if err := j.validateSetDataAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetDataLakeS3Uri(val *string) {
	if err := j.validateSetDataLakeS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLakeS3Uri",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetFlywheelName(val *string) {
	if err := j.validateSetFlywheelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flywheelName",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetModelType(val *string) {
	if err := j.validateSetModelTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelType",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComprehendFlywheel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ComprehendFlywheel resource upon running "cdktf plan <stack-name>".
func ComprehendFlywheel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComprehendFlywheel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
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
func ComprehendFlywheel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendFlywheel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComprehendFlywheel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendFlywheel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComprehendFlywheel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendFlywheel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComprehendFlywheel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.comprehendFlywheel.ComprehendFlywheel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComprehendFlywheel) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendFlywheel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendFlywheel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendFlywheel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendFlywheel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendFlywheel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendFlywheel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendFlywheel) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendFlywheel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendFlywheel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendFlywheel) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) PutDataSecurityConfig(value *ComprehendFlywheelDataSecurityConfig) {
	if err := c.validatePutDataSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataSecurityConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) PutTaskConfig(value *ComprehendFlywheelTaskConfig) {
	if err := c.validatePutTaskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaskConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetActiveModelArn() {
	_jsii_.InvokeVoid(
		c,
		"resetActiveModelArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetDataSecurityConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDataSecurityConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetModelType() {
	_jsii_.InvokeVoid(
		c,
		"resetModelType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) ResetTaskConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendFlywheel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendFlywheel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendFlywheel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendFlywheel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

