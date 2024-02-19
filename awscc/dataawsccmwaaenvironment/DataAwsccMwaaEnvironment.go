package dataawsccmwaaenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccmwaaenvironment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/mwaa_environment awscc_mwaa_environment}.
type DataAwsccMwaaEnvironment interface {
	cdktf.TerraformDataSource
	AirflowConfigurationOptions() *string
	AirflowVersion() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CeleryExecutorQueue() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DagS3Path() *string
	DatabaseVpcEndpointService() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointManagement() *string
	EnvironmentClass() *string
	ExecutionRoleArn() *string
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
	KmsKey() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfiguration() DataAwsccMwaaEnvironmentLoggingConfigurationOutputReference
	MaxWorkers() *float64
	MinWorkers() *float64
	Name() *string
	NetworkConfiguration() DataAwsccMwaaEnvironmentNetworkConfigurationOutputReference
	// The tree node.
	Node() constructs.Node
	PluginsS3ObjectVersion() *string
	PluginsS3Path() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RequirementsS3ObjectVersion() *string
	RequirementsS3Path() *string
	Schedulers() *float64
	SourceBucketArn() *string
	StartupScriptS3ObjectVersion() *string
	StartupScriptS3Path() *string
	Tags() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebserverAccessMode() *string
	WebserverUrl() *string
	WebserverVpcEndpointService() *string
	WeeklyMaintenanceWindowStart() *string
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
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsccMwaaEnvironment
type jsiiProxy_DataAwsccMwaaEnvironment struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) AirflowConfigurationOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) AirflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) CeleryExecutorQueue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"celeryExecutorQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) DatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) EndpointManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) EnvironmentClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) LoggingConfiguration() DataAwsccMwaaEnvironmentLoggingConfigurationOutputReference {
	var returns DataAwsccMwaaEnvironmentLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) NetworkConfiguration() DataAwsccMwaaEnvironmentNetworkConfigurationOutputReference {
	var returns DataAwsccMwaaEnvironmentNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) SourceBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) Tags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) WebserverAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) WebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/mwaa_environment awscc_mwaa_environment} Data Source.
func NewDataAwsccMwaaEnvironment(scope constructs.Construct, id *string, config *DataAwsccMwaaEnvironmentConfig) DataAwsccMwaaEnvironment {
	_init_.Initialize()

	if err := validateNewDataAwsccMwaaEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMwaaEnvironment{}

	_jsii_.Create(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/mwaa_environment awscc_mwaa_environment} Data Source.
func NewDataAwsccMwaaEnvironment_Override(d DataAwsccMwaaEnvironment, scope constructs.Construct, id *string, config *DataAwsccMwaaEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMwaaEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccMwaaEnvironment resource upon running "cdktf plan <stack-name>".
func DataAwsccMwaaEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccMwaaEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
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
func DataAwsccMwaaEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccMwaaEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccMwaaEnvironment_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccMwaaEnvironment_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccMwaaEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccMwaaEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccMwaaEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccMwaaEnvironment.DataAwsccMwaaEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMwaaEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMwaaEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

