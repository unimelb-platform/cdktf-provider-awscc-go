package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pipespipe/internal"
)

type PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList
type jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList{}

	_jsii_.Create(
		"awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList_Override(p PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := p.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		p,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) Get(index *float64) PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsMultiMeasureAttributeMappingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

