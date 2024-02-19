package sagemakerinferencecomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakerinferencecomponent/internal"
)

type SagemakerInferenceComponentSpecificationStartupParametersOutputReference interface {
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
	ContainerStartupHealthCheckTimeoutInSeconds() *float64
	SetContainerStartupHealthCheckTimeoutInSeconds(val *float64)
	ContainerStartupHealthCheckTimeoutInSecondsInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelDataDownloadTimeoutInSeconds() *float64
	SetModelDataDownloadTimeoutInSeconds(val *float64)
	ModelDataDownloadTimeoutInSecondsInput() *float64
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
	ResetContainerStartupHealthCheckTimeoutInSeconds()
	ResetModelDataDownloadTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerInferenceComponentSpecificationStartupParametersOutputReference
type jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerInferenceComponentSpecificationStartupParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerInferenceComponentSpecificationStartupParametersOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerInferenceComponentSpecificationStartupParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationStartupParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerInferenceComponentSpecificationStartupParametersOutputReference_Override(s SagemakerInferenceComponentSpecificationStartupParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerInferenceComponent.SagemakerInferenceComponentSpecificationStartupParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentSpecificationStartupParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

