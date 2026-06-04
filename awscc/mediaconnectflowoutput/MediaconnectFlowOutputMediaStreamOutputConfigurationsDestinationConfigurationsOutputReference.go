package mediaconnectflowoutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediaconnectflowoutput/internal"
)

type MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference interface {
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
	DestinationIp() *string
	SetDestinationIp(val *string)
	DestinationIpInput() *string
	DestinationPort() *float64
	SetDestinationPort(val *float64)
	DestinationPortInput() *float64
	// Experimental.
	Fqn() *string
	Interface() MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterfaceOutputReference
	InterfaceInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutInterface(value *MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterface)
	ResetDestinationIp()
	ResetDestinationPort()
	ResetInterface()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference
type jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) DestinationIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) DestinationIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) DestinationPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) DestinationPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) Interface() MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterfaceOutputReference {
	var returns MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterfaceOutputReference
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) InterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference{}

	_jsii_.Create(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference_Override(m MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediaconnectFlowOutput.MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetDestinationIp(val *string) {
	if err := j.validateSetDestinationIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIp",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetDestinationPort(val *float64) {
	if err := j.validateSetDestinationPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) PutInterface(value *MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterface) {
	if err := m.validatePutInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInterface",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ResetDestinationIp() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		m,
		"resetInterface",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

