package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
)

type KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsEncryptionConfig() KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference
	KmsEncryptionConfigInput() interface{}
	NoEncryptionConfig() *string
	SetNoEncryptionConfig(val *string)
	NoEncryptionConfigInput() *string
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
	PutKmsEncryptionConfig(value *KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig)
	ResetKmsEncryptionConfig()
	ResetNoEncryptionConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) KmsEncryptionConfig() KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference {
	var returns KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"kmsEncryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) KmsEncryptionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) NoEncryptionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noEncryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) NoEncryptionConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noEncryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetNoEncryptionConfig(val *string) {
	if err := j.validateSetNoEncryptionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEncryptionConfig",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) PutKmsEncryptionConfig(value *KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationKmsEncryptionConfig) {
	if err := k.validatePutKmsEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKmsEncryptionConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ResetKmsEncryptionConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetKmsEncryptionConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ResetNoEncryptionConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetNoEncryptionConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamS3DestinationConfigurationEncryptionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

