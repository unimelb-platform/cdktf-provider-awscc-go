package maciefindingsfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/maciefindingsfilter/internal"
)

type MacieFindingsFilterFindingCriteriaCriterionMap interface {
	cdktf.ComplexMap
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
	// Experimental.
	ComputeFqn() *string
	Get(key *string) MacieFindingsFilterFindingCriteriaCriterionOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MacieFindingsFilterFindingCriteriaCriterionMap
type jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacieFindingsFilterFindingCriteriaCriterionMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MacieFindingsFilterFindingCriteriaCriterionMap {
	_init_.Initialize()

	if err := validateNewMacieFindingsFilterFindingCriteriaCriterionMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap{}

	_jsii_.Create(
		"awscc.macieFindingsFilter.MacieFindingsFilterFindingCriteriaCriterionMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMacieFindingsFilterFindingCriteriaCriterionMap_Override(m MacieFindingsFilterFindingCriteriaCriterionMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.macieFindingsFilter.MacieFindingsFilterFindingCriteriaCriterionMap",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) Get(key *string) MacieFindingsFilterFindingCriteriaCriterionOutputReference {
	if err := m.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns MacieFindingsFilterFindingCriteriaCriterionOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacieFindingsFilterFindingCriteriaCriterionMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

