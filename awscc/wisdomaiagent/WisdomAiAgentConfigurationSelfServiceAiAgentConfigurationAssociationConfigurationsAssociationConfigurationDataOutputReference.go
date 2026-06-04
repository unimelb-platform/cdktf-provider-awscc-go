package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/wisdomaiagent/internal"
)

type WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference interface {
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
	KnowledgeBaseAssociationConfigurationData() WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataOutputReference
	KnowledgeBaseAssociationConfigurationDataInput() interface{}
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
	PutKnowledgeBaseAssociationConfigurationData(value *WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationData)
	ResetKnowledgeBaseAssociationConfigurationData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference
type jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) KnowledgeBaseAssociationConfigurationData() WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataOutputReference {
	var returns WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataOutputReference
	_jsii_.Get(
		j,
		"knowledgeBaseAssociationConfigurationData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) KnowledgeBaseAssociationConfigurationDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseAssociationConfigurationDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference{}

	_jsii_.Create(
		"awscc.wisdomAiAgent.WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference_Override(w WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.wisdomAiAgent.WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) PutKnowledgeBaseAssociationConfigurationData(value *WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationData) {
	if err := w.validatePutKnowledgeBaseAssociationConfigurationDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putKnowledgeBaseAssociationConfigurationData",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) ResetKnowledgeBaseAssociationConfigurationData() {
	_jsii_.InvokeVoid(
		w,
		"resetKnowledgeBaseAssociationConfigurationData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

