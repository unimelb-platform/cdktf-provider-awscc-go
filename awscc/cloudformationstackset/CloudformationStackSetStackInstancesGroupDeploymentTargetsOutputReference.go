package cloudformationstackset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudformationstackset/internal"
)

type CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference interface {
	cdktf.ComplexObject
	AccountFilterType() *string
	SetAccountFilterType(val *string)
	AccountFilterTypeInput() *string
	Accounts() *[]*string
	SetAccounts(val *[]*string)
	AccountsInput() *[]*string
	AccountsUrl() *string
	SetAccountsUrl(val *string)
	AccountsUrlInput() *string
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
	InternalValue() *CloudformationStackSetStackInstancesGroupDeploymentTargets
	SetInternalValue(val *CloudformationStackSetStackInstancesGroupDeploymentTargets)
	OrganizationalUnitIds() *[]*string
	SetOrganizationalUnitIds(val *[]*string)
	OrganizationalUnitIdsInput() *[]*string
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
	ResetAccountFilterType()
	ResetAccounts()
	ResetAccountsUrl()
	ResetOrganizationalUnitIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference
type jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) AccountFilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) AccountFilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) AccountsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) AccountsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) InternalValue() *CloudformationStackSetStackInstancesGroupDeploymentTargets {
	var returns *CloudformationStackSetStackInstancesGroupDeploymentTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) OrganizationalUnitIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) OrganizationalUnitIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference{}

	_jsii_.Create(
		"awscc.cloudformationStackSet.CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference_Override(c CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudformationStackSet.CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetAccountFilterType(val *string) {
	if err := j.validateSetAccountFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountFilterType",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetAccountsUrl(val *string) {
	if err := j.validateSetAccountsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountsUrl",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetInternalValue(val *CloudformationStackSetStackInstancesGroupDeploymentTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetOrganizationalUnitIds(val *[]*string) {
	if err := j.validateSetOrganizationalUnitIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitIds",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ResetAccountFilterType() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountFilterType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ResetAccountsUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountsUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ResetOrganizationalUnitIds() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganizationalUnitIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetStackInstancesGroupDeploymentTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

