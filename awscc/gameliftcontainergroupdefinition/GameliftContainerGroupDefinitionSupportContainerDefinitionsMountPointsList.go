package gameliftcontainergroupdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/gameliftcontainergroupdefinition/internal"
)

type GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList interface {
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
	Get(index *float64) GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList
type jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList {
	_init_.Initialize()

	if err := validateNewGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList{}

	_jsii_.Create(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList_Override(g GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) Get(index *float64) GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

