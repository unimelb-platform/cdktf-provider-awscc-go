package s3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3storagelens/internal"
)

type S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference interface {
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
	InternalValue() *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics
	SetInternalValue(val *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics)
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	SelectionCriteria() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteriaOutputReference
	SelectionCriteriaInput() interface{}
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
	PutSelectionCriteria(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria)
	ResetIsEnabled()
	ResetSelectionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference
type jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) InternalValue() *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics {
	var returns *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) SelectionCriteria() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteriaOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteriaOutputReference
	_jsii_.Get(
		j,
		"selectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) SelectionCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference{}

	_jsii_.Create(
		"awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference_Override(s S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetInternalValue(val *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) PutSelectionCriteria(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria) {
	if err := s.validatePutSelectionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSelectionCriteria",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ResetSelectionCriteria() {
	_jsii_.InvokeVoid(
		s,
		"resetSelectionCriteria",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

