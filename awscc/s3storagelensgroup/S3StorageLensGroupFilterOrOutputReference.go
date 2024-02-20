package s3storagelensgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3storagelensgroup/internal"
)

type S3StorageLensGroupFilterOrOutputReference interface {
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
	MatchAnyPrefix() *[]*string
	SetMatchAnyPrefix(val *[]*string)
	MatchAnyPrefixInput() *[]*string
	MatchAnySuffix() *[]*string
	SetMatchAnySuffix(val *[]*string)
	MatchAnySuffixInput() *[]*string
	MatchAnyTag() S3StorageLensGroupFilterOrMatchAnyTagList
	MatchAnyTagInput() interface{}
	MatchObjectAge() S3StorageLensGroupFilterOrMatchObjectAgeOutputReference
	MatchObjectAgeInput() interface{}
	MatchObjectSize() S3StorageLensGroupFilterOrMatchObjectSizeOutputReference
	MatchObjectSizeInput() interface{}
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
	PutMatchAnyTag(value interface{})
	PutMatchObjectAge(value *S3StorageLensGroupFilterOrMatchObjectAge)
	PutMatchObjectSize(value *S3StorageLensGroupFilterOrMatchObjectSize)
	ResetMatchAnyPrefix()
	ResetMatchAnySuffix()
	ResetMatchAnyTag()
	ResetMatchObjectAge()
	ResetMatchObjectSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensGroupFilterOrOutputReference
type jsiiProxy_S3StorageLensGroupFilterOrOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnyPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnyPrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnySuffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnySuffixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnyTag() S3StorageLensGroupFilterOrMatchAnyTagList {
	var returns S3StorageLensGroupFilterOrMatchAnyTagList
	_jsii_.Get(
		j,
		"matchAnyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchAnyTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchAnyTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchObjectAge() S3StorageLensGroupFilterOrMatchObjectAgeOutputReference {
	var returns S3StorageLensGroupFilterOrMatchObjectAgeOutputReference
	_jsii_.Get(
		j,
		"matchObjectAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchObjectAgeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchObjectSize() S3StorageLensGroupFilterOrMatchObjectSizeOutputReference {
	var returns S3StorageLensGroupFilterOrMatchObjectSizeOutputReference
	_jsii_.Get(
		j,
		"matchObjectSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) MatchObjectSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensGroupFilterOrOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3StorageLensGroupFilterOrOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensGroupFilterOrOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensGroupFilterOrOutputReference{}

	_jsii_.Create(
		"awscc.s3StorageLensGroup.S3StorageLensGroupFilterOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensGroupFilterOrOutputReference_Override(s S3StorageLensGroupFilterOrOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3StorageLensGroup.S3StorageLensGroupFilterOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetMatchAnyPrefix(val *[]*string) {
	if err := j.validateSetMatchAnyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnyPrefix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetMatchAnySuffix(val *[]*string) {
	if err := j.validateSetMatchAnySuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnySuffix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOrOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) PutMatchAnyTag(value interface{}) {
	if err := s.validatePutMatchAnyTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchAnyTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) PutMatchObjectAge(value *S3StorageLensGroupFilterOrMatchObjectAge) {
	if err := s.validatePutMatchObjectAgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectAge",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) PutMatchObjectSize(value *S3StorageLensGroupFilterOrMatchObjectSize) {
	if err := s.validatePutMatchObjectSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectSize",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ResetMatchAnyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ResetMatchAnySuffix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnySuffix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ResetMatchAnyTag() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ResetMatchObjectAge() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectAge",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ResetMatchObjectSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOrOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

