package sesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sesmailmanagerruleset/internal"
)

type SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference interface {
	cdktf.ComplexObject
	Analysis() SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysisOutputReference
	AnalysisInput() interface{}
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
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
	IsInAddressList() SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStructOutputReference
	IsInAddressListInput() interface{}
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
	PutAnalysis(value *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysis)
	PutIsInAddressList(value *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStruct)
	ResetAnalysis()
	ResetAttribute()
	ResetIsInAddressList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference
type jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) Analysis() SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysisOutputReference {
	var returns SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysisOutputReference
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) AnalysisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) IsInAddressList() SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStructOutputReference {
	var returns SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStructOutputReference
	_jsii_.Get(
		j,
		"isInAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) IsInAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference{}

	_jsii_.Create(
		"awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference_Override(s SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sesMailManagerRuleSet.SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) PutAnalysis(value *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysis) {
	if err := s.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnalysis",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) PutIsInAddressList(value *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStruct) {
	if err := s.validatePutIsInAddressListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIsInAddressList",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		s,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ResetIsInAddressList() {
	_jsii_.InvokeVoid(
		s,
		"resetIsInAddressList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

