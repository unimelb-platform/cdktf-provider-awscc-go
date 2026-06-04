package ec2vpnconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2vpnconnection/internal"
)

type Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogEnabled() interface{}
	SetLogEnabled(val interface{})
	LogEnabledInput() interface{}
	LogGroupArn() *string
	SetLogGroupArn(val *string)
	LogGroupArnInput() *string
	LogOutputFormat() *string
	SetLogOutputFormat(val *string)
	LogOutputFormatInput() *string
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
	ResetLogEnabled()
	ResetLogGroupArn()
	ResetLogOutputFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference
type jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) LogOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference{}

	_jsii_.Create(
		"awscc.ec2VpnConnection.Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference_Override(e Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VpnConnection.Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetLogEnabled(val interface{}) {
	if err := j.validateSetLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEnabled",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetLogGroupArn(val *string) {
	if err := j.validateSetLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupArn",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetLogOutputFormat(val *string) {
	if err := j.validateSetLogOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ResetLogEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLogEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ResetLogGroupArn() {
	_jsii_.InvokeVoid(
		e,
		"resetLogGroupArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ResetLogOutputFormat() {
	_jsii_.InvokeVoid(
		e,
		"resetLogOutputFormat",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

