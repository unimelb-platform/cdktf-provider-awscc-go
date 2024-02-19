package ec2networkinsightsaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2networkinsightsaccessscope/internal"
)

type Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference interface {
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
	DestinationAddresses() *[]*string
	SetDestinationAddresses(val *[]*string)
	DestinationAddressesInput() *[]*string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	DestinationPrefixLists() *[]*string
	SetDestinationPrefixLists(val *[]*string)
	DestinationPrefixListsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourcePorts() *[]*string
	SetSourcePorts(val *[]*string)
	SourcePortsInput() *[]*string
	SourcePrefixLists() *[]*string
	SetSourcePrefixLists(val *[]*string)
	SourcePrefixListsInput() *[]*string
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
	ResetDestinationAddresses()
	ResetDestinationPorts()
	ResetDestinationPrefixLists()
	ResetProtocols()
	ResetSourceAddresses()
	ResetSourcePorts()
	ResetSourcePrefixLists()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference
type jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationPrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) DestinationPrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourcePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourcePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourcePrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) SourcePrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference{}

	_jsii_.Create(
		"awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference_Override(e Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2NetworkInsightsAccessScope.Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetDestinationPrefixLists(val *[]*string) {
	if err := j.validateSetDestinationPrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixLists",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetSourcePorts(val *[]*string) {
	if err := j.validateSetSourcePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePorts",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetSourcePrefixLists(val *[]*string) {
	if err := j.validateSetSourcePrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePrefixLists",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetDestinationPrefixLists() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPrefixLists",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocols",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ResetSourcePrefixLists() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePrefixLists",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAccessScopeMatchPathsSourcePacketHeaderStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

