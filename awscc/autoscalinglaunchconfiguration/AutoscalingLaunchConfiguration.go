package autoscalinglaunchconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/autoscalinglaunchconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_launch_configuration awscc_autoscaling_launch_configuration}.
type AutoscalingLaunchConfiguration interface {
	cdktf.TerraformResource
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
	BlockDeviceMappings() AutoscalingLaunchConfigurationBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClassicLinkVpcId() *string
	SetClassicLinkVpcId(val *string)
	ClassicLinkVpcIdInput() *string
	ClassicLinkVpcSecurityGroups() *[]*string
	SetClassicLinkVpcSecurityGroups(val *[]*string)
	ClassicLinkVpcSecurityGroupsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InstanceMonitoring() interface{}
	SetInstanceMonitoring(val interface{})
	InstanceMonitoringInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	KernelId() *string
	SetKernelId(val *string)
	KernelIdInput() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
	LaunchConfigurationNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetadataOptions() AutoscalingLaunchConfigurationMetadataOptionsOutputReference
	MetadataOptionsInput() interface{}
	// The tree node.
	Node() constructs.Node
	PlacementTenancy() *string
	SetPlacementTenancy(val *string)
	PlacementTenancyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RamDiskId() *string
	SetRamDiskId(val *string)
	RamDiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SpotPrice() *string
	SetSpotPrice(val *string)
	SpotPriceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBlockDeviceMappings(value interface{})
	PutMetadataOptions(value *AutoscalingLaunchConfigurationMetadataOptions)
	ResetAssociatePublicIpAddress()
	ResetBlockDeviceMappings()
	ResetClassicLinkVpcId()
	ResetClassicLinkVpcSecurityGroups()
	ResetEbsOptimized()
	ResetIamInstanceProfile()
	ResetInstanceId()
	ResetInstanceMonitoring()
	ResetKernelId()
	ResetKeyName()
	ResetLaunchConfigurationName()
	ResetMetadataOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementTenancy()
	ResetRamDiskId()
	ResetSecurityGroups()
	ResetSpotPrice()
	ResetUserData()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutoscalingLaunchConfiguration
type jsiiProxy_AutoscalingLaunchConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) BlockDeviceMappings() AutoscalingLaunchConfigurationBlockDeviceMappingsList {
	var returns AutoscalingLaunchConfigurationBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ClassicLinkVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicLinkVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ClassicLinkVpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicLinkVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ClassicLinkVpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classicLinkVpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ClassicLinkVpcSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classicLinkVpcSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) LaunchConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) MetadataOptions() AutoscalingLaunchConfigurationMetadataOptionsOutputReference {
	var returns AutoscalingLaunchConfigurationMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) MetadataOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) PlacementTenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) RamDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_launch_configuration awscc_autoscaling_launch_configuration} Resource.
func NewAutoscalingLaunchConfiguration(scope constructs.Construct, id *string, config *AutoscalingLaunchConfigurationConfig) AutoscalingLaunchConfiguration {
	_init_.Initialize()

	if err := validateNewAutoscalingLaunchConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingLaunchConfiguration{}

	_jsii_.Create(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_launch_configuration awscc_autoscaling_launch_configuration} Resource.
func NewAutoscalingLaunchConfiguration_Override(a AutoscalingLaunchConfiguration, scope constructs.Construct, id *string, config *AutoscalingLaunchConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetClassicLinkVpcId(val *string) {
	if err := j.validateSetClassicLinkVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classicLinkVpcId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetClassicLinkVpcSecurityGroups(val *[]*string) {
	if err := j.validateSetClassicLinkVpcSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classicLinkVpcSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetInstanceMonitoring(val interface{}) {
	if err := j.validateSetInstanceMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMonitoring",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetLaunchConfigurationName(val *string) {
	if err := j.validateSetLaunchConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetPlacementTenancy(val *string) {
	if err := j.validateSetPlacementTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetRamDiskId(val *string) {
	if err := j.validateSetRamDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_AutoscalingLaunchConfiguration)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Generates CDKTF code for importing a AutoscalingLaunchConfiguration resource upon running "cdktf plan <stack-name>".
func AutoscalingLaunchConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAutoscalingLaunchConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AutoscalingLaunchConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingLaunchConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingLaunchConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingLaunchConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingLaunchConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingLaunchConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutoscalingLaunchConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.autoscalingLaunchConfiguration.AutoscalingLaunchConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingLaunchConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) PutBlockDeviceMappings(value interface{}) {
	if err := a.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) PutMetadataOptions(value *AutoscalingLaunchConfigurationMetadataOptions) {
	if err := a.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetClassicLinkVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetClassicLinkVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetClassicLinkVpcSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetClassicLinkVpcSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetInstanceId() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetInstanceMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetKernelId() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetLaunchConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetPlacementTenancy() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementTenancy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetRamDiskId() {
	_jsii_.InvokeVoid(
		a,
		"resetRamDiskId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingLaunchConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

