package ssmresourcedatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ssmresourcedatasync/internal"
)

type SsmResourceDataSyncSyncSourceOutputReference interface {
	cdktf.ComplexObject
	AwsOrganizationsSource() SsmResourceDataSyncSyncSourceAwsOrganizationsSourceOutputReference
	AwsOrganizationsSourceInput() interface{}
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
	IncludeFutureRegions() interface{}
	SetIncludeFutureRegions(val interface{})
	IncludeFutureRegionsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceRegions() *[]*string
	SetSourceRegions(val *[]*string)
	SourceRegionsInput() *[]*string
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
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
	PutAwsOrganizationsSource(value *SsmResourceDataSyncSyncSourceAwsOrganizationsSource)
	ResetAwsOrganizationsSource()
	ResetIncludeFutureRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmResourceDataSyncSyncSourceOutputReference
type jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) AwsOrganizationsSource() SsmResourceDataSyncSyncSourceAwsOrganizationsSourceOutputReference {
	var returns SsmResourceDataSyncSyncSourceAwsOrganizationsSourceOutputReference
	_jsii_.Get(
		j,
		"awsOrganizationsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) AwsOrganizationsSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsOrganizationsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) IncludeFutureRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeFutureRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) IncludeFutureRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeFutureRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) SourceRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) SourceRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSsmResourceDataSyncSyncSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SsmResourceDataSyncSyncSourceOutputReference {
	_init_.Initialize()

	if err := validateNewSsmResourceDataSyncSyncSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference{}

	_jsii_.Create(
		"awscc.ssmResourceDataSync.SsmResourceDataSyncSyncSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSsmResourceDataSyncSyncSourceOutputReference_Override(s SsmResourceDataSyncSyncSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ssmResourceDataSync.SsmResourceDataSyncSyncSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetIncludeFutureRegions(val interface{}) {
	if err := j.validateSetIncludeFutureRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFutureRegions",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetSourceRegions(val *[]*string) {
	if err := j.validateSetSourceRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegions",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetSourceType(val *string) {
	if err := j.validateSetSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) PutAwsOrganizationsSource(value *SsmResourceDataSyncSyncSourceAwsOrganizationsSource) {
	if err := s.validatePutAwsOrganizationsSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsOrganizationsSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ResetAwsOrganizationsSource() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsOrganizationsSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ResetIncludeFutureRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeFutureRegions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmResourceDataSyncSyncSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

