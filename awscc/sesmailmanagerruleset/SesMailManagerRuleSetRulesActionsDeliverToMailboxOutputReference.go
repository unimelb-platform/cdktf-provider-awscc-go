package sesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sesmailmanagerruleset/internal"
)

type SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference interface {
	cdktf.ComplexObject
	ActionFailurePolicy() *string
	SetActionFailurePolicy(val *string)
	ActionFailurePolicyInput() *string
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
	MailboxArn() *string
	SetMailboxArn(val *string)
	MailboxArnInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetActionFailurePolicy()
	ResetMailboxArn()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference
type jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ActionFailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ActionFailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) MailboxArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) MailboxArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference{}

	_jsii_.Create(
		"awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference_Override(s SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetActionFailurePolicy(val *string) {
	if err := j.validateSetActionFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionFailurePolicy",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetMailboxArn(val *string) {
	if err := j.validateSetMailboxArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailboxArn",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ResetActionFailurePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetActionFailurePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ResetMailboxArn() {
	_jsii_.InvokeVoid(
		s,
		"resetMailboxArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

