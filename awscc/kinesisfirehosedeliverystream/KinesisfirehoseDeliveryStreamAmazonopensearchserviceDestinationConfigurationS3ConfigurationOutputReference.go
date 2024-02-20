package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
)

type KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference interface {
	cdktf.ComplexObject
	BucketArn() *string
	SetBucketArn(val *string)
	BucketArnInput() *string
	BufferingHints() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference
	BufferingHintsInput() interface{}
	CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CompressionFormat() *string
	SetCompressionFormat(val *string)
	CompressionFormatInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionConfiguration() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() interface{}
	ErrorOutputPrefix() *string
	SetErrorOutputPrefix(val *string)
	ErrorOutputPrefixInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration
	SetInternalValue(val *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBufferingHints(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHints)
	PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions)
	PutEncryptionConfiguration(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfiguration)
	ResetBufferingHints()
	ResetCloudwatchLoggingOptions()
	ResetCompressionFormat()
	ResetEncryptionConfiguration()
	ResetErrorOutputPrefix()
	ResetPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) BufferingHints() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference {
	var returns KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHintsOutputReference
	_jsii_.Get(
		j,
		"bufferingHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) BufferingHintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferingHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) CloudwatchLoggingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) EncryptionConfiguration() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) EncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) InternalValue() *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration {
	var returns *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetErrorOutputPrefix(val *string) {
	if err := j.validateSetErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetInternalValue(val *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3Configuration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) PutBufferingHints(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationBufferingHints) {
	if err := k.validatePutBufferingHintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putBufferingHints",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) PutEncryptionConfiguration(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationEncryptionConfiguration) {
	if err := k.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetBufferingHints() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingHints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetErrorOutputPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationS3ConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

