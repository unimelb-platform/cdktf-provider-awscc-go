package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference interface {
	cdktf.ComplexObject
	ApplicationPolicies() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPoliciesOutputReference
	ApplicationPoliciesInput() interface{}
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
	InternalValue() *PcaconnectoradTemplateDefinitionTemplateV3Extensions
	SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV3Extensions)
	KeyUsage() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsageOutputReference
	KeyUsageInput() *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage
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
	PutApplicationPolicies(value *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPolicies)
	PutKeyUsage(value *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage)
	ResetApplicationPolicies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ApplicationPolicies() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPoliciesOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPoliciesOutputReference
	_jsii_.Get(
		j,
		"applicationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ApplicationPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) InternalValue() *PcaconnectoradTemplateDefinitionTemplateV3Extensions {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3Extensions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) KeyUsage() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsageOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) KeyUsageInput() *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference)SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV3Extensions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) PutApplicationPolicies(value *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPolicies) {
	if err := p.validatePutApplicationPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApplicationPolicies",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) PutKeyUsage(value *PcaconnectoradTemplateDefinitionTemplateV3ExtensionsKeyUsage) {
	if err := p.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ResetApplicationPolicies() {
	_jsii_.InvokeVoid(
		p,
		"resetApplicationPolicies",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

