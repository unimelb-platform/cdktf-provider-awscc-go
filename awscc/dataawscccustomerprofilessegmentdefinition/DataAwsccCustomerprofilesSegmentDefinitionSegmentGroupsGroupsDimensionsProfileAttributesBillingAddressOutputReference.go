package dataawscccustomerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccustomerprofilessegmentdefinition/internal"
)

type DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference interface {
	cdktf.ComplexObject
	City() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCityOutputReference
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
	Country() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountryOutputReference
	County() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountyOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress
	SetInternalValue(val *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress)
	PostalCode() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressPostalCodeOutputReference
	Province() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressProvinceOutputReference
	State() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressStateOutputReference
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

// The jsii proxy struct for DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference
type jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) City() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCityOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCityOutputReference
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) Country() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountryOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountryOutputReference
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) County() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountyOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountyOutputReference
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) InternalValue() *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress {
	var returns *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) PostalCode() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressPostalCodeOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressPostalCodeOutputReference
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) Province() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressProvinceOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressProvinceOutputReference
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) State() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressStateOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressStateOutputReference
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference_Override(d DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference)SetInternalValue(val *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

