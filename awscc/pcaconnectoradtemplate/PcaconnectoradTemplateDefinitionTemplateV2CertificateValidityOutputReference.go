package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference interface {
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
	InternalValue() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity
	SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity)
	RenewalPeriod() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriodOutputReference
	RenewalPeriodInput() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidityPeriod() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriodOutputReference
	ValidityPeriodInput() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod
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
	PutRenewalPeriod(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod)
	PutValidityPeriod(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) InternalValue() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity {
	var returns *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) RenewalPeriod() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriodOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriodOutputReference
	_jsii_.Get(
		j,
		"renewalPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) RenewalPeriodInput() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod {
	var returns *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod
	_jsii_.Get(
		j,
		"renewalPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ValidityPeriod() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriodOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriodOutputReference
	_jsii_.Get(
		j,
		"validityPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ValidityPeriodInput() *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod {
	var returns *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod
	_jsii_.Get(
		j,
		"validityPeriodInput",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference)SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) PutRenewalPeriod(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod) {
	if err := p.validatePutRenewalPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRenewalPeriod",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) PutValidityPeriod(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod) {
	if err := p.validatePutValidityPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putValidityPeriod",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

