package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/batchjobdefinition/internal"
)

type BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference interface {
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
	DependsOn() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList
	DependsOnInput() interface{}
	Environment() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList
	EnvironmentInput() interface{}
	Essential() interface{}
	SetEssential(val interface{})
	EssentialInput() interface{}
	FirelensConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxParameters() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference
	LinuxParametersInput() interface{}
	LogConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	MountPoints() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList
	MountPointsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	RepositoryCredentials() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference
	RepositoryCredentialsInput() interface{}
	ResourceRequirements() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList
	ResourceRequirementsInput() interface{}
	Secrets() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList
	SecretsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ulimits() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList
	UlimitsInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutDependsOn(value interface{})
	PutEnvironment(value interface{})
	PutFirelensConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfiguration)
	PutLinuxParameters(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParameters)
	PutLogConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfiguration)
	PutMountPoints(value interface{})
	PutRepositoryCredentials(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentials)
	PutResourceRequirements(value interface{})
	PutSecrets(value interface{})
	PutUlimits(value interface{})
	ResetCommand()
	ResetDependsOn()
	ResetEnvironment()
	ResetEssential()
	ResetFirelensConfiguration()
	ResetImage()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMountPoints()
	ResetName()
	ResetPrivileged()
	ResetReadonlyRootFilesystem()
	ResetRepositoryCredentials()
	ResetResourceRequirements()
	ResetSecrets()
	ResetUlimits()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference
type jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) DependsOn() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Environment() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) FirelensConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LinuxParameters() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LogConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) MountPoints() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) RepositoryCredentials() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResourceRequirements() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Secrets() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Ulimits() BatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) UlimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference{}

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference_Override(b BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutDependsOn(value interface{}) {
	if err := b.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutEnvironment(value interface{}) {
	if err := b.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutFirelensConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfiguration) {
	if err := b.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutLinuxParameters(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParameters) {
	if err := b.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutLogConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfiguration) {
	if err := b.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutMountPoints(value interface{}) {
	if err := b.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutRepositoryCredentials(value *BatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentials) {
	if err := b.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutResourceRequirements(value interface{}) {
	if err := b.validatePutResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResourceRequirements",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutSecrets(value interface{}) {
	if err := b.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSecrets",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) PutUlimits(value interface{}) {
	if err := b.validatePutUlimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUlimits",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		b,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		b,
		"resetEssential",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		b,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		b,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		b,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetResourceRequirements() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceRequirements",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		b,
		"resetSecrets",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetUlimits() {
	_jsii_.InvokeVoid(
		b,
		"resetUlimits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		b,
		"resetUser",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

