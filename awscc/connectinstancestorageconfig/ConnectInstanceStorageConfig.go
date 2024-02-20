package connectinstancestorageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/connectinstancestorageconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config awscc_connect_instance_storage_config}.
type ConnectInstanceStorageConfig interface {
	cdktf.TerraformResource
	AssociationId() *string
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
	InstanceArn() *string
	SetInstanceArn(val *string)
	InstanceArnInput() *string
	KinesisFirehoseConfig() ConnectInstanceStorageConfigKinesisFirehoseConfigOutputReference
	KinesisFirehoseConfigInput() interface{}
	KinesisStreamConfig() ConnectInstanceStorageConfigKinesisStreamConfigOutputReference
	KinesisStreamConfigInput() interface{}
	KinesisVideoStreamConfig() ConnectInstanceStorageConfigKinesisVideoStreamConfigOutputReference
	KinesisVideoStreamConfigInput() interface{}
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
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	S3Config() ConnectInstanceStorageConfigS3ConfigOutputReference
	S3ConfigInput() interface{}
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	PutKinesisFirehoseConfig(value *ConnectInstanceStorageConfigKinesisFirehoseConfig)
	PutKinesisStreamConfig(value *ConnectInstanceStorageConfigKinesisStreamConfig)
	PutKinesisVideoStreamConfig(value *ConnectInstanceStorageConfigKinesisVideoStreamConfig)
	PutS3Config(value *ConnectInstanceStorageConfigS3Config)
	ResetKinesisFirehoseConfig()
	ResetKinesisStreamConfig()
	ResetKinesisVideoStreamConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetS3Config()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ConnectInstanceStorageConfig
type jsiiProxy_ConnectInstanceStorageConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) InstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisFirehoseConfig() ConnectInstanceStorageConfigKinesisFirehoseConfigOutputReference {
	var returns ConnectInstanceStorageConfigKinesisFirehoseConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisFirehoseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisFirehoseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisStreamConfig() ConnectInstanceStorageConfigKinesisStreamConfigOutputReference {
	var returns ConnectInstanceStorageConfigKinesisStreamConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisStreamConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisVideoStreamConfig() ConnectInstanceStorageConfigKinesisVideoStreamConfigOutputReference {
	var returns ConnectInstanceStorageConfigKinesisVideoStreamConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) KinesisVideoStreamConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) S3Config() ConnectInstanceStorageConfigS3ConfigOutputReference {
	var returns ConnectInstanceStorageConfigS3ConfigOutputReference
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) S3ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config awscc_connect_instance_storage_config} Resource.
func NewConnectInstanceStorageConfig(scope constructs.Construct, id *string, config *ConnectInstanceStorageConfigConfig) ConnectInstanceStorageConfig {
	_init_.Initialize()

	if err := validateNewConnectInstanceStorageConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectInstanceStorageConfig{}

	_jsii_.Create(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config awscc_connect_instance_storage_config} Resource.
func NewConnectInstanceStorageConfig_Override(c ConnectInstanceStorageConfig, scope constructs.Construct, id *string, config *ConnectInstanceStorageConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetInstanceArn(val *string) {
	if err := j.validateSetInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceArn",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfig)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

// Generates CDKTF code for importing a ConnectInstanceStorageConfig resource upon running "cdktf plan <stack-name>".
func ConnectInstanceStorageConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConnectInstanceStorageConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
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
func ConnectInstanceStorageConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectInstanceStorageConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConnectInstanceStorageConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectInstanceStorageConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConnectInstanceStorageConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectInstanceStorageConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConnectInstanceStorageConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.connectInstanceStorageConfig.ConnectInstanceStorageConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectInstanceStorageConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) PutKinesisFirehoseConfig(value *ConnectInstanceStorageConfigKinesisFirehoseConfig) {
	if err := c.validatePutKinesisFirehoseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisFirehoseConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) PutKinesisStreamConfig(value *ConnectInstanceStorageConfigKinesisStreamConfig) {
	if err := c.validatePutKinesisStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisStreamConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) PutKinesisVideoStreamConfig(value *ConnectInstanceStorageConfigKinesisVideoStreamConfig) {
	if err := c.validatePutKinesisVideoStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisVideoStreamConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) PutS3Config(value *ConnectInstanceStorageConfigS3Config) {
	if err := c.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3Config",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ResetKinesisFirehoseConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisFirehoseConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ResetKinesisStreamConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisStreamConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ResetKinesisVideoStreamConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisVideoStreamConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ResetS3Config() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Config",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

