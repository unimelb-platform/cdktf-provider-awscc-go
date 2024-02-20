package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference interface {
	cdktf.ComplexObject
	ClientVersion() *string
	SetClientVersion(val *string)
	ClientVersionInput() *string
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
	ExportableKey() interface{}
	SetExportableKey(val interface{})
	ExportableKeyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags
	SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags)
	RequireAlternateSignatureAlgorithm() interface{}
	SetRequireAlternateSignatureAlgorithm(val interface{})
	RequireAlternateSignatureAlgorithmInput() interface{}
	RequireSameKeyRenewal() interface{}
	SetRequireSameKeyRenewal(val interface{})
	RequireSameKeyRenewalInput() interface{}
	StrongKeyProtectionRequired() interface{}
	SetStrongKeyProtectionRequired(val interface{})
	StrongKeyProtectionRequiredInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseLegacyProvider() interface{}
	SetUseLegacyProvider(val interface{})
	UseLegacyProviderInput() interface{}
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
	ResetExportableKey()
	ResetRequireAlternateSignatureAlgorithm()
	ResetRequireSameKeyRenewal()
	ResetStrongKeyProtectionRequired()
	ResetUseLegacyProvider()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ClientVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ClientVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ExportableKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportableKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ExportableKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportableKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) InternalValue() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) RequireAlternateSignatureAlgorithm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAlternateSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) RequireAlternateSignatureAlgorithmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAlternateSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) RequireSameKeyRenewal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSameKeyRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) RequireSameKeyRenewalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSameKeyRenewalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) StrongKeyProtectionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strongKeyProtectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) StrongKeyProtectionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strongKeyProtectionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) UseLegacyProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) UseLegacyProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacyProviderInput",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetClientVersion(val *string) {
	if err := j.validateSetClientVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientVersion",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetExportableKey(val interface{}) {
	if err := j.validateSetExportableKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportableKey",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetInternalValue(val *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetRequireAlternateSignatureAlgorithm(val interface{}) {
	if err := j.validateSetRequireAlternateSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAlternateSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetRequireSameKeyRenewal(val interface{}) {
	if err := j.validateSetRequireSameKeyRenewalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSameKeyRenewal",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetStrongKeyProtectionRequired(val interface{}) {
	if err := j.validateSetStrongKeyProtectionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strongKeyProtectionRequired",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference)SetUseLegacyProvider(val interface{}) {
	if err := j.validateSetUseLegacyProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLegacyProvider",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ResetExportableKey() {
	_jsii_.InvokeVoid(
		p,
		"resetExportableKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ResetRequireAlternateSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireAlternateSignatureAlgorithm",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ResetRequireSameKeyRenewal() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireSameKeyRenewal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ResetStrongKeyProtectionRequired() {
	_jsii_.InvokeVoid(
		p,
		"resetStrongKeyProtectionRequired",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ResetUseLegacyProvider() {
	_jsii_.InvokeVoid(
		p,
		"resetUseLegacyProvider",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

