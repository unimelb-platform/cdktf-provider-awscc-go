package elasticloadbalancingv2listener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2listener/internal"
)

type Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference interface {
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
	TargetGroups() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupsList
	TargetGroupsInput() interface{}
	TargetGroupStickinessConfig() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfigOutputReference
	TargetGroupStickinessConfigInput() interface{}
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
	PutTargetGroups(value interface{})
	PutTargetGroupStickinessConfig(value *Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfig)
	ResetTargetGroups()
	ResetTargetGroupStickinessConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference
type jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TargetGroups() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupsList {
	var returns Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupsList
	_jsii_.Get(
		j,
		"targetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TargetGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TargetGroupStickinessConfig() Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfigOutputReference {
	var returns Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfigOutputReference
	_jsii_.Get(
		j,
		"targetGroupStickinessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TargetGroupStickinessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupStickinessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference_Override(e Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2Listener.Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) PutTargetGroups(value interface{}) {
	if err := e.validatePutTargetGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetGroups",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) PutTargetGroupStickinessConfig(value *Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfig) {
	if err := e.validatePutTargetGroupStickinessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetGroupStickinessConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ResetTargetGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ResetTargetGroupStickinessConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetGroupStickinessConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerDefaultActionsForwardConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

