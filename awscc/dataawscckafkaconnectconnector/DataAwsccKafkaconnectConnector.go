package dataawscckafkaconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscckafkaconnectconnector/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kafkaconnect_connector awscc_kafkaconnect_connector}.
type DataAwsccKafkaconnectConnector interface {
	cdktf.TerraformDataSource
	Capacity() DataAwsccKafkaconnectConnectorCapacityOutputReference
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectorArn() *string
	ConnectorConfiguration() cdktf.StringMap
	ConnectorDescription() *string
	ConnectorName() *string
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
	SetId(val *string)
	IdInput() *string
	KafkaCluster() DataAwsccKafkaconnectConnectorKafkaClusterOutputReference
	KafkaClusterClientAuthentication() DataAwsccKafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference
	KafkaClusterEncryptionInTransit() DataAwsccKafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	KafkaConnectVersion() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogDelivery() DataAwsccKafkaconnectConnectorLogDeliveryOutputReference
	// The tree node.
	Node() constructs.Node
	Plugins() DataAwsccKafkaconnectConnectorPluginsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceExecutionRoleArn() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkerConfiguration() DataAwsccKafkaconnectConnectorWorkerConfigurationOutputReference
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsccKafkaconnectConnector
type jsiiProxy_DataAwsccKafkaconnectConnector struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Capacity() DataAwsccKafkaconnectConnectorCapacityOutputReference {
	var returns DataAwsccKafkaconnectConnectorCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ConnectorConfiguration() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"connectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ConnectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) KafkaCluster() DataAwsccKafkaconnectConnectorKafkaClusterOutputReference {
	var returns DataAwsccKafkaconnectConnectorKafkaClusterOutputReference
	_jsii_.Get(
		j,
		"kafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) KafkaClusterClientAuthentication() DataAwsccKafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference {
	var returns DataAwsccKafkaconnectConnectorKafkaClusterClientAuthenticationOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) KafkaClusterEncryptionInTransit() DataAwsccKafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference {
	var returns DataAwsccKafkaconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) KafkaConnectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaConnectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) LogDelivery() DataAwsccKafkaconnectConnectorLogDeliveryOutputReference {
	var returns DataAwsccKafkaconnectConnectorLogDeliveryOutputReference
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Plugins() DataAwsccKafkaconnectConnectorPluginsList {
	var returns DataAwsccKafkaconnectConnectorPluginsList
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector) WorkerConfiguration() DataAwsccKafkaconnectConnectorWorkerConfigurationOutputReference {
	var returns DataAwsccKafkaconnectConnectorWorkerConfigurationOutputReference
	_jsii_.Get(
		j,
		"workerConfiguration",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kafkaconnect_connector awscc_kafkaconnect_connector} Data Source.
func NewDataAwsccKafkaconnectConnector(scope constructs.Construct, id *string, config *DataAwsccKafkaconnectConnectorConfig) DataAwsccKafkaconnectConnector {
	_init_.Initialize()

	if err := validateNewDataAwsccKafkaconnectConnectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKafkaconnectConnector{}

	_jsii_.Create(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kafkaconnect_connector awscc_kafkaconnect_connector} Data Source.
func NewDataAwsccKafkaconnectConnector_Override(d DataAwsccKafkaconnectConnector, scope constructs.Construct, id *string, config *DataAwsccKafkaconnectConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKafkaconnectConnector)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccKafkaconnectConnector resource upon running "cdktf plan <stack-name>".
func DataAwsccKafkaconnectConnector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccKafkaconnectConnector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
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
func DataAwsccKafkaconnectConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKafkaconnectConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccKafkaconnectConnector_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKafkaconnectConnector_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccKafkaconnectConnector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKafkaconnectConnector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccKafkaconnectConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccKafkaconnectConnector.DataAwsccKafkaconnectConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKafkaconnectConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

