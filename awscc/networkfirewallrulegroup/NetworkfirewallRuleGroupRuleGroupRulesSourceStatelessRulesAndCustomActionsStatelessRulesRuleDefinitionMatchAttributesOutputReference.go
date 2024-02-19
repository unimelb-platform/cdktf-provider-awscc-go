package networkfirewallrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkfirewallrulegroup/internal"
)

type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference interface {
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
	DestinationPorts() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationPortsList
	DestinationPortsInput() interface{}
	Destinations() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationsList
	DestinationsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributes
	SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributes)
	Protocols() *[]*float64
	SetProtocols(val *[]*float64)
	ProtocolsInput() *[]*float64
	SourcePorts() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcePortsList
	SourcePortsInput() interface{}
	Sources() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcesList
	SourcesInput() interface{}
	TcpFlags() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesTcpFlagsList
	TcpFlagsInput() interface{}
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
	PutDestinationPorts(value interface{})
	PutDestinations(value interface{})
	PutSourcePorts(value interface{})
	PutSources(value interface{})
	PutTcpFlags(value interface{})
	ResetDestinationPorts()
	ResetDestinations()
	ResetProtocols()
	ResetSourcePorts()
	ResetSources()
	ResetTcpFlags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference
type jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) DestinationPorts() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationPortsList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationPortsList
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) DestinationPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) Destinations() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationsList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributes {
	var returns *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) Protocols() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ProtocolsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) SourcePorts() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcePortsList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcePortsList
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) SourcePortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) Sources() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcesList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesSourcesList
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) SourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) TcpFlags() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesTcpFlagsList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesTcpFlagsList
	_jsii_.Get(
		j,
		"tcpFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) TcpFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference{}

	_jsii_.Create(
		"awscc.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference_Override(n NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetProtocols(val *[]*float64) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) PutDestinationPorts(value interface{}) {
	if err := n.validatePutDestinationPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestinationPorts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) PutDestinations(value interface{}) {
	if err := n.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestinations",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) PutSourcePorts(value interface{}) {
	if err := n.validatePutSourcePortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSourcePorts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) PutSources(value interface{}) {
	if err := n.validatePutSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSources",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) PutTcpFlags(value interface{}) {
	if err := n.validatePutTcpFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTcpFlags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinations",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocols",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		n,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetSources() {
	_jsii_.InvokeVoid(
		n,
		"resetSources",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ResetTcpFlags() {
	_jsii_.InvokeVoid(
		n,
		"resetTcpFlags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRulesRuleDefinitionMatchAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

