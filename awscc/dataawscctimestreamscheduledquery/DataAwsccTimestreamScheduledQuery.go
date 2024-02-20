package dataawscctimestreamscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscctimestreamscheduledquery/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/timestream_scheduled_query awscc_timestream_scheduled_query}.
type DataAwsccTimestreamScheduledQuery interface {
	cdktf.TerraformDataSource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientToken() *string
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
	ErrorReportConfiguration() DataAwsccTimestreamScheduledQueryErrorReportConfigurationOutputReference
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
	KmsKeyId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() DataAwsccTimestreamScheduledQueryNotificationConfigurationOutputReference
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	QueryString() *string
	// Experimental.
	RawOverrides() interface{}
	ScheduleConfiguration() DataAwsccTimestreamScheduledQueryScheduleConfigurationOutputReference
	ScheduledQueryExecutionRoleArn() *string
	ScheduledQueryName() *string
	SqErrorReportConfiguration() *string
	SqKmsKeyId() *string
	SqName() *string
	SqNotificationConfiguration() *string
	SqQueryString() *string
	SqScheduleConfiguration() *string
	SqScheduledQueryExecutionRoleArn() *string
	SqTargetConfiguration() *string
	Tags() DataAwsccTimestreamScheduledQueryTagsList
	TargetConfiguration() DataAwsccTimestreamScheduledQueryTargetConfigurationOutputReference
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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

// The jsii proxy struct for DataAwsccTimestreamScheduledQuery
type jsiiProxy_DataAwsccTimestreamScheduledQuery struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ErrorReportConfiguration() DataAwsccTimestreamScheduledQueryErrorReportConfigurationOutputReference {
	var returns DataAwsccTimestreamScheduledQueryErrorReportConfigurationOutputReference
	_jsii_.Get(
		j,
		"errorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) NotificationConfiguration() DataAwsccTimestreamScheduledQueryNotificationConfigurationOutputReference {
	var returns DataAwsccTimestreamScheduledQueryNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ScheduleConfiguration() DataAwsccTimestreamScheduledQueryScheduleConfigurationOutputReference {
	var returns DataAwsccTimestreamScheduledQueryScheduleConfigurationOutputReference
	_jsii_.Get(
		j,
		"scheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) ScheduledQueryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqErrorReportConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqErrorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqNotificationConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqNotificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqScheduleConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqScheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqScheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) SqTargetConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqTargetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) Tags() DataAwsccTimestreamScheduledQueryTagsList {
	var returns DataAwsccTimestreamScheduledQueryTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) TargetConfiguration() DataAwsccTimestreamScheduledQueryTargetConfigurationOutputReference {
	var returns DataAwsccTimestreamScheduledQueryTargetConfigurationOutputReference
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/timestream_scheduled_query awscc_timestream_scheduled_query} Data Source.
func NewDataAwsccTimestreamScheduledQuery(scope constructs.Construct, id *string, config *DataAwsccTimestreamScheduledQueryConfig) DataAwsccTimestreamScheduledQuery {
	_init_.Initialize()

	if err := validateNewDataAwsccTimestreamScheduledQueryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccTimestreamScheduledQuery{}

	_jsii_.Create(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/timestream_scheduled_query awscc_timestream_scheduled_query} Data Source.
func NewDataAwsccTimestreamScheduledQuery_Override(d DataAwsccTimestreamScheduledQuery, scope constructs.Construct, id *string, config *DataAwsccTimestreamScheduledQueryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTimestreamScheduledQuery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccTimestreamScheduledQuery resource upon running "cdktf plan <stack-name>".
func DataAwsccTimestreamScheduledQuery_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccTimestreamScheduledQuery_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
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
func DataAwsccTimestreamScheduledQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccTimestreamScheduledQuery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccTimestreamScheduledQuery_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccTimestreamScheduledQuery_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccTimestreamScheduledQuery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccTimestreamScheduledQuery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccTimestreamScheduledQuery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccTimestreamScheduledQuery.DataAwsccTimestreamScheduledQuery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccTimestreamScheduledQuery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

