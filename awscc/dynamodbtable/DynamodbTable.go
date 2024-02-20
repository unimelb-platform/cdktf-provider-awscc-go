package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dynamodbtable/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table awscc_dynamodb_table}.
type DynamodbTable interface {
	cdktf.TerraformResource
	Arn() *string
	AttributeDefinitions() DynamodbTableAttributeDefinitionsList
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
	ContributorInsightsSpecification() DynamodbTableContributorInsightsSpecificationOutputReference
	ContributorInsightsSpecificationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
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
	GlobalSecondaryIndexes() DynamodbTableGlobalSecondaryIndexesList
	GlobalSecondaryIndexesInput() interface{}
	Id() *string
	ImportSourceSpecification() DynamodbTableImportSourceSpecificationOutputReference
	ImportSourceSpecificationInput() interface{}
	KeySchema() *string
	SetKeySchema(val *string)
	KeySchemaInput() *string
	KinesisStreamSpecification() DynamodbTableKinesisStreamSpecificationOutputReference
	KinesisStreamSpecificationInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalSecondaryIndexes() DynamodbTableLocalSecondaryIndexesList
	LocalSecondaryIndexesInput() interface{}
	// The tree node.
	Node() constructs.Node
	PointInTimeRecoverySpecification() DynamodbTablePointInTimeRecoverySpecificationOutputReference
	PointInTimeRecoverySpecificationInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedThroughput() DynamodbTableProvisionedThroughputOutputReference
	ProvisionedThroughputInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SseSpecification() DynamodbTableSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	StreamArn() *string
	StreamSpecification() DynamodbTableStreamSpecificationOutputReference
	StreamSpecificationInput() interface{}
	TableClass() *string
	SetTableClass(val *string)
	TableClassInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	Tags() DynamodbTableTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeToLiveSpecification() DynamodbTableTimeToLiveSpecificationOutputReference
	TimeToLiveSpecificationInput() interface{}
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
	PutContributorInsightsSpecification(value *DynamodbTableContributorInsightsSpecification)
	PutGlobalSecondaryIndexes(value interface{})
	PutImportSourceSpecification(value *DynamodbTableImportSourceSpecification)
	PutKinesisStreamSpecification(value *DynamodbTableKinesisStreamSpecification)
	PutLocalSecondaryIndexes(value interface{})
	PutPointInTimeRecoverySpecification(value *DynamodbTablePointInTimeRecoverySpecification)
	PutProvisionedThroughput(value *DynamodbTableProvisionedThroughput)
	PutSseSpecification(value *DynamodbTableSseSpecification)
	PutStreamSpecification(value *DynamodbTableStreamSpecification)
	PutTags(value interface{})
	PutTimeToLiveSpecification(value *DynamodbTableTimeToLiveSpecification)
	ResetAttributeDefinitions()
	ResetBillingMode()
	ResetContributorInsightsSpecification()
	ResetDeletionProtectionEnabled()
	ResetGlobalSecondaryIndexes()
	ResetImportSourceSpecification()
	ResetKinesisStreamSpecification()
	ResetLocalSecondaryIndexes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecoverySpecification()
	ResetProvisionedThroughput()
	ResetSseSpecification()
	ResetStreamSpecification()
	ResetTableClass()
	ResetTableName()
	ResetTags()
	ResetTimeToLiveSpecification()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DynamodbTable
type jsiiProxy_DynamodbTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DynamodbTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) AttributeDefinitions() DynamodbTableAttributeDefinitionsList {
	var returns DynamodbTableAttributeDefinitionsList
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) AttributeDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ContributorInsightsSpecification() DynamodbTableContributorInsightsSpecificationOutputReference {
	var returns DynamodbTableContributorInsightsSpecificationOutputReference
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ContributorInsightsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndexes() DynamodbTableGlobalSecondaryIndexesList {
	var returns DynamodbTableGlobalSecondaryIndexesList
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportSourceSpecification() DynamodbTableImportSourceSpecificationOutputReference {
	var returns DynamodbTableImportSourceSpecificationOutputReference
	_jsii_.Get(
		j,
		"importSourceSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportSourceSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSourceSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KeySchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KeySchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KinesisStreamSpecification() DynamodbTableKinesisStreamSpecificationOutputReference {
	var returns DynamodbTableKinesisStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KinesisStreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndexes() DynamodbTableLocalSecondaryIndexesList {
	var returns DynamodbTableLocalSecondaryIndexesList
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecoverySpecification() DynamodbTablePointInTimeRecoverySpecificationOutputReference {
	var returns DynamodbTablePointInTimeRecoverySpecificationOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecoverySpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ProvisionedThroughput() DynamodbTableProvisionedThroughputOutputReference {
	var returns DynamodbTableProvisionedThroughputOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ProvisionedThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) SseSpecification() DynamodbTableSseSpecificationOutputReference {
	var returns DynamodbTableSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamSpecification() DynamodbTableStreamSpecificationOutputReference {
	var returns DynamodbTableStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Tags() DynamodbTableTagsList {
	var returns DynamodbTableTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TimeToLiveSpecification() DynamodbTableTimeToLiveSpecificationOutputReference {
	var returns DynamodbTableTimeToLiveSpecificationOutputReference
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TimeToLiveSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecificationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table awscc_dynamodb_table} Resource.
func NewDynamodbTable(scope constructs.Construct, id *string, config *DynamodbTableConfig) DynamodbTable {
	_init_.Initialize()

	if err := validateNewDynamodbTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTable{}

	_jsii_.Create(
		"awscc.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table awscc_dynamodb_table} Resource.
func NewDynamodbTable_Override(d DynamodbTable, scope constructs.Construct, id *string, config *DynamodbTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetKeySchema(val *string) {
	if err := j.validateSetKeySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTF code for importing a DynamodbTable resource upon running "cdktf plan <stack-name>".
func DynamodbTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dynamodbTable.DynamodbTable",
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
func DynamodbTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbTable.DynamodbTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbTable.DynamodbTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dynamodbTable.DynamodbTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dynamodbTable.DynamodbTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbTable) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbTable) PutAttributeDefinitions(value interface{}) {
	if err := d.validatePutAttributeDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributeDefinitions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutContributorInsightsSpecification(value *DynamodbTableContributorInsightsSpecification) {
	if err := d.validatePutContributorInsightsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContributorInsightsSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutGlobalSecondaryIndexes(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutImportSourceSpecification(value *DynamodbTableImportSourceSpecification) {
	if err := d.validatePutImportSourceSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImportSourceSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutKinesisStreamSpecification(value *DynamodbTableKinesisStreamSpecification) {
	if err := d.validatePutKinesisStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKinesisStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutLocalSecondaryIndexes(value interface{}) {
	if err := d.validatePutLocalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutPointInTimeRecoverySpecification(value *DynamodbTablePointInTimeRecoverySpecification) {
	if err := d.validatePutPointInTimeRecoverySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointInTimeRecoverySpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutProvisionedThroughput(value *DynamodbTableProvisionedThroughput) {
	if err := d.validatePutProvisionedThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisionedThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutSseSpecification(value *DynamodbTableSseSpecification) {
	if err := d.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutStreamSpecification(value *DynamodbTableStreamSpecification) {
	if err := d.validatePutStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTimeToLiveSpecification(value *DynamodbTableTimeToLiveSpecification) {
	if err := d.validatePutTimeToLiveSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeToLiveSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) ResetAttributeDefinitions() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributeDefinitions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetContributorInsightsSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetContributorInsightsSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetGlobalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetImportSourceSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetImportSourceSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetKinesisStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetLocalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetPointInTimeRecoverySpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetPointInTimeRecoverySpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTableClass() {
	_jsii_.InvokeVoid(
		d,
		"resetTableClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTimeToLiveSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeToLiveSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

