package kafkaconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kafkaconnectconnector/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector awscc_kafkaconnect_connector}.
type KafkaconnectConnector interface {
	cdktf.TerraformResource
	Capacity() KafkaconnectConnectorCapacityOutputReference
	CapacityInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorArn() *string
	ConnectorConfiguration() *map[string]*string
	SetConnectorConfiguration(val *map[string]*string)
	ConnectorConfigurationInput() *map[string]*string
	ConnectorDescription() *string
	SetConnectorDescription(val *string)
	ConnectorDescriptionInput() *string
	ConnectorName() *string
	SetConnectorName(val *string)
	ConnectorNameInput() *string
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
	KafkaCluster() KafkaconnectConnectorKafkaClusterOutputReference
	KafkaClusterClientAuthentication() KafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference
	KafkaClusterClientAuthenticationInput() interface{}
	KafkaClusterEncryptionInTransit() KafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	KafkaClusterEncryptionInTransitInput() interface{}
	KafkaClusterInput() interface{}
	KafkaConnectVersion() *string
	SetKafkaConnectVersion(val *string)
	KafkaConnectVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogDelivery() KafkaconnectConnectorLogDeliveryOutputReference
	LogDeliveryInput() interface{}
	// The tree node.
	Node() constructs.Node
	Plugins() KafkaconnectConnectorPluginsList
	PluginsInput() interface{}
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
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	ServiceExecutionRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkerConfiguration() KafkaconnectConnectorWorkerConfigurationOutputReference
	WorkerConfigurationInput() interface{}
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
	PutCapacity(value *KafkaconnectConnectorCapacity)
	PutKafkaCluster(value *KafkaconnectConnectorKafkaCluster)
	PutKafkaClusterClientAuthentication(value *KafkaconnectConnectorKafkaClusterClientAuthentication)
	PutKafkaClusterEncryptionInTransit(value *KafkaconnectConnectorKafkaClusterEncryptionInTransit)
	PutLogDelivery(value *KafkaconnectConnectorLogDelivery)
	PutPlugins(value interface{})
	PutWorkerConfiguration(value *KafkaconnectConnectorWorkerConfiguration)
	ResetConnectorDescription()
	ResetLogDelivery()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetWorkerConfiguration()
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

// The jsii proxy struct for KafkaconnectConnector
type jsiiProxy_KafkaconnectConnector struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KafkaconnectConnector) Capacity() KafkaconnectConnectorCapacityOutputReference {
	var returns KafkaconnectConnectorCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) CapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorConfiguration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConnectorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaCluster() KafkaconnectConnectorKafkaClusterOutputReference {
	var returns KafkaconnectConnectorKafkaClusterOutputReference
	_jsii_.Get(
		j,
		"kafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaClusterClientAuthentication() KafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference {
	var returns KafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaClusterClientAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaClusterEncryptionInTransit() KafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference {
	var returns KafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaClusterEncryptionInTransitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaConnectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaConnectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) KafkaConnectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaConnectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) LogDelivery() KafkaconnectConnectorLogDeliveryOutputReference {
	var returns KafkaconnectConnectorLogDeliveryOutputReference
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) LogDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Plugins() KafkaconnectConnectorPluginsList {
	var returns KafkaconnectConnectorPluginsList
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) PluginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) ServiceExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) WorkerConfiguration() KafkaconnectConnectorWorkerConfigurationOutputReference {
	var returns KafkaconnectConnectorWorkerConfigurationOutputReference
	_jsii_.Get(
		j,
		"workerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnector) WorkerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector awscc_kafkaconnect_connector} Resource.
func NewKafkaconnectConnector(scope constructs.Construct, id *string, config *KafkaconnectConnectorConfig) KafkaconnectConnector {
	_init_.Initialize()

	if err := validateNewKafkaconnectConnectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KafkaconnectConnector{}

	_jsii_.Create(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector awscc_kafkaconnect_connector} Resource.
func NewKafkaconnectConnector_Override(k KafkaconnectConnector, scope constructs.Construct, id *string, config *KafkaconnectConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetConnectorConfiguration(val *map[string]*string) {
	if err := j.validateSetConnectorConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorConfiguration",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetConnectorDescription(val *string) {
	if err := j.validateSetConnectorDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorDescription",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetConnectorName(val *string) {
	if err := j.validateSetConnectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorName",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetKafkaConnectVersion(val *string) {
	if err := j.validateSetKafkaConnectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaConnectVersion",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnector)SetServiceExecutionRoleArn(val *string) {
	if err := j.validateSetServiceExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

// Generates CDKTF code for importing a KafkaconnectConnector resource upon running "cdktf plan <stack-name>".
func KafkaconnectConnector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKafkaconnectConnector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
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
func KafkaconnectConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKafkaconnectConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KafkaconnectConnector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKafkaconnectConnector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KafkaconnectConnector_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKafkaconnectConnector_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KafkaconnectConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.kafkaconnectConnector.KafkaconnectConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutCapacity(value *KafkaconnectConnectorCapacity) {
	if err := k.validatePutCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCapacity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutKafkaCluster(value *KafkaconnectConnectorKafkaCluster) {
	if err := k.validatePutKafkaClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKafkaCluster",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutKafkaClusterClientAuthentication(value *KafkaconnectConnectorKafkaClusterClientAuthentication) {
	if err := k.validatePutKafkaClusterClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKafkaClusterClientAuthentication",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutKafkaClusterEncryptionInTransit(value *KafkaconnectConnectorKafkaClusterEncryptionInTransit) {
	if err := k.validatePutKafkaClusterEncryptionInTransitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKafkaClusterEncryptionInTransit",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutLogDelivery(value *KafkaconnectConnectorLogDelivery) {
	if err := k.validatePutLogDeliveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLogDelivery",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutPlugins(value interface{}) {
	if err := k.validatePutPluginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPlugins",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) PutWorkerConfiguration(value *KafkaconnectConnectorWorkerConfiguration) {
	if err := k.validatePutWorkerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWorkerConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnector) ResetConnectorDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetConnectorDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnector) ResetLogDelivery() {
	_jsii_.InvokeVoid(
		k,
		"resetLogDelivery",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnector) ResetWorkerConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetWorkerConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

