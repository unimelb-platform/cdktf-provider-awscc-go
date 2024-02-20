package greengrassv2componentversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/greengrassv2componentversion/internal"
)

type Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap interface {
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
	Get(key *string) Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesOutputReference
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

// The jsii proxy struct for Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap
type jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGreengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap {
	_init_.Initialize()

	if err := validateNewGreengrassv2ComponentVersionLambdaFunctionComponentDependenciesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap{}

	_jsii_.Create(
		"awscc.greengrassv2ComponentVersion.Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGreengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap_Override(g Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.greengrassv2ComponentVersion.Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) Get(key *string) Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesOutputReference {
	if err := g.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_Greengrassv2ComponentVersionLambdaFunctionComponentDependenciesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

