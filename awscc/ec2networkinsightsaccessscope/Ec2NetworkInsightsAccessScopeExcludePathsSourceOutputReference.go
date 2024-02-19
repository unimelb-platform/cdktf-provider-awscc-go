package ec2networkinsightsaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2networkinsightsaccessscope/internal"
)

type Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference interface {
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
	PacketHeaderStatement() Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference
	PacketHeaderStatementInput() interface{}
	ResourceStatement() Ec2NetworkInsightsAccessScopeExcludePathsSourceResourceStatementOutputReference
	ResourceStatementInput() interface{}
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
	PutPacketHeaderStatement(value *Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatement)
	PutResourceStatement(value *Ec2NetworkInsightsAccessScopeExcludePathsSourceResourceStatement)
	ResetPacketHeaderStatement()
	ResetResourceStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference
type jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) PacketHeaderStatement() Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference {
	var returns Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatementOutputReference
	_jsii_.Get(
		j,
		"packetHeaderStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) PacketHeaderStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packetHeaderStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ResourceStatement() Ec2NetworkInsightsAccessScopeExcludePathsSourceResourceStatementOutputReference {
	var returns Ec2NetworkInsightsAccessScopeExcludePathsSourceResourceStatementOutputReference
	_jsii_.Get(
		j,
		"resourceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ResourceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAccessScopeExcludePathsSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAccessScopeExcludePathsSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference{}

	_jsii_.Create(
		"awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAccessScopeExcludePathsSourceOutputReference_Override(e Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) PutPacketHeaderStatement(value *Ec2NetworkInsightsAccessScopeExcludePathsSourcePacketHeaderStatement) {
	if err := e.validatePutPacketHeaderStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPacketHeaderStatement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) PutResourceStatement(value *Ec2NetworkInsightsAccessScopeExcludePathsSourceResourceStatement) {
	if err := e.validatePutResourceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceStatement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ResetPacketHeaderStatement() {
	_jsii_.InvokeVoid(
		e,
		"resetPacketHeaderStatement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ResetResourceStatement() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceStatement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeExcludePathsSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

