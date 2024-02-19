package route53resolverfirewallrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/route53resolverfirewallrulegroup/internal"
)

type Route53ResolverFirewallRuleGroupFirewallRulesOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	BlockOverrideDnsType() *string
	SetBlockOverrideDnsType(val *string)
	BlockOverrideDnsTypeInput() *string
	BlockOverrideDomain() *string
	SetBlockOverrideDomain(val *string)
	BlockOverrideDomainInput() *string
	BlockOverrideTtl() *float64
	SetBlockOverrideTtl(val *float64)
	BlockOverrideTtlInput() *float64
	BlockResponse() *string
	SetBlockResponse(val *string)
	BlockResponseInput() *string
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
	FirewallDomainListId() *string
	SetFirewallDomainListId(val *string)
	FirewallDomainListIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Qtype() *string
	SetQtype(val *string)
	QtypeInput() *string
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
	ResetBlockOverrideDnsType()
	ResetBlockOverrideDomain()
	ResetBlockOverrideTtl()
	ResetBlockResponse()
	ResetQtype()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53ResolverFirewallRuleGroupFirewallRulesOutputReference
type jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideDnsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockOverrideTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) BlockResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) FirewallDomainListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) Qtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) QtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoute53ResolverFirewallRuleGroupFirewallRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Route53ResolverFirewallRuleGroupFirewallRulesOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53ResolverFirewallRuleGroupFirewallRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference{}

	_jsii_.Create(
		"awscc.route53ResolverFirewallRuleGroup.Route53ResolverFirewallRuleGroupFirewallRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRoute53ResolverFirewallRuleGroupFirewallRulesOutputReference_Override(r Route53ResolverFirewallRuleGroupFirewallRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.route53ResolverFirewallRuleGroup.Route53ResolverFirewallRuleGroupFirewallRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetBlockOverrideDnsType(val *string) {
	if err := j.validateSetBlockOverrideDnsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDnsType",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetBlockOverrideDomain(val *string) {
	if err := j.validateSetBlockOverrideDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDomain",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetBlockOverrideTtl(val *float64) {
	if err := j.validateSetBlockOverrideTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideTtl",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetBlockResponse(val *string) {
	if err := j.validateSetBlockResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockResponse",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetFirewallDomainListId(val *string) {
	if err := j.validateSetFirewallDomainListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallDomainListId",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetQtype(val *string) {
	if err := j.validateSetQtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qtype",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ResetBlockOverrideDnsType() {
	_jsii_.InvokeVoid(
		r,
		"resetBlockOverrideDnsType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ResetBlockOverrideDomain() {
	_jsii_.InvokeVoid(
		r,
		"resetBlockOverrideDomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ResetBlockOverrideTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetBlockOverrideTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ResetBlockResponse() {
	_jsii_.InvokeVoid(
		r,
		"resetBlockResponse",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ResetQtype() {
	_jsii_.InvokeVoid(
		r,
		"resetQtype",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53ResolverFirewallRuleGroupFirewallRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

