package ec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2spotfleet/internal"
)

type Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference interface {
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
	LaunchTemplateId() *string
	SetLaunchTemplateId(val *string)
	LaunchTemplateIdInput() *string
	LaunchTemplateName() *string
	SetLaunchTemplateName(val *string)
	LaunchTemplateNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetLaunchTemplateId()
	ResetLaunchTemplateName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference
type jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) LaunchTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) LaunchTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference{}

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference_Override(e Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetLaunchTemplateId(val *string) {
	if err := j.validateSetLaunchTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateId",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetLaunchTemplateName(val *string) {
	if err := j.validateSetLaunchTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateName",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ResetLaunchTemplateId() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchTemplateId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ResetLaunchTemplateName() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchTemplateName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

