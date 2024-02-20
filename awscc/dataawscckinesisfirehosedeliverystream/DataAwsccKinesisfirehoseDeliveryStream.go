package dataawscckinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscckinesisfirehosedeliverystream/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream}.
type DataAwsccKinesisfirehoseDeliveryStream interface {
	cdktf.TerraformDataSource
	AmazonOpenSearchServerlessDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference
	AmazonopensearchserviceDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeliveryStreamEncryptionConfigurationInput() DataAwsccKinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference
	DeliveryStreamName() *string
	DeliveryStreamType() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ElasticsearchDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference
	ExtendedS3DestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpEndpointDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference
	Id() *string
	SetId(val *string)
	IdInput() *string
	KinesisStreamSourceConfiguration() DataAwsccKinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MskSourceConfiguration() DataAwsccKinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RedshiftDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference
	S3DestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference
	SnowflakeDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference
	SplunkDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference
	Tags() DataAwsccKinesisfirehoseDeliveryStreamTagsList
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

// The jsii proxy struct for DataAwsccKinesisfirehoseDeliveryStream
type jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) AmazonOpenSearchServerlessDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchServerlessDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) AmazonopensearchserviceDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) DeliveryStreamEncryptionConfigurationInput() DataAwsccKinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) DeliveryStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ElasticsearchDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ExtendedS3DestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"extendedS3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) HttpEndpointDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) KinesisStreamSourceConfiguration() DataAwsccKinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) MskSourceConfiguration() DataAwsccKinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) RedshiftDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"redshiftDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) S3DestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) SnowflakeDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"snowflakeDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) SplunkDestinationConfiguration() DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference {
	var returns DataAwsccKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"splunkDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) Tags() DataAwsccKinesisfirehoseDeliveryStreamTagsList {
	var returns DataAwsccKinesisfirehoseDeliveryStreamTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream} Data Source.
func NewDataAwsccKinesisfirehoseDeliveryStream(scope constructs.Construct, id *string, config *DataAwsccKinesisfirehoseDeliveryStreamConfig) DataAwsccKinesisfirehoseDeliveryStream {
	_init_.Initialize()

	if err := validateNewDataAwsccKinesisfirehoseDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream{}

	_jsii_.Create(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream} Data Source.
func NewDataAwsccKinesisfirehoseDeliveryStream_Override(d DataAwsccKinesisfirehoseDeliveryStream, scope constructs.Construct, id *string, config *DataAwsccKinesisfirehoseDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccKinesisfirehoseDeliveryStream resource upon running "cdktf plan <stack-name>".
func DataAwsccKinesisfirehoseDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccKinesisfirehoseDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
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
func DataAwsccKinesisfirehoseDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKinesisfirehoseDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccKinesisfirehoseDeliveryStream_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKinesisfirehoseDeliveryStream_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccKinesisfirehoseDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccKinesisfirehoseDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccKinesisfirehoseDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccKinesisfirehoseDeliveryStream.DataAwsccKinesisfirehoseDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKinesisfirehoseDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

