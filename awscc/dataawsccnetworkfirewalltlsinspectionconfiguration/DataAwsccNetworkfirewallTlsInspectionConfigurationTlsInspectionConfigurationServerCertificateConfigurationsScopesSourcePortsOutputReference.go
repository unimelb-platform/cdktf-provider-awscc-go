package dataawsccnetworkfirewalltlsinspectionconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccnetworkfirewalltlsinspectionconfiguration/internal"
)

type DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference interface {
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
	FromPort() *float64
	InternalValue() *DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePorts
	SetInternalValue(val *DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePorts)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToPort() *float64
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

// The jsii proxy struct for DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference
type jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) InternalValue() *DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePorts {
	var returns *DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePorts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}


func NewDataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccNetworkfirewallTlsInspectionConfiguration.DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference_Override(d DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccNetworkfirewallTlsInspectionConfiguration.DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference)SetInternalValue(val *DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePorts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccNetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesSourcePortsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

