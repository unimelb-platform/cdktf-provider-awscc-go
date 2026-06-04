package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/batchjobdefinition/internal"
)

type BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList interface {
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
	Get(index *float64) BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList
type jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList{}

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList_Override(b BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) Get(index *float64) BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

