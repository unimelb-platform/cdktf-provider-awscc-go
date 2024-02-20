package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/batchjobdefinition/internal"
)

type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference interface {
	cdktf.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Environment() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	EnvironmentInput() interface{}
	EphemeralStorage() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference
	EphemeralStorageInput() interface{}
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	FargatePlatformConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference
	FargatePlatformConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobRoleArn() *string
	SetJobRoleArn(val *string)
	JobRoleArnInput() *string
	LinuxParameters() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference
	LinuxParametersInput() interface{}
	LogConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MountPoints() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	MountPointsInput() interface{}
	NetworkConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	ResourceRequirements() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	ResourceRequirementsInput() interface{}
	RuntimePlatform() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference
	RuntimePlatformInput() interface{}
	Secrets() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	SecretsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ulimits() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	UlimitsInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	Vcpus() *float64
	SetVcpus(val *float64)
	VcpusInput() *float64
	Volumes() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
	VolumesInput() interface{}
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
	PutEnvironment(value interface{})
	PutEphemeralStorage(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorage)
	PutFargatePlatformConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfiguration)
	PutLinuxParameters(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParameters)
	PutLogConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfiguration)
	PutMountPoints(value interface{})
	PutNetworkConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfiguration)
	PutResourceRequirements(value interface{})
	PutRuntimePlatform(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatform)
	PutSecrets(value interface{})
	PutUlimits(value interface{})
	PutVolumes(value interface{})
	ResetCommand()
	ResetEnvironment()
	ResetEphemeralStorage()
	ResetExecutionRoleArn()
	ResetFargatePlatformConfiguration()
	ResetInstanceType()
	ResetJobRoleArn()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMemory()
	ResetMountPoints()
	ResetNetworkConfiguration()
	ResetPrivileged()
	ResetReadonlyRootFilesystem()
	ResetResourceRequirements()
	ResetRuntimePlatform()
	ResetSecrets()
	ResetUlimits()
	ResetUser()
	ResetVcpus()
	ResetVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference
type jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Environment() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) EphemeralStorage() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) EphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) FargatePlatformConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationOutputReference
	_jsii_.Get(
		j,
		"fargatePlatformConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) FargatePlatformConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fargatePlatformConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) JobRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) JobRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LinuxParameters() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LogConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) MountPoints() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) NetworkConfiguration() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResourceRequirements() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) RuntimePlatform() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformOutputReference
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) RuntimePlatformInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Secrets() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Ulimits() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) UlimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) VcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Volumes() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference{}

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference_Override(b BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetJobRoleArn(val *string) {
	if err := j.validateSetJobRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobRoleArn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetVcpus(val *float64) {
	if err := j.validateSetVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpus",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutEnvironment(value interface{}) {
	if err := b.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutEphemeralStorage(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorage) {
	if err := b.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutFargatePlatformConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfiguration) {
	if err := b.validatePutFargatePlatformConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFargatePlatformConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutLinuxParameters(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParameters) {
	if err := b.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutLogConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfiguration) {
	if err := b.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutMountPoints(value interface{}) {
	if err := b.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutNetworkConfiguration(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfiguration) {
	if err := b.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutResourceRequirements(value interface{}) {
	if err := b.validatePutResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResourceRequirements",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutRuntimePlatform(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatform) {
	if err := b.validatePutRuntimePlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRuntimePlatform",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutSecrets(value interface{}) {
	if err := b.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSecrets",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutUlimits(value interface{}) {
	if err := b.validatePutUlimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUlimits",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) PutVolumes(value interface{}) {
	if err := b.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVolumes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		b,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetFargatePlatformConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetFargatePlatformConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetJobRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetJobRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		b,
		"resetMemory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		b,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		b,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetResourceRequirements() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceRequirements",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetRuntimePlatform() {
	_jsii_.InvokeVoid(
		b,
		"resetRuntimePlatform",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		b,
		"resetSecrets",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetUlimits() {
	_jsii_.InvokeVoid(
		b,
		"resetUlimits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		b,
		"resetUser",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetVcpus() {
	_jsii_.InvokeVoid(
		b,
		"resetVcpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

