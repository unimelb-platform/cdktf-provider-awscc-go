package dataawsccec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccec2networkinsightsanalysis/internal"
)

type DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference interface {
	cdktf.ComplexObject
	AclRule() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsAclRuleOutputReference
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
	Component() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsComponentOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationVpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcOutputReference
	// Experimental.
	Fqn() *string
	InboundHeader() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderOutputReference
	InternalValue() *DataAwsccEc2NetworkInsightsAnalysisForwardPathComponents
	SetInternalValue(val *DataAwsccEc2NetworkInsightsAnalysisForwardPathComponents)
	OutboundHeader() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderOutputReference
	RouteTableRoute() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteOutputReference
	SecurityGroupRule() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleOutputReference
	SequenceNumber() *float64
	SourceVpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSourceVpcOutputReference
	Subnet() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSubnetOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Vpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsVpcOutputReference
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

// The jsii proxy struct for DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference
type jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) AclRule() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsAclRuleOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsAclRuleOutputReference
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Component() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsComponentOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsComponentOutputReference
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) DestinationVpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcOutputReference
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InboundHeader() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderOutputReference
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InternalValue() *DataAwsccEc2NetworkInsightsAnalysisForwardPathComponents {
	var returns *DataAwsccEc2NetworkInsightsAnalysisForwardPathComponents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) OutboundHeader() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderOutputReference
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) RouteTableRoute() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteOutputReference
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SecurityGroupRule() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleOutputReference
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SourceVpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSourceVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSourceVpcOutputReference
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Subnet() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSubnetOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsSubnetOutputReference
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Vpc() DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsVpcOutputReference {
	var returns DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsVpcOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccEc2NetworkInsightsAnalysis.DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference_Override(d DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEc2NetworkInsightsAnalysis.DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetInternalValue(val *DataAwsccEc2NetworkInsightsAnalysisForwardPathComponents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

