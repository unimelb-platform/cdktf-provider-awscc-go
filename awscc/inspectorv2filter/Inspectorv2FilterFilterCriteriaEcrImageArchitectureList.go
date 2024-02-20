package inspectorv2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/inspectorv2filter/internal"
)

type Inspectorv2FilterFilterCriteriaEcrImageArchitectureList interface {
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
	Get(index *float64) Inspectorv2FilterFilterCriteriaEcrImageArchitectureOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2FilterFilterCriteriaEcrImageArchitectureList
type jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewInspectorv2FilterFilterCriteriaEcrImageArchitectureList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Inspectorv2FilterFilterCriteriaEcrImageArchitectureList {
	_init_.Initialize()

	if err := validateNewInspectorv2FilterFilterCriteriaEcrImageArchitectureListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList{}

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaEcrImageArchitectureList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewInspectorv2FilterFilterCriteriaEcrImageArchitectureList_Override(i Inspectorv2FilterFilterCriteriaEcrImageArchitectureList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaEcrImageArchitectureList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) Get(index *float64) Inspectorv2FilterFilterCriteriaEcrImageArchitectureOutputReference {
	if err := i.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Inspectorv2FilterFilterCriteriaEcrImageArchitectureOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaEcrImageArchitectureList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

