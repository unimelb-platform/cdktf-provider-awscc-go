package networkfirewallrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/networkfirewallrulegroup/internal"
)

type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference interface {
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
	InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions
	SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions)
	Keyword() *string
	SetKeyword(val *string)
	KeywordInput() *string
	Settings() *[]*string
	SetSettings(val *[]*string)
	SettingsInput() *[]*string
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
	ResetSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference
type jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions {
	var returns *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) Keyword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) KeywordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) Settings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) SettingsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference{}

	_jsii_.Create(
		"awscc.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference_Override(n NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetKeyword(val *string) {
	if err := j.validateSetKeywordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyword",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetSettings(val *[]*string) {
	if err := j.validateSetSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		n,
		"resetSettings",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

