package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference interface {
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
	InternalValue() *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags
	SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags)
	RequireCommonName() interface{}
	SetRequireCommonName(val interface{})
	RequireCommonNameInput() interface{}
	RequireDirectoryPath() interface{}
	SetRequireDirectoryPath(val interface{})
	RequireDirectoryPathInput() interface{}
	RequireDnsAsCn() interface{}
	SetRequireDnsAsCn(val interface{})
	RequireDnsAsCnInput() interface{}
	RequireEmail() interface{}
	SetRequireEmail(val interface{})
	RequireEmailInput() interface{}
	SanRequireDirectoryGuid() interface{}
	SetSanRequireDirectoryGuid(val interface{})
	SanRequireDirectoryGuidInput() interface{}
	SanRequireDns() interface{}
	SetSanRequireDns(val interface{})
	SanRequireDnsInput() interface{}
	SanRequireDomainDns() interface{}
	SetSanRequireDomainDns(val interface{})
	SanRequireDomainDnsInput() interface{}
	SanRequireEmail() interface{}
	SetSanRequireEmail(val interface{})
	SanRequireEmailInput() interface{}
	SanRequireSpn() interface{}
	SetSanRequireSpn(val interface{})
	SanRequireSpnInput() interface{}
	SanRequireUpn() interface{}
	SetSanRequireUpn(val interface{})
	SanRequireUpnInput() interface{}
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
	ResetRequireCommonName()
	ResetRequireDirectoryPath()
	ResetRequireDnsAsCn()
	ResetRequireEmail()
	ResetSanRequireDirectoryGuid()
	ResetSanRequireDns()
	ResetSanRequireDomainDns()
	ResetSanRequireEmail()
	ResetSanRequireSpn()
	ResetSanRequireUpn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) InternalValue() *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireCommonName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireCommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireDirectoryPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireDirectoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireDirectoryPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireDirectoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireDnsAsCn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireDnsAsCn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireDnsAsCnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireDnsAsCnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireEmail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) RequireEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDirectoryGuid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDirectoryGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDirectoryGuidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDirectoryGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDomainDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDomainDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireDomainDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireDomainDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireEmail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireSpn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireSpn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireSpnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireSpnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireUpn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireUpn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) SanRequireUpnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sanRequireUpnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetRequireCommonName(val interface{}) {
	if err := j.validateSetRequireCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCommonName",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetRequireDirectoryPath(val interface{}) {
	if err := j.validateSetRequireDirectoryPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireDirectoryPath",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetRequireDnsAsCn(val interface{}) {
	if err := j.validateSetRequireDnsAsCnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireDnsAsCn",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetRequireEmail(val interface{}) {
	if err := j.validateSetRequireEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEmail",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireDirectoryGuid(val interface{}) {
	if err := j.validateSetSanRequireDirectoryGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireDirectoryGuid",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireDns(val interface{}) {
	if err := j.validateSetSanRequireDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireDns",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireDomainDns(val interface{}) {
	if err := j.validateSetSanRequireDomainDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireDomainDns",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireEmail(val interface{}) {
	if err := j.validateSetSanRequireEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireEmail",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireSpn(val interface{}) {
	if err := j.validateSetSanRequireSpnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireSpn",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetSanRequireUpn(val interface{}) {
	if err := j.validateSetSanRequireUpnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sanRequireUpn",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetRequireCommonName() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireCommonName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetRequireDirectoryPath() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireDirectoryPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetRequireDnsAsCn() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireDnsAsCn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetRequireEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireDirectoryGuid() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireDirectoryGuid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireDns() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireDns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireDomainDns() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireDomainDns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireSpn() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireSpn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ResetSanRequireUpn() {
	_jsii_.InvokeVoid(
		p,
		"resetSanRequireUpn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

