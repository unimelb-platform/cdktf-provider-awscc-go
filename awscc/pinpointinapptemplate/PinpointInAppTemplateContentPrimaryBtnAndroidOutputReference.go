package pinpointinapptemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pinpointinapptemplate/internal"
)

type PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference interface {
	cdktf.ComplexObject
	ButtonAction() *string
	SetButtonAction(val *string)
	ButtonActionInput() *string
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
	Link() *string
	SetLink(val *string)
	LinkInput() *string
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
	ResetButtonAction()
	ResetLink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference
type jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ButtonAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ButtonActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) Link() *string {
	var returns *string
	_jsii_.Get(
		j,
		"link",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) LinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPinpointInAppTemplateContentPrimaryBtnAndroidOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointInAppTemplateContentPrimaryBtnAndroidOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference{}

	_jsii_.Create(
		"awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPinpointInAppTemplateContentPrimaryBtnAndroidOutputReference_Override(p PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetButtonAction(val *string) {
	if err := j.validateSetButtonActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buttonAction",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetLink(val *string) {
	if err := j.validateSetLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"link",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ResetButtonAction() {
	_jsii_.InvokeVoid(
		p,
		"resetButtonAction",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ResetLink() {
	_jsii_.InvokeVoid(
		p,
		"resetLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

