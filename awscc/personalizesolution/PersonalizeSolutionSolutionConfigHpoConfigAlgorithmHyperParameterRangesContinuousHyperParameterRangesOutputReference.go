package personalizesolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/personalizesolution/internal"
)

type PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference interface {
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
	MaxValue() *float64
	SetMaxValue(val *float64)
	MaxValueInput() *float64
	MinValue() *float64
	SetMinValue(val *float64)
	MinValueInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetMaxValue()
	ResetMinValue()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference
type jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MaxValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MaxValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MinValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) MinValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference {
	_init_.Initialize()

	if err := validateNewPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference{}

	_jsii_.Create(
		"awscc.personalizeSolution.PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference_Override(p PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.personalizeSolution.PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetMaxValue(val *float64) {
	if err := j.validateSetMaxValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetMinValue(val *float64) {
	if err := j.validateSetMinValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ResetMaxValue() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ResetMinValue() {
	_jsii_.InvokeVoid(
		p,
		"resetMinValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesContinuousHyperParameterRangesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

