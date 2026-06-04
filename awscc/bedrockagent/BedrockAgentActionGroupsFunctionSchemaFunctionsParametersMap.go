package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/bedrockagent/internal"
)

type BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap interface {
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
	Get(key *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference
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

// The jsii proxy struct for BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap
type jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap {
	_init_.Initialize()

	if err := validateNewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap{}

	_jsii_.Create(
		"awscc.bedrockAgent.BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap_Override(b BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.bedrockAgent.BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Get(key *string) BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference {
	if err := b.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns BedrockAgentActionGroupsFunctionSchemaFunctionsParametersOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockAgentActionGroupsFunctionSchemaFunctionsParametersMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

