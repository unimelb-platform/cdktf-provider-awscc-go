package dataawsccecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccecstaskdefinition/internal"
)

type DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference interface {
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
	Cpu() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CredentialSpecs() *[]*string
	DependsOn() DataAwsccEcsTaskDefinitionContainerDefinitionsDependsOnList
	DisableNetworking() cdktf.IResolvable
	DnsSearchDomains() *[]*string
	DnsServers() *[]*string
	DockerLabels() cdktf.StringMap
	DockerSecurityOptions() *[]*string
	EntryPoint() *[]*string
	Environment() DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentList
	EnvironmentFiles() DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentFilesList
	Essential() cdktf.IResolvable
	ExtraHosts() DataAwsccEcsTaskDefinitionContainerDefinitionsExtraHostsList
	FirelensConfiguration() DataAwsccEcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	// Experimental.
	Fqn() *string
	HealthCheck() DataAwsccEcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	Hostname() *string
	Image() *string
	Interactive() cdktf.IResolvable
	InternalValue() *DataAwsccEcsTaskDefinitionContainerDefinitions
	SetInternalValue(val *DataAwsccEcsTaskDefinitionContainerDefinitions)
	Links() *[]*string
	LinuxParameters() DataAwsccEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	LogConfiguration() DataAwsccEcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	Memory() *float64
	MemoryReservation() *float64
	MountPoints() DataAwsccEcsTaskDefinitionContainerDefinitionsMountPointsList
	Name() *string
	PortMappings() DataAwsccEcsTaskDefinitionContainerDefinitionsPortMappingsList
	Privileged() cdktf.IResolvable
	PseudoTerminal() cdktf.IResolvable
	ReadonlyRootFilesystem() cdktf.IResolvable
	RepositoryCredentials() DataAwsccEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	ResourceRequirements() DataAwsccEcsTaskDefinitionContainerDefinitionsResourceRequirementsList
	Secrets() DataAwsccEcsTaskDefinitionContainerDefinitionsSecretsList
	StartTimeout() *float64
	StopTimeout() *float64
	SystemControls() DataAwsccEcsTaskDefinitionContainerDefinitionsSystemControlsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ulimits() DataAwsccEcsTaskDefinitionContainerDefinitionsUlimitsList
	User() *string
	VolumesFrom() DataAwsccEcsTaskDefinitionContainerDefinitionsVolumesFromList
	WorkingDirectory() *string
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

// The jsii proxy struct for DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference
type jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) CredentialSpecs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"credentialSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DependsOn() DataAwsccEcsTaskDefinitionContainerDefinitionsDependsOnList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DisableNetworking() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DnsSearchDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSearchDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DockerLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"dockerLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) DockerSecurityOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dockerSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Environment() DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFiles() DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentFilesList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsEnvironmentFilesList
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Essential() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ExtraHosts() DataAwsccEcsTaskDefinitionContainerDefinitionsExtraHostsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsExtraHostsList
	_jsii_.Get(
		j,
		"extraHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) FirelensConfiguration() DataAwsccEcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) HealthCheck() DataAwsccEcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Interactive() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) InternalValue() *DataAwsccEcsTaskDefinitionContainerDefinitions {
	var returns *DataAwsccEcsTaskDefinitionContainerDefinitions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Links() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) LinuxParameters() DataAwsccEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) LogConfiguration() DataAwsccEcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) MountPoints() DataAwsccEcsTaskDefinitionContainerDefinitionsMountPointsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) PortMappings() DataAwsccEcsTaskDefinitionContainerDefinitionsPortMappingsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsPortMappingsList
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Privileged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystem() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentials() DataAwsccEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ResourceRequirements() DataAwsccEcsTaskDefinitionContainerDefinitionsResourceRequirementsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Secrets() DataAwsccEcsTaskDefinitionContainerDefinitionsSecretsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) SystemControls() DataAwsccEcsTaskDefinitionContainerDefinitionsSystemControlsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsSystemControlsList
	_jsii_.Get(
		j,
		"systemControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Ulimits() DataAwsccEcsTaskDefinitionContainerDefinitionsUlimitsList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) VolumesFrom() DataAwsccEcsTaskDefinitionContainerDefinitionsVolumesFromList {
	var returns DataAwsccEcsTaskDefinitionContainerDefinitionsVolumesFromList
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


func NewDataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEcsTaskDefinitionContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccEcsTaskDefinition.DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference_Override(d DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEcsTaskDefinition.DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference)SetInternalValue(val *DataAwsccEcsTaskDefinitionContainerDefinitions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEcsTaskDefinitionContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

