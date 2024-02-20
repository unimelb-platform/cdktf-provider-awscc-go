package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

type SagemakerModelPackageInferenceSpecificationOutputReference interface {
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
	Containers() SagemakerModelPackageInferenceSpecificationContainersList
	ContainersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SupportedContentTypes() *[]*string
	SetSupportedContentTypes(val *[]*string)
	SupportedContentTypesInput() *[]*string
	SupportedRealtimeInferenceInstanceTypes() *[]*string
	SetSupportedRealtimeInferenceInstanceTypes(val *[]*string)
	SupportedRealtimeInferenceInstanceTypesInput() *[]*string
	SupportedResponseMimeTypes() *[]*string
	SetSupportedResponseMimeTypes(val *[]*string)
	SupportedResponseMimeTypesInput() *[]*string
	SupportedTransformInstanceTypes() *[]*string
	SetSupportedTransformInstanceTypes(val *[]*string)
	SupportedTransformInstanceTypesInput() *[]*string
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
	PutContainers(value interface{})
	ResetSupportedRealtimeInferenceInstanceTypes()
	ResetSupportedTransformInstanceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageInferenceSpecificationOutputReference
type jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) Containers() SagemakerModelPackageInferenceSpecificationContainersList {
	var returns SagemakerModelPackageInferenceSpecificationContainersList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedContentTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedContentTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedRealtimeInferenceInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedRealtimeInferenceInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedResponseMimeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedResponseMimeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedTransformInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) SupportedTransformInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageInferenceSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageInferenceSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageInferenceSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageInferenceSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageInferenceSpecificationOutputReference_Override(s SagemakerModelPackageInferenceSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageInferenceSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetSupportedContentTypes(val *[]*string) {
	if err := j.validateSetSupportedContentTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedContentTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetSupportedRealtimeInferenceInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedRealtimeInferenceInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetSupportedResponseMimeTypes(val *[]*string) {
	if err := j.validateSetSupportedResponseMimeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedResponseMimeTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetSupportedTransformInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedTransformInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedTransformInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) PutContainers(value interface{}) {
	if err := s.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ResetSupportedRealtimeInferenceInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedRealtimeInferenceInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ResetSupportedTransformInstanceTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportedTransformInstanceTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageInferenceSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

