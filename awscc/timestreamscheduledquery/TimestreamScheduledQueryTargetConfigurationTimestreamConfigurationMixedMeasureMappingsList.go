package timestreamscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/timestreamscheduledquery/internal"
)

type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList
type jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList {
	_init_.Initialize()

	if err := validateNewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList{}

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewTimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList_Override(t TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.timestreamScheduledQuery.TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) Get(index *float64) TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
