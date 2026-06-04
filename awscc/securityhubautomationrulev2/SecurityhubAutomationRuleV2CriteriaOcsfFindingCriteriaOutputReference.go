package securityhubautomationrulev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/securityhubautomationrulev2/internal"
)

type SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference interface {
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
	CompositeFilters() SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersList
	CompositeFiltersInput() interface{}
	CompositeOperator() *string
	SetCompositeOperator(val *string)
	CompositeOperatorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutCompositeFilters(value interface{})
	ResetCompositeFilters()
	ResetCompositeOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference
type jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) CompositeFilters() SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersList {
	var returns SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaCompositeFiltersList
	_jsii_.Get(
		j,
		"compositeFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) CompositeFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) CompositeOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compositeOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) CompositeOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compositeOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference{}

	_jsii_.Create(
		"awscc.securityhubAutomationRuleV2.SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference_Override(s SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.securityhubAutomationRuleV2.SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetCompositeOperator(val *string) {
	if err := j.validateSetCompositeOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compositeOperator",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) PutCompositeFilters(value interface{}) {
	if err := s.validatePutCompositeFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCompositeFilters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ResetCompositeFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetCompositeFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ResetCompositeOperator() {
	_jsii_.InvokeVoid(
		s,
		"resetCompositeOperator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubAutomationRuleV2CriteriaOcsfFindingCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

