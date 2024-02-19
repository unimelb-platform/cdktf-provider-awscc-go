package timestreamscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/timestreamscheduledquery/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query awscc_timestream_scheduled_query}.
type TimestreamScheduledQuery interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientToken() *string
	SetClientToken(val *string)
	ClientTokenInput() *string
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
	ErrorReportConfiguration() TimestreamScheduledQueryErrorReportConfigurationOutputReference
	ErrorReportConfigurationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() TimestreamScheduledQueryNotificationConfigurationOutputReference
	NotificationConfigurationInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	// Experimental.
	RawOverrides() interface{}
	ScheduleConfiguration() TimestreamScheduledQueryScheduleConfigurationOutputReference
	ScheduleConfigurationInput() interface{}
	ScheduledQueryExecutionRoleArn() *string
	SetScheduledQueryExecutionRoleArn(val *string)
	ScheduledQueryExecutionRoleArnInput() *string
	ScheduledQueryName() *string
	SetScheduledQueryName(val *string)
	ScheduledQueryNameInput() *string
	SqErrorReportConfiguration() *string
	SqKmsKeyId() *string
	SqName() *string
	SqNotificationConfiguration() *string
	SqQueryString() *string
	SqScheduleConfiguration() *string
	SqScheduledQueryExecutionRoleArn() *string
	SqTargetConfiguration() *string
	Tags() TimestreamScheduledQueryTagsList
	TagsInput() interface{}
	TargetConfiguration() TimestreamScheduledQueryTargetConfigurationOutputReference
	TargetConfigurationInput() interface{}
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
	PutErrorReportConfiguration(value *TimestreamScheduledQueryErrorReportConfiguration)
	PutNotificationConfiguration(value *TimestreamScheduledQueryNotificationConfiguration)
	PutScheduleConfiguration(value *TimestreamScheduledQueryScheduleConfiguration)
	PutTags(value interface{})
	PutTargetConfiguration(value *TimestreamScheduledQueryTargetConfiguration)
	ResetClientToken()
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheduledQueryName()
	ResetTags()
	ResetTargetConfiguration()
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

// The jsii proxy struct for TimestreamScheduledQuery
type jsiiProxy_TimestreamScheduledQuery struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TimestreamScheduledQuery) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ClientTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ErrorReportConfiguration() TimestreamScheduledQueryErrorReportConfigurationOutputReference {
	var returns TimestreamScheduledQueryErrorReportConfigurationOutputReference
	_jsii_.Get(
		j,
		"errorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ErrorReportConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) NotificationConfiguration() TimestreamScheduledQueryNotificationConfigurationOutputReference {
	var returns TimestreamScheduledQueryNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) NotificationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduleConfiguration() TimestreamScheduledQueryScheduleConfigurationOutputReference {
	var returns TimestreamScheduledQueryScheduleConfigurationOutputReference
	_jsii_.Get(
		j,
		"scheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduledQueryExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduledQueryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) ScheduledQueryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqErrorReportConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqErrorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqNotificationConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqNotificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqScheduleConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqScheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqScheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) SqTargetConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqTargetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) Tags() TimestreamScheduledQueryTagsList {
	var returns TimestreamScheduledQueryTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TargetConfiguration() TimestreamScheduledQueryTargetConfigurationOutputReference {
	var returns TimestreamScheduledQueryTargetConfigurationOutputReference
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TargetConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQuery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query awscc_timestream_scheduled_query} Resource.
func NewTimestreamScheduledQuery(scope constructs.Construct, id *string, config *TimestreamScheduledQueryConfig) TimestreamScheduledQuery {
	_init_.Initialize()

	if err := validateNewTimestreamScheduledQueryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamScheduledQuery{}

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query awscc_timestream_scheduled_query} Resource.
func NewTimestreamScheduledQuery_Override(t TimestreamScheduledQuery, scope constructs.Construct, id *string, config *TimestreamScheduledQueryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetClientToken(val *string) {
	if err := j.validateSetClientTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetScheduledQueryExecutionRoleArn(val *string) {
	if err := j.validateSetScheduledQueryExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledQueryExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQuery)SetScheduledQueryName(val *string) {
	if err := j.validateSetScheduledQueryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledQueryName",
		val,
	)
}

// Generates CDKTF code for importing a TimestreamScheduledQuery resource upon running "cdktf plan <stack-name>".
func TimestreamScheduledQuery_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTimestreamScheduledQuery_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
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
func TimestreamScheduledQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamScheduledQuery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamScheduledQuery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamScheduledQuery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamScheduledQuery_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamScheduledQuery_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TimestreamScheduledQuery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQuery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) PutErrorReportConfiguration(value *TimestreamScheduledQueryErrorReportConfiguration) {
	if err := t.validatePutErrorReportConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putErrorReportConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) PutNotificationConfiguration(value *TimestreamScheduledQueryNotificationConfiguration) {
	if err := t.validatePutNotificationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotificationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) PutScheduleConfiguration(value *TimestreamScheduledQueryScheduleConfiguration) {
	if err := t.validatePutScheduleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScheduleConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) PutTags(value interface{}) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) PutTargetConfiguration(value *TimestreamScheduledQueryTargetConfiguration) {
	if err := t.validatePutTargetConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetClientToken() {
	_jsii_.InvokeVoid(
		t,
		"resetClientToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetScheduledQueryName() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduledQueryName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) ResetTargetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamScheduledQuery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQuery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

