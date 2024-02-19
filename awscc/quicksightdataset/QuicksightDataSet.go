package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksightdataset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set awscc_quicksight_data_set}.
type QuicksightDataSet interface {
	cdktf.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ColumnGroups() QuicksightDataSetColumnGroupsList
	ColumnGroupsInput() interface{}
	ColumnLevelPermissionRules() QuicksightDataSetColumnLevelPermissionRulesList
	ColumnLevelPermissionRulesInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumedSpiceCapacityInBytes() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedTime() *string
	DataSetId() *string
	SetDataSetId(val *string)
	DataSetIdInput() *string
	DataSetUsageConfiguration() QuicksightDataSetDataSetUsageConfigurationOutputReference
	DataSetUsageConfigurationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FieldFolders() QuicksightDataSetFieldFoldersMap
	FieldFoldersInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	ImportMode() *string
	SetImportMode(val *string)
	ImportModeInput() *string
	IngestionWaitPolicy() QuicksightDataSetIngestionWaitPolicyOutputReference
	IngestionWaitPolicyInput() interface{}
	LastUpdatedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogicalTableMap() QuicksightDataSetLogicalTableMapMap
	LogicalTableMapInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputColumns() QuicksightDataSetOutputColumnsList
	Permissions() QuicksightDataSetPermissionsList
	PermissionsInput() interface{}
	PhysicalTableMap() QuicksightDataSetPhysicalTableMapMap
	PhysicalTableMapInput() interface{}
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
	RowLevelPermissionDataSet() QuicksightDataSetRowLevelPermissionDataSetOutputReference
	RowLevelPermissionDataSetInput() interface{}
	Tags() QuicksightDataSetTagsList
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
	PutColumnGroups(value interface{})
	PutColumnLevelPermissionRules(value interface{})
	PutDataSetUsageConfiguration(value *QuicksightDataSetDataSetUsageConfiguration)
	PutFieldFolders(value interface{})
	PutIngestionWaitPolicy(value *QuicksightDataSetIngestionWaitPolicy)
	PutLogicalTableMap(value interface{})
	PutPermissions(value interface{})
	PutPhysicalTableMap(value interface{})
	PutRowLevelPermissionDataSet(value *QuicksightDataSetRowLevelPermissionDataSet)
	PutTags(value interface{})
	ResetAwsAccountId()
	ResetColumnGroups()
	ResetColumnLevelPermissionRules()
	ResetDataSetId()
	ResetDataSetUsageConfiguration()
	ResetFieldFolders()
	ResetImportMode()
	ResetIngestionWaitPolicy()
	ResetLogicalTableMap()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetPhysicalTableMap()
	ResetRowLevelPermissionDataSet()
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

// The jsii proxy struct for QuicksightDataSet
type jsiiProxy_QuicksightDataSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightDataSet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnGroups() QuicksightDataSetColumnGroupsList {
	var returns QuicksightDataSetColumnGroupsList
	_jsii_.Get(
		j,
		"columnGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnLevelPermissionRules() QuicksightDataSetColumnLevelPermissionRulesList {
	var returns QuicksightDataSetColumnLevelPermissionRulesList
	_jsii_.Get(
		j,
		"columnLevelPermissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnLevelPermissionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnLevelPermissionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ConsumedSpiceCapacityInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consumedSpiceCapacityInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetUsageConfiguration() QuicksightDataSetDataSetUsageConfigurationOutputReference {
	var returns QuicksightDataSetDataSetUsageConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataSetUsageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetUsageConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetUsageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FieldFolders() QuicksightDataSetFieldFoldersMap {
	var returns QuicksightDataSetFieldFoldersMap
	_jsii_.Get(
		j,
		"fieldFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FieldFoldersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ImportMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ImportModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) IngestionWaitPolicy() QuicksightDataSetIngestionWaitPolicyOutputReference {
	var returns QuicksightDataSetIngestionWaitPolicyOutputReference
	_jsii_.Get(
		j,
		"ingestionWaitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) IngestionWaitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingestionWaitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LogicalTableMap() QuicksightDataSetLogicalTableMapMap {
	var returns QuicksightDataSetLogicalTableMapMap
	_jsii_.Get(
		j,
		"logicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LogicalTableMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicalTableMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) OutputColumns() QuicksightDataSetOutputColumnsList {
	var returns QuicksightDataSetOutputColumnsList
	_jsii_.Get(
		j,
		"outputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Permissions() QuicksightDataSetPermissionsList {
	var returns QuicksightDataSetPermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PhysicalTableMap() QuicksightDataSetPhysicalTableMapMap {
	var returns QuicksightDataSetPhysicalTableMapMap
	_jsii_.Get(
		j,
		"physicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PhysicalTableMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalTableMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RowLevelPermissionDataSet() QuicksightDataSetRowLevelPermissionDataSetOutputReference {
	var returns QuicksightDataSetRowLevelPermissionDataSetOutputReference
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RowLevelPermissionDataSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Tags() QuicksightDataSetTagsList {
	var returns QuicksightDataSetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set awscc_quicksight_data_set} Resource.
func NewQuicksightDataSet(scope constructs.Construct, id *string, config *QuicksightDataSetConfig) QuicksightDataSet {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSet{}

	_jsii_.Create(
		"awscc.quicksightDataSet.QuicksightDataSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set awscc_quicksight_data_set} Resource.
func NewQuicksightDataSet_Override(q QuicksightDataSet, scope constructs.Construct, id *string, config *QuicksightDataSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightDataSet.QuicksightDataSet",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetDataSetId(val *string) {
	if err := j.validateSetDataSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSetId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetImportMode(val *string) {
	if err := j.validateSetImportModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importMode",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a QuicksightDataSet resource upon running "cdktf plan <stack-name>".
func QuicksightDataSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateQuicksightDataSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.quicksightDataSet.QuicksightDataSet",
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
func QuicksightDataSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightDataSet.QuicksightDataSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightDataSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightDataSet.QuicksightDataSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightDataSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.quicksightDataSet.QuicksightDataSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightDataSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.quicksightDataSet.QuicksightDataSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QuicksightDataSet) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QuicksightDataSet) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QuicksightDataSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightDataSet) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QuicksightDataSet) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightDataSet) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutColumnGroups(value interface{}) {
	if err := q.validatePutColumnGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putColumnGroups",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutColumnLevelPermissionRules(value interface{}) {
	if err := q.validatePutColumnLevelPermissionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putColumnLevelPermissionRules",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutDataSetUsageConfiguration(value *QuicksightDataSetDataSetUsageConfiguration) {
	if err := q.validatePutDataSetUsageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataSetUsageConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutFieldFolders(value interface{}) {
	if err := q.validatePutFieldFoldersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFieldFolders",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutIngestionWaitPolicy(value *QuicksightDataSetIngestionWaitPolicy) {
	if err := q.validatePutIngestionWaitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIngestionWaitPolicy",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutLogicalTableMap(value interface{}) {
	if err := q.validatePutLogicalTableMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLogicalTableMap",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutPermissions(value interface{}) {
	if err := q.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPermissions",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutPhysicalTableMap(value interface{}) {
	if err := q.validatePutPhysicalTableMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPhysicalTableMap",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutRowLevelPermissionDataSet(value *QuicksightDataSetRowLevelPermissionDataSet) {
	if err := q.validatePutRowLevelPermissionDataSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRowLevelPermissionDataSet",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetColumnGroups() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnGroups",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetColumnLevelPermissionRules() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnLevelPermissionRules",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetDataSetId() {
	_jsii_.InvokeVoid(
		q,
		"resetDataSetId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetDataSetUsageConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetDataSetUsageConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetFieldFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetFieldFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetImportMode() {
	_jsii_.InvokeVoid(
		q,
		"resetImportMode",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetIngestionWaitPolicy() {
	_jsii_.InvokeVoid(
		q,
		"resetIngestionWaitPolicy",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetLogicalTableMap() {
	_jsii_.InvokeVoid(
		q,
		"resetLogicalTableMap",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetName() {
	_jsii_.InvokeVoid(
		q,
		"resetName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetPermissions() {
	_jsii_.InvokeVoid(
		q,
		"resetPermissions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetPhysicalTableMap() {
	_jsii_.InvokeVoid(
		q,
		"resetPhysicalTableMap",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetRowLevelPermissionDataSet() {
	_jsii_.InvokeVoid(
		q,
		"resetRowLevelPermissionDataSet",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

