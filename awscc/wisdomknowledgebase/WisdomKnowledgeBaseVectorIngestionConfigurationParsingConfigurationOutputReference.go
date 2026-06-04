package wisdomknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/wisdomknowledgebase/internal"
)

type WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference interface {
	cdktf.ComplexObject
	BedrockFoundationModelConfiguration() WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfigurationOutputReference
	BedrockFoundationModelConfigurationInput() interface{}
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
	ParsingStrategy() *string
	SetParsingStrategy(val *string)
	ParsingStrategyInput() *string
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
	PutBedrockFoundationModelConfiguration(value *WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration)
	ResetBedrockFoundationModelConfiguration()
	ResetParsingStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference
type jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) BedrockFoundationModelConfiguration() WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfigurationOutputReference {
	var returns WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfigurationOutputReference
	_jsii_.Get(
		j,
		"bedrockFoundationModelConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) BedrockFoundationModelConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bedrockFoundationModelConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ParsingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parsingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ParsingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parsingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.wisdomKnowledgeBase.WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference_Override(w WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.wisdomKnowledgeBase.WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetParsingStrategy(val *string) {
	if err := j.validateSetParsingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parsingStrategy",
		val,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) PutBedrockFoundationModelConfiguration(value *WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration) {
	if err := w.validatePutBedrockFoundationModelConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBedrockFoundationModelConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ResetBedrockFoundationModelConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetBedrockFoundationModelConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ResetParsingStrategy() {
	_jsii_.InvokeVoid(
		w,
		"resetParsingStrategy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

