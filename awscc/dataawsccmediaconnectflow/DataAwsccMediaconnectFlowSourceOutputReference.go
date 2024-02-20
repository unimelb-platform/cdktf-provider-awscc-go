package dataawsccmediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccmediaconnectflow/internal"
)

type DataAwsccMediaconnectFlowSourceOutputReference interface {
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
	Decryption() DataAwsccMediaconnectFlowSourceDecryptionOutputReference
	Description() *string
	EntitlementArn() *string
	// Experimental.
	Fqn() *string
	GatewayBridgeSource() DataAwsccMediaconnectFlowSourceGatewayBridgeSourceOutputReference
	IngestIp() *string
	IngestPort() *float64
	InternalValue() *DataAwsccMediaconnectFlowSource
	SetInternalValue(val *DataAwsccMediaconnectFlowSource)
	MaxBitrate() *float64
	MaxLatency() *float64
	MinLatency() *float64
	Name() *string
	Protocol() *string
	SenderControlPort() *float64
	SenderIpAddress() *string
	SourceArn() *string
	SourceIngestPort() *string
	SourceListenerAddress() *string
	SourceListenerPort() *float64
	StreamId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcInterfaceName() *string
	WhitelistCidr() *string
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

// The jsii proxy struct for DataAwsccMediaconnectFlowSourceOutputReference
type jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Decryption() DataAwsccMediaconnectFlowSourceDecryptionOutputReference {
	var returns DataAwsccMediaconnectFlowSourceDecryptionOutputReference
	_jsii_.Get(
		j,
		"decryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) EntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GatewayBridgeSource() DataAwsccMediaconnectFlowSourceGatewayBridgeSourceOutputReference {
	var returns DataAwsccMediaconnectFlowSourceGatewayBridgeSourceOutputReference
	_jsii_.Get(
		j,
		"gatewayBridgeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) IngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) IngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) InternalValue() *DataAwsccMediaconnectFlowSource {
	var returns *DataAwsccMediaconnectFlowSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SenderControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SenderIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SourceListenerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) SourceListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) VpcInterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) WhitelistCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidr",
		&returns,
	)
	return returns
}


func NewDataAwsccMediaconnectFlowSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccMediaconnectFlowSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediaconnectFlowSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccMediaconnectFlow.DataAwsccMediaconnectFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccMediaconnectFlowSourceOutputReference_Override(d DataAwsccMediaconnectFlowSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccMediaconnectFlow.DataAwsccMediaconnectFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference)SetInternalValue(val *DataAwsccMediaconnectFlowSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediaconnectFlowSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

