package groundstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/groundstationconfig/internal"
)

type GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference interface {
	cdktf.ComplexObject
	Bandwidth() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigBandwidthOutputReference
	BandwidthInput() interface{}
	CenterFrequency() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigCenterFrequencyOutputReference
	CenterFrequencyInput() interface{}
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
	Polarization() *string
	SetPolarization(val *string)
	PolarizationInput() *string
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
	PutBandwidth(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigBandwidth)
	PutCenterFrequency(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigCenterFrequency)
	ResetBandwidth()
	ResetCenterFrequency()
	ResetPolarization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference
type jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) Bandwidth() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigBandwidthOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigBandwidthOutputReference
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) BandwidthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) CenterFrequency() GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigCenterFrequencyOutputReference {
	var returns GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigCenterFrequencyOutputReference
	_jsii_.Get(
		j,
		"centerFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) CenterFrequencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"centerFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) Polarization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) PolarizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polarizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference{}

	_jsii_.Create(
		"awscc.groundstationConfig.GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference_Override(g GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.groundstationConfig.GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetPolarization(val *string) {
	if err := j.validateSetPolarizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"polarization",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) PutBandwidth(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigBandwidth) {
	if err := g.validatePutBandwidthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBandwidth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) PutCenterFrequency(value *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigCenterFrequency) {
	if err := g.validatePutCenterFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCenterFrequency",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ResetBandwidth() {
	_jsii_.InvokeVoid(
		g,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ResetCenterFrequency() {
	_jsii_.InvokeVoid(
		g,
		"resetCenterFrequency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ResetPolarization() {
	_jsii_.InvokeVoid(
		g,
		"resetPolarization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

