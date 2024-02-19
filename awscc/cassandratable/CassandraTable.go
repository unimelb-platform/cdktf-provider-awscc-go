package cassandratable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cassandratable/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table awscc_cassandra_table}.
type CassandraTable interface {
	cdktf.TerraformResource
	AutoScalingSpecifications() CassandraTableAutoScalingSpecificationsOutputReference
	AutoScalingSpecificationsInput() interface{}
	BillingMode() CassandraTableBillingModeOutputReference
	BillingModeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientSideTimestampsEnabled() interface{}
	SetClientSideTimestampsEnabled(val interface{})
	ClientSideTimestampsEnabledInput() interface{}
	ClusteringKeyColumns() CassandraTableClusteringKeyColumnsList
	ClusteringKeyColumnsInput() interface{}
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
	DefaultTimeToLive() *float64
	SetDefaultTimeToLive(val *float64)
	DefaultTimeToLiveInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionSpecification() CassandraTableEncryptionSpecificationOutputReference
	EncryptionSpecificationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KeyspaceName() *string
	SetKeyspaceName(val *string)
	KeyspaceNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PartitionKeyColumns() CassandraTablePartitionKeyColumnsList
	PartitionKeyColumnsInput() interface{}
	PointInTimeRecoveryEnabled() interface{}
	SetPointInTimeRecoveryEnabled(val interface{})
	PointInTimeRecoveryEnabledInput() interface{}
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
	RegularColumns() CassandraTableRegularColumnsList
	RegularColumnsInput() interface{}
	ReplicaSpecifications() CassandraTableReplicaSpecificationsList
	ReplicaSpecificationsInput() interface{}
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	Tags() CassandraTableTagsList
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
	PutAutoScalingSpecifications(value *CassandraTableAutoScalingSpecifications)
	PutBillingMode(value *CassandraTableBillingMode)
	PutClusteringKeyColumns(value interface{})
	PutEncryptionSpecification(value *CassandraTableEncryptionSpecification)
	PutPartitionKeyColumns(value interface{})
	PutRegularColumns(value interface{})
	PutReplicaSpecifications(value interface{})
	PutTags(value interface{})
	ResetAutoScalingSpecifications()
	ResetBillingMode()
	ResetClientSideTimestampsEnabled()
	ResetClusteringKeyColumns()
	ResetDefaultTimeToLive()
	ResetEncryptionSpecification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecoveryEnabled()
	ResetRegularColumns()
	ResetReplicaSpecifications()
	ResetTableName()
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

// The jsii proxy struct for CassandraTable
type jsiiProxy_CassandraTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CassandraTable) AutoScalingSpecifications() CassandraTableAutoScalingSpecificationsOutputReference {
	var returns CassandraTableAutoScalingSpecificationsOutputReference
	_jsii_.Get(
		j,
		"autoScalingSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) AutoScalingSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) BillingMode() CassandraTableBillingModeOutputReference {
	var returns CassandraTableBillingModeOutputReference
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) BillingModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ClientSideTimestampsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSideTimestampsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ClientSideTimestampsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSideTimestampsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ClusteringKeyColumns() CassandraTableClusteringKeyColumnsList {
	var returns CassandraTableClusteringKeyColumnsList
	_jsii_.Get(
		j,
		"clusteringKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ClusteringKeyColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteringKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) DefaultTimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) DefaultTimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) EncryptionSpecification() CassandraTableEncryptionSpecificationOutputReference {
	var returns CassandraTableEncryptionSpecificationOutputReference
	_jsii_.Get(
		j,
		"encryptionSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) EncryptionSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) KeyspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) KeyspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) PartitionKeyColumns() CassandraTablePartitionKeyColumnsList {
	var returns CassandraTablePartitionKeyColumnsList
	_jsii_.Get(
		j,
		"partitionKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) PartitionKeyColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) PointInTimeRecoveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) PointInTimeRecoveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) RegularColumns() CassandraTableRegularColumnsList {
	var returns CassandraTableRegularColumnsList
	_jsii_.Get(
		j,
		"regularColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) RegularColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regularColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ReplicaSpecifications() CassandraTableReplicaSpecificationsList {
	var returns CassandraTableReplicaSpecificationsList
	_jsii_.Get(
		j,
		"replicaSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) ReplicaSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) Tags() CassandraTableTagsList {
	var returns CassandraTableTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table awscc_cassandra_table} Resource.
func NewCassandraTable(scope constructs.Construct, id *string, config *CassandraTableConfig) CassandraTable {
	_init_.Initialize()

	if err := validateNewCassandraTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CassandraTable{}

	_jsii_.Create(
		"awscc.cassandraTable.CassandraTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table awscc_cassandra_table} Resource.
func NewCassandraTable_Override(c CassandraTable, scope constructs.Construct, id *string, config *CassandraTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cassandraTable.CassandraTable",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CassandraTable)SetClientSideTimestampsEnabled(val interface{}) {
	if err := j.validateSetClientSideTimestampsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSideTimestampsEnabled",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetDefaultTimeToLive(val *float64) {
	if err := j.validateSetDefaultTimeToLiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeToLive",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetKeyspaceName(val *string) {
	if err := j.validateSetKeyspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyspaceName",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetPointInTimeRecoveryEnabled(val interface{}) {
	if err := j.validateSetPointInTimeRecoveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRecoveryEnabled",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CassandraTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTF code for importing a CassandraTable resource upon running "cdktf plan <stack-name>".
func CassandraTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCassandraTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cassandraTable.CassandraTable",
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
func CassandraTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCassandraTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cassandraTable.CassandraTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CassandraTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCassandraTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cassandraTable.CassandraTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CassandraTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCassandraTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cassandraTable.CassandraTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CassandraTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cassandraTable.CassandraTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CassandraTable) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CassandraTable) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CassandraTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CassandraTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CassandraTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CassandraTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CassandraTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CassandraTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CassandraTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CassandraTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CassandraTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CassandraTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CassandraTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CassandraTable) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CassandraTable) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CassandraTable) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CassandraTable) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CassandraTable) PutAutoScalingSpecifications(value *CassandraTableAutoScalingSpecifications) {
	if err := c.validatePutAutoScalingSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoScalingSpecifications",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutBillingMode(value *CassandraTableBillingMode) {
	if err := c.validatePutBillingModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBillingMode",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutClusteringKeyColumns(value interface{}) {
	if err := c.validatePutClusteringKeyColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClusteringKeyColumns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutEncryptionSpecification(value *CassandraTableEncryptionSpecification) {
	if err := c.validatePutEncryptionSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionSpecification",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutPartitionKeyColumns(value interface{}) {
	if err := c.validatePutPartitionKeyColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPartitionKeyColumns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutRegularColumns(value interface{}) {
	if err := c.validatePutRegularColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegularColumns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutReplicaSpecifications(value interface{}) {
	if err := c.validatePutReplicaSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReplicaSpecifications",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTable) ResetAutoScalingSpecifications() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingSpecifications",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetClientSideTimestampsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSideTimestampsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetClusteringKeyColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetClusteringKeyColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetDefaultTimeToLive() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTimeToLive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetEncryptionSpecification() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionSpecification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetPointInTimeRecoveryEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPointInTimeRecoveryEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetRegularColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetRegularColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetReplicaSpecifications() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicaSpecifications",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetTableName() {
	_jsii_.InvokeVoid(
		c,
		"resetTableName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

