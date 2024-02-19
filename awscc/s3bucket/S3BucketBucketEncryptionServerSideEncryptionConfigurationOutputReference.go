package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3bucket/internal"
)

type S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference interface {
	cdktf.ComplexObject
	BucketKeyEnabled() interface{}
	SetBucketKeyEnabled(val interface{})
	BucketKeyEnabledInput() interface{}
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
	InternalValue() *S3BucketBucketEncryptionServerSideEncryptionConfiguration
	SetInternalValue(val *S3BucketBucketEncryptionServerSideEncryptionConfiguration)
	ServerSideEncryptionByDefault() S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefaultOutputReference
	ServerSideEncryptionByDefaultInput() interface{}
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
	PutServerSideEncryptionByDefault(value *S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefault)
	ResetBucketKeyEnabled()
	ResetServerSideEncryptionByDefault()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference
type jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) InternalValue() *S3BucketBucketEncryptionServerSideEncryptionConfiguration {
	var returns *S3BucketBucketEncryptionServerSideEncryptionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ServerSideEncryptionByDefault() S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefaultOutputReference {
	var returns S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefaultOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryptionByDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ServerSideEncryptionByDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionByDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference_Override(s S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetInternalValue(val *S3BucketBucketEncryptionServerSideEncryptionConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) PutServerSideEncryptionByDefault(value *S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefault) {
	if err := s.validatePutServerSideEncryptionByDefaultParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServerSideEncryptionByDefault",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ResetServerSideEncryptionByDefault() {
	_jsii_.InvokeVoid(
		s,
		"resetServerSideEncryptionByDefault",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketBucketEncryptionServerSideEncryptionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

