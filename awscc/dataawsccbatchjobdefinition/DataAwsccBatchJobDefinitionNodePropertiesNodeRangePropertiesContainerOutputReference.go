package dataawsccbatchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccbatchjobdefinition/internal"
)

type DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference interface {
	cdktf.ComplexObject
	Command() *[]*string
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
	Environment() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	EphemeralStorage() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference
	ExecutionRoleArn() *string
	FargatePlatformConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference
	// Experimental.
	Fqn() *string
	Image() *string
	InstanceType() *string
	InternalValue() *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer
	SetInternalValue(val *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer)
	JobRoleArn() *string
	LinuxParameters() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference
	LogConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference
	Memory() *float64
	MountPoints() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	NetworkConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference
	Privileged() cdktf.IResolvable
	ReadonlyRootFilesystem() cdktf.IResolvable
	ResourceRequirements() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	RuntimePlatform() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference
	Secrets() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ulimits() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	User() *string
	Vcpus() *float64
	Volumes() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
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

// The jsii proxy struct for DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference
type jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Environment() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) EphemeralStorage() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) FargatePlatformConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference
	_jsii_.Get(
		j,
		"fargatePlatformConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InternalValue() *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer {
	var returns *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) JobRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LinuxParameters() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LogConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) MountPoints() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) NetworkConfiguration() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Privileged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ReadonlyRootFilesystem() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResourceRequirements() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) RuntimePlatform() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Secrets() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Ulimits() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Volumes() DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList {
	var returns DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference_Override(d DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetInternalValue(val *DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

