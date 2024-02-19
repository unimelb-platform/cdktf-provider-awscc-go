package dataawsccnetworkfirewallfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccnetworkfirewallfirewallpolicy/internal"
)

type DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap interface {
	cdktf.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference
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

// The jsii proxy struct for DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap
type jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap {
	_init_.Initialize()

	if err := validateNewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap{}

	_jsii_.Create(
		"awscc.dataAwsccNetworkfirewallFirewallPolicy.DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap_Override(d DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccNetworkfirewallFirewallPolicy.DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Get(key *string) DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

