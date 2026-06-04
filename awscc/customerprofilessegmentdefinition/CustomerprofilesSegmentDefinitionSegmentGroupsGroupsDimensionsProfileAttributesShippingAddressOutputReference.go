package customerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/customerprofilessegmentdefinition/internal"
)

type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference interface {
	cdktf.ComplexObject
	City() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCityOutputReference
	CityInput() interface{}
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
	Country() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountryOutputReference
	CountryInput() interface{}
	County() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountyOutputReference
	CountyInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PostalCode() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference
	PostalCodeInput() interface{}
	Province() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvinceOutputReference
	ProvinceInput() interface{}
	State() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressStateOutputReference
	StateInput() interface{}
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
	PutCity(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCity)
	PutCountry(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountry)
	PutCounty(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCounty)
	PutPostalCode(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode)
	PutProvince(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvince)
	PutState(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressState)
	ResetCity()
	ResetCountry()
	ResetCounty()
	ResetPostalCode()
	ResetProvince()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference
type jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) City() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCityOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCityOutputReference
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) CityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) Country() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountryOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountryOutputReference
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) CountryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) County() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountyOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountyOutputReference
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) CountyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PostalCode() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCodeOutputReference
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PostalCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) Province() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvinceOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvinceOutputReference
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ProvinceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) State() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressStateOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressStateOutputReference
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) StateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference{}

	_jsii_.Create(
		"awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference_Override(c CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutCity(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCity) {
	if err := c.validatePutCityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutCountry(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountry) {
	if err := c.validatePutCountryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCountry",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutCounty(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCounty) {
	if err := c.validatePutCountyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCounty",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutPostalCode(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode) {
	if err := c.validatePutPostalCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostalCode",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutProvince(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvince) {
	if err := c.validatePutProvinceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProvince",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) PutState(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressState) {
	if err := c.validatePutStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putState",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		c,
		"resetCity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		c,
		"resetCountry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetCounty() {
	_jsii_.InvokeVoid(
		c,
		"resetCounty",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		c,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetProvince() {
	_jsii_.InvokeVoid(
		c,
		"resetProvince",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		c,
		"resetState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

