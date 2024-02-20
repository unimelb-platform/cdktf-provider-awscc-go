package mskcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mskcluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster awscc_msk_cluster}.
type MskCluster interface {
	cdktf.TerraformResource
	Arn() *string
	BrokerNodeGroupInfo() MskClusterBrokerNodeGroupInfoOutputReference
	BrokerNodeGroupInfoInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAuthentication() MskClusterClientAuthenticationOutputReference
	ClientAuthenticationInput() interface{}
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	ConfigurationInfo() MskClusterConfigurationInfoOutputReference
	ConfigurationInfoInput() interface{}
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
	CurrentVersion() *string
	SetCurrentVersion(val *string)
	CurrentVersionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionInfo() MskClusterEncryptionInfoOutputReference
	EncryptionInfoInput() interface{}
	EnhancedMonitoring() *string
	SetEnhancedMonitoring(val *string)
	EnhancedMonitoringInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KafkaVersion() *string
	SetKafkaVersion(val *string)
	KafkaVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingInfo() MskClusterLoggingInfoOutputReference
	LoggingInfoInput() interface{}
	// The tree node.
	Node() constructs.Node
	NumberOfBrokerNodes() *float64
	SetNumberOfBrokerNodes(val *float64)
	NumberOfBrokerNodesInput() *float64
	OpenMonitoring() MskClusterOpenMonitoringOutputReference
	OpenMonitoringInput() interface{}
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
	StorageMode() *string
	SetStorageMode(val *string)
	StorageModeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	PutBrokerNodeGroupInfo(value *MskClusterBrokerNodeGroupInfo)
	PutClientAuthentication(value *MskClusterClientAuthentication)
	PutConfigurationInfo(value *MskClusterConfigurationInfo)
	PutEncryptionInfo(value *MskClusterEncryptionInfo)
	PutLoggingInfo(value *MskClusterLoggingInfo)
	PutOpenMonitoring(value *MskClusterOpenMonitoring)
	ResetClientAuthentication()
	ResetConfigurationInfo()
	ResetCurrentVersion()
	ResetEncryptionInfo()
	ResetEnhancedMonitoring()
	ResetLoggingInfo()
	ResetOpenMonitoring()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageMode()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MskCluster
type jsiiProxy_MskCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MskCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) BrokerNodeGroupInfo() MskClusterBrokerNodeGroupInfoOutputReference {
	var returns MskClusterBrokerNodeGroupInfoOutputReference
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) BrokerNodeGroupInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"brokerNodeGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ClientAuthentication() MskClusterClientAuthenticationOutputReference {
	var returns MskClusterClientAuthenticationOutputReference
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ClientAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ConfigurationInfo() MskClusterConfigurationInfoOutputReference {
	var returns MskClusterConfigurationInfoOutputReference
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ConfigurationInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) CurrentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) EncryptionInfo() MskClusterEncryptionInfoOutputReference {
	var returns MskClusterEncryptionInfoOutputReference
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) EncryptionInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) EnhancedMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) KafkaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) LoggingInfo() MskClusterLoggingInfoOutputReference {
	var returns MskClusterLoggingInfoOutputReference
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) LoggingInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) NumberOfBrokerNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) OpenMonitoring() MskClusterOpenMonitoringOutputReference {
	var returns MskClusterOpenMonitoringOutputReference
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) OpenMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) StorageMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) StorageModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster awscc_msk_cluster} Resource.
func NewMskCluster(scope constructs.Construct, id *string, config *MskClusterConfig) MskCluster {
	_init_.Initialize()

	if err := validateNewMskClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskCluster{}

	_jsii_.Create(
		"awscc.mskCluster.MskCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster awscc_msk_cluster} Resource.
func NewMskCluster_Override(m MskCluster, scope constructs.Construct, id *string, config *MskClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mskCluster.MskCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MskCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetCurrentVersion(val *string) {
	if err := j.validateSetCurrentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentVersion",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetEnhancedMonitoring(val *string) {
	if err := j.validateSetEnhancedMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetKafkaVersion(val *string) {
	if err := j.validateSetKafkaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetNumberOfBrokerNodes(val *float64) {
	if err := j.validateSetNumberOfBrokerNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetStorageMode(val *string) {
	if err := j.validateSetStorageModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMode",
		val,
	)
}

func (j *jsiiProxy_MskCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a MskCluster resource upon running "cdktf plan <stack-name>".
func MskCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMskCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mskCluster.MskCluster",
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
func MskCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mskCluster.MskCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MskCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mskCluster.MskCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MskCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mskCluster.MskCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MskCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mskCluster.MskCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MskCluster) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MskCluster) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MskCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MskCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MskCluster) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MskCluster) PutBrokerNodeGroupInfo(value *MskClusterBrokerNodeGroupInfo) {
	if err := m.validatePutBrokerNodeGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBrokerNodeGroupInfo",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) PutClientAuthentication(value *MskClusterClientAuthentication) {
	if err := m.validatePutClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientAuthentication",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) PutConfigurationInfo(value *MskClusterConfigurationInfo) {
	if err := m.validatePutConfigurationInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putConfigurationInfo",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) PutEncryptionInfo(value *MskClusterEncryptionInfo) {
	if err := m.validatePutEncryptionInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionInfo",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) PutLoggingInfo(value *MskClusterLoggingInfo) {
	if err := m.validatePutLoggingInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLoggingInfo",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) PutOpenMonitoring(value *MskClusterOpenMonitoring) {
	if err := m.validatePutOpenMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOpenMonitoring",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskCluster) ResetClientAuthentication() {
	_jsii_.InvokeVoid(
		m,
		"resetClientAuthentication",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetConfigurationInfo() {
	_jsii_.InvokeVoid(
		m,
		"resetConfigurationInfo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetCurrentVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetCurrentVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetEncryptionInfo() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionInfo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetEnhancedMonitoring() {
	_jsii_.InvokeVoid(
		m,
		"resetEnhancedMonitoring",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetLoggingInfo() {
	_jsii_.InvokeVoid(
		m,
		"resetLoggingInfo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetOpenMonitoring() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenMonitoring",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetStorageMode() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

