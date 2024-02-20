package applicationinsightsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationinsightsapplication/internal"
)

type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference interface {
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
	HostPort() *string
	SetHostPort(val *string)
	HostPortInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jmxurl() *string
	SetJmxurl(val *string)
	JmxurlInput() *string
	PrometheusPort() *string
	SetPrometheusPort(val *string)
	PrometheusPortInput() *string
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
	ResetHostPort()
	ResetJmxurl()
	ResetPrometheusPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference
type jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) HostPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) HostPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) Jmxurl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jmxurl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) JmxurlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jmxurlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) PrometheusPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prometheusPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) PrometheusPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prometheusPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference{}

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference_Override(a ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationinsightsApplication.ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetHostPort(val *string) {
	if err := j.validateSetHostPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPort",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetJmxurl(val *string) {
	if err := j.validateSetJmxurlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jmxurl",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetPrometheusPort(val *string) {
	if err := j.validateSetPrometheusPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prometheusPort",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ResetHostPort() {
	_jsii_.InvokeVoid(
		a,
		"resetHostPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ResetJmxurl() {
	_jsii_.InvokeVoid(
		a,
		"resetJmxurl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ResetPrometheusPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPrometheusPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

