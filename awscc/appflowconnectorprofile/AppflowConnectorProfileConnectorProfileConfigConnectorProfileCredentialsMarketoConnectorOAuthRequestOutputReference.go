package appflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appflowconnectorprofile/internal"
)

type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference interface {
	cdktf.ComplexObject
	AuthCode() *string
	SetAuthCode(val *string)
	AuthCodeInput() *string
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
	RedirectUri() *string
	SetRedirectUri(val *string)
	RedirectUriInput() *string
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
	ResetAuthCode()
	ResetRedirectUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference
type jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) RedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) RedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference{}

	_jsii_.Create(
		"awscc.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference_Override(a AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetRedirectUri(val *string) {
	if err := j.validateSetRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUri",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ResetRedirectUri() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
