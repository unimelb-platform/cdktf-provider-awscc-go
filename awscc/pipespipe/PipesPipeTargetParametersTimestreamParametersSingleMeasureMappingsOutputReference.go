package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pipespipe/internal"
)

type PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference interface {
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
	MeasureName() *string
	SetMeasureName(val *string)
	MeasureNameInput() *string
	MeasureValue() *string
	SetMeasureValue(val *string)
	MeasureValueInput() *string
	MeasureValueType() *string
	SetMeasureValueType(val *string)
	MeasureValueTypeInput() *string
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
	ResetMeasureName()
	ResetMeasureValue()
	ResetMeasureValueType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference
type jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) MeasureValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference{}

	_jsii_.Create(
		"awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference_Override(p PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetMeasureName(val *string) {
	if err := j.validateSetMeasureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureName",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetMeasureValue(val *string) {
	if err := j.validateSetMeasureValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetMeasureValueType(val *string) {
	if err := j.validateSetMeasureValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureValueType",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ResetMeasureName() {
	_jsii_.InvokeVoid(
		p,
		"resetMeasureName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ResetMeasureValue() {
	_jsii_.InvokeVoid(
		p,
		"resetMeasureValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ResetMeasureValueType() {
	_jsii_.InvokeVoid(
		p,
		"resetMeasureValueType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

