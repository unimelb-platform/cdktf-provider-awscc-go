package networkfirewallfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkfirewallfirewallpolicy/internal"
)

type NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap interface {
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
	Get(key *string) NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference
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

// The jsii proxy struct for NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap
type jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap {
	_init_.Initialize()

	if err := validateNewNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap{}

	_jsii_.Create(
		"awscc.networkfirewallFirewallPolicy.NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap_Override(n NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkfirewallFirewallPolicy.NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Get(key *string) NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference {
	if err := n.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

