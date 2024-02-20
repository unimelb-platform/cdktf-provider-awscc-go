package s3outpostsbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3outpostsbucket/internal"
)

type S3OutpostsBucketLifecycleConfigurationRulesOutputReference interface {
	cdktf.ComplexObject
	AbortIncompleteMultipartUpload() S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	AbortIncompleteMultipartUploadInput() interface{}
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
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	ExpirationInDays() *float64
	SetExpirationInDays(val *float64)
	ExpirationInDaysInput() *float64
	Filter() S3OutpostsBucketLifecycleConfigurationRulesFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *S3OutpostsBucketLifecycleConfigurationRules
	SetInternalValue(val *S3OutpostsBucketLifecycleConfigurationRules)
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	PutAbortIncompleteMultipartUpload(value *S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload)
	PutFilter(value *S3OutpostsBucketLifecycleConfigurationRulesFilter)
	ResetAbortIncompleteMultipartUpload()
	ResetExpirationDate()
	ResetExpirationInDays()
	ResetFilter()
	ResetId()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3OutpostsBucketLifecycleConfigurationRulesOutputReference
type jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUpload() S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference {
	var returns S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) Filter() S3OutpostsBucketLifecycleConfigurationRulesFilterOutputReference {
	var returns S3OutpostsBucketLifecycleConfigurationRulesFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) InternalValue() *S3OutpostsBucketLifecycleConfigurationRules {
	var returns *S3OutpostsBucketLifecycleConfigurationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3OutpostsBucketLifecycleConfigurationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3OutpostsBucketLifecycleConfigurationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewS3OutpostsBucketLifecycleConfigurationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference{}

	_jsii_.Create(
		"awscc.s3OutpostsBucket.S3OutpostsBucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3OutpostsBucketLifecycleConfigurationRulesOutputReference_Override(s S3OutpostsBucketLifecycleConfigurationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3OutpostsBucket.S3OutpostsBucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetExpirationDate(val *string) {
	if err := j.validateSetExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetExpirationInDays(val *float64) {
	if err := j.validateSetExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetInternalValue(val *S3OutpostsBucketLifecycleConfigurationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) PutAbortIncompleteMultipartUpload(value *S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload) {
	if err := s.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) PutFilter(value *S3OutpostsBucketLifecycleConfigurationRulesFilter) {
	if err := s.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetExpirationInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3OutpostsBucketLifecycleConfigurationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

