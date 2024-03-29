package dynamodbglobaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dynamodbglobaltable/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table}.
type DynamodbGlobalTable interface {
	cdktf.TerraformResource
	Arn() *string
	AttributeDefinitions() DynamodbGlobalTableAttributeDefinitionsList
	AttributeDefinitionsInput() interface{}
	BillingMode() *string
	SetBillingMode(val *string)
	BillingModeInput() *string
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
	GlobalSecondaryIndexes() DynamodbGlobalTableGlobalSecondaryIndexesList
	GlobalSecondaryIndexesInput() interface{}
	Id() *string
	KeySchema() DynamodbGlobalTableKeySchemaList
	KeySchemaInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalSecondaryIndexes() DynamodbGlobalTableLocalSecondaryIndexesList
	LocalSecondaryIndexesInput() interface{}
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
	Replicas() DynamodbGlobalTableReplicasList
	ReplicasInput() interface{}
	SseSpecification() DynamodbGlobalTableSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	StreamArn() *string
	StreamSpecification() DynamodbGlobalTableStreamSpecificationOutputReference
	StreamSpecificationInput() interface{}
	TableId() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeToLiveSpecification() DynamodbGlobalTableTimeToLiveSpecificationOutputReference
	TimeToLiveSpecificationInput() interface{}
	WriteProvisionedThroughputSettings() DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference
	WriteProvisionedThroughputSettingsInput() interface{}
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
	PutAttributeDefinitions(value interface{})
	PutGlobalSecondaryIndexes(value interface{})
	PutKeySchema(value interface{})
	PutLocalSecondaryIndexes(value interface{})
	PutReplicas(value interface{})
	PutSseSpecification(value *DynamodbGlobalTableSseSpecification)
	PutStreamSpecification(value *DynamodbGlobalTableStreamSpecification)
	PutTimeToLiveSpecification(value *DynamodbGlobalTableTimeToLiveSpecification)
	PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableWriteProvisionedThroughputSettings)
	ResetBillingMode()
	ResetGlobalSecondaryIndexes()
	ResetLocalSecondaryIndexes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSseSpecification()
	ResetStreamSpecification()
	ResetTableName()
	ResetTimeToLiveSpecification()
	ResetWriteProvisionedThroughputSettings()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DynamodbGlobalTable
type jsiiProxy_DynamodbGlobalTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DynamodbGlobalTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) AttributeDefinitions() DynamodbGlobalTableAttributeDefinitionsList {
	var returns DynamodbGlobalTableAttributeDefinitionsList
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) AttributeDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalSecondaryIndexes() DynamodbGlobalTableGlobalSecondaryIndexesList {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesList
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) KeySchema() DynamodbGlobalTableKeySchemaList {
	var returns DynamodbGlobalTableKeySchemaList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) LocalSecondaryIndexes() DynamodbGlobalTableLocalSecondaryIndexesList {
	var returns DynamodbGlobalTableLocalSecondaryIndexesList
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) LocalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Replicas() DynamodbGlobalTableReplicasList {
	var returns DynamodbGlobalTableReplicasList
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReplicasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) SseSpecification() DynamodbGlobalTableSseSpecificationOutputReference {
	var returns DynamodbGlobalTableSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamSpecification() DynamodbGlobalTableStreamSpecificationOutputReference {
	var returns DynamodbGlobalTableStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TimeToLiveSpecification() DynamodbGlobalTableTimeToLiveSpecificationOutputReference {
	var returns DynamodbGlobalTableTimeToLiveSpecificationOutputReference
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TimeToLiveSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteProvisionedThroughputSettings() DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table} Resource.
func NewDynamodbGlobalTable(scope constructs.Construct, id *string, config *DynamodbGlobalTableConfig) DynamodbGlobalTable {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalTable{}

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table} Resource.
func NewDynamodbGlobalTable_Override(d DynamodbGlobalTable, scope constructs.Construct, id *string, config *DynamodbGlobalTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTF code for importing a DynamodbGlobalTable resource upon running "cdktf plan <stack-name>".
func DynamodbGlobalTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
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
func DynamodbGlobalTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbGlobalTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTable) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutAttributeDefinitions(value interface{}) {
	if err := d.validatePutAttributeDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributeDefinitions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutGlobalSecondaryIndexes(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutKeySchema(value interface{}) {
	if err := d.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutLocalSecondaryIndexes(value interface{}) {
	if err := d.validatePutLocalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutReplicas(value interface{}) {
	if err := d.validatePutReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReplicas",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutSseSpecification(value *DynamodbGlobalTableSseSpecification) {
	if err := d.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutStreamSpecification(value *DynamodbGlobalTableStreamSpecification) {
	if err := d.validatePutStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutTimeToLiveSpecification(value *DynamodbGlobalTableTimeToLiveSpecification) {
	if err := d.validatePutTimeToLiveSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeToLiveSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableWriteProvisionedThroughputSettings) {
	if err := d.validatePutWriteProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWriteProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetGlobalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetLocalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetTimeToLiveSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeToLiveSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetWriteProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

