package dataawsccappflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccappflowconnectorprofile/internal"
)

type DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference interface {
	cdktf.ComplexObject
	AccessToken() *string
	ClientId() *string
	ClientSecret() *string
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
	InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2
	SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2)
	OAuthRequest() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OAuthRequestOutputReference
	RefreshToken() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference
type jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 {
	var returns *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) OAuthRequest() DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OAuthRequestOutputReference {
	var returns DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OAuthRequestOutputReference
	_jsii_.Get(
		j,
		"oAuthRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference_Override(d DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference)SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
