package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

type SagemakerModelPackageModelMetricsOutputReference interface {
	cdktf.ComplexObject
	Bias() SagemakerModelPackageModelMetricsBiasOutputReference
	BiasInput() interface{}
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
	Explainability() SagemakerModelPackageModelMetricsExplainabilityOutputReference
	ExplainabilityInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelDataQuality() SagemakerModelPackageModelMetricsModelDataQualityOutputReference
	ModelDataQualityInput() interface{}
	ModelQuality() SagemakerModelPackageModelMetricsModelQualityOutputReference
	ModelQualityInput() interface{}
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
	PutBias(value *SagemakerModelPackageModelMetricsBias)
	PutExplainability(value *SagemakerModelPackageModelMetricsExplainability)
	PutModelDataQuality(value *SagemakerModelPackageModelMetricsModelDataQuality)
	PutModelQuality(value *SagemakerModelPackageModelMetricsModelQuality)
	ResetBias()
	ResetExplainability()
	ResetModelDataQuality()
	ResetModelQuality()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageModelMetricsOutputReference
type jsiiProxy_SagemakerModelPackageModelMetricsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) Bias() SagemakerModelPackageModelMetricsBiasOutputReference {
	var returns SagemakerModelPackageModelMetricsBiasOutputReference
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) BiasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) Explainability() SagemakerModelPackageModelMetricsExplainabilityOutputReference {
	var returns SagemakerModelPackageModelMetricsExplainabilityOutputReference
	_jsii_.Get(
		j,
		"explainability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ExplainabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explainabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ModelDataQuality() SagemakerModelPackageModelMetricsModelDataQualityOutputReference {
	var returns SagemakerModelPackageModelMetricsModelDataQualityOutputReference
	_jsii_.Get(
		j,
		"modelDataQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ModelDataQualityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelDataQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ModelQuality() SagemakerModelPackageModelMetricsModelQualityOutputReference {
	var returns SagemakerModelPackageModelMetricsModelQualityOutputReference
	_jsii_.Get(
		j,
		"modelQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ModelQualityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageModelMetricsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageModelMetricsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageModelMetricsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageModelMetricsOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageModelMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageModelMetricsOutputReference_Override(s SagemakerModelPackageModelMetricsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageModelMetricsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) PutBias(value *SagemakerModelPackageModelMetricsBias) {
	if err := s.validatePutBiasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBias",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) PutExplainability(value *SagemakerModelPackageModelMetricsExplainability) {
	if err := s.validatePutExplainabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExplainability",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) PutModelDataQuality(value *SagemakerModelPackageModelMetricsModelDataQuality) {
	if err := s.validatePutModelDataQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelDataQuality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) PutModelQuality(value *SagemakerModelPackageModelMetricsModelQuality) {
	if err := s.validatePutModelQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQuality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		s,
		"resetBias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ResetExplainability() {
	_jsii_.InvokeVoid(
		s,
		"resetExplainability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ResetModelDataQuality() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataQuality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ResetModelQuality() {
	_jsii_.InvokeVoid(
		s,
		"resetModelQuality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

