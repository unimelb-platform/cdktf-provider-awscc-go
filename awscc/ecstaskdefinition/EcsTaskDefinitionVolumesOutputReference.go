package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ecstaskdefinition/internal"
)

type EcsTaskDefinitionVolumesOutputReference interface {
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
	ConfiguredAtLaunch() interface{}
	SetConfiguredAtLaunch(val interface{})
	ConfiguredAtLaunchInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerVolumeConfiguration() EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference
	DockerVolumeConfigurationInput() interface{}
	EfsVolumeConfiguration() EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference
	EfsVolumeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	Host() EcsTaskDefinitionVolumesHostOutputReference
	HostInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutDockerVolumeConfiguration(value *EcsTaskDefinitionVolumesDockerVolumeConfiguration)
	PutEfsVolumeConfiguration(value *EcsTaskDefinitionVolumesEfsVolumeConfiguration)
	PutHost(value *EcsTaskDefinitionVolumesHost)
	ResetConfiguredAtLaunch()
	ResetDockerVolumeConfiguration()
	ResetEfsVolumeConfiguration()
	ResetHost()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsTaskDefinitionVolumesOutputReference
type jsiiProxy_EcsTaskDefinitionVolumesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ConfiguredAtLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuredAtLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ConfiguredAtLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuredAtLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) DockerVolumeConfiguration() EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"dockerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) DockerVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dockerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) EfsVolumeConfiguration() EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"efsVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) EfsVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Host() EcsTaskDefinitionVolumesHostOutputReference {
	var returns EcsTaskDefinitionVolumesHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionVolumesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsTaskDefinitionVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionVolumesOutputReference{}

	_jsii_.Create(
		"awscc.ecsTaskDefinition.EcsTaskDefinitionVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionVolumesOutputReference_Override(e EcsTaskDefinitionVolumesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ecsTaskDefinition.EcsTaskDefinitionVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetConfiguredAtLaunch(val interface{}) {
	if err := j.validateSetConfiguredAtLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuredAtLaunch",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutDockerVolumeConfiguration(value *EcsTaskDefinitionVolumesDockerVolumeConfiguration) {
	if err := e.validatePutDockerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDockerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutEfsVolumeConfiguration(value *EcsTaskDefinitionVolumesEfsVolumeConfiguration) {
	if err := e.validatePutEfsVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEfsVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutHost(value *EcsTaskDefinitionVolumesHost) {
	if err := e.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHost",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetConfiguredAtLaunch() {
	_jsii_.InvokeVoid(
		e,
		"resetConfiguredAtLaunch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetDockerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetDockerVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetEfsVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetEfsVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		e,
		"resetHost",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

