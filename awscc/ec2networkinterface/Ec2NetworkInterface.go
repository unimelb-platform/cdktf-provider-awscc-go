package ec2networkinterface

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2networkinterface/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface awscc_ec2_network_interface}.
type Ec2NetworkInterface interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionTrackingSpecification() Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference
	ConnectionTrackingSpecificationInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnablePrimaryIpv6() interface{}
	SetEnablePrimaryIpv6(val interface{})
	EnablePrimaryIpv6Input() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupSet() *[]*string
	SetGroupSet(val *[]*string)
	GroupSetInput() *[]*string
	Id() *string
	InterfaceType() *string
	SetInterfaceType(val *string)
	InterfaceTypeInput() *string
	Ipv4PrefixCount() *float64
	SetIpv4PrefixCount(val *float64)
	Ipv4PrefixCountInput() *float64
	Ipv4Prefixes() Ec2NetworkInterfaceIpv4PrefixesList
	Ipv4PrefixesInput() interface{}
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() Ec2NetworkInterfaceIpv6AddressesList
	Ipv6AddressesInput() interface{}
	Ipv6PrefixCount() *float64
	SetIpv6PrefixCount(val *float64)
	Ipv6PrefixCountInput() *float64
	Ipv6Prefixes() Ec2NetworkInterfaceIpv6PrefixesList
	Ipv6PrefixesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrimaryIpv6Address() *string
	PrimaryPrivateIpAddress() *string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddresses() Ec2NetworkInterfacePrivateIpAddressesList
	PrivateIpAddressesInput() interface{}
	PrivateIpAddressInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SecondaryPrivateIpAddressCount() *float64
	SetSecondaryPrivateIpAddressCount(val *float64)
	SecondaryPrivateIpAddressCountInput() *float64
	SecondaryPrivateIpAddresses() *[]*string
	SourceDestCheck() interface{}
	SetSourceDestCheck(val interface{})
	SourceDestCheckInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() Ec2NetworkInterfaceTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutConnectionTrackingSpecification(value *Ec2NetworkInterfaceConnectionTrackingSpecification)
	PutIpv4Prefixes(value interface{})
	PutIpv6Addresses(value interface{})
	PutIpv6Prefixes(value interface{})
	PutPrivateIpAddresses(value interface{})
	PutTags(value interface{})
	ResetConnectionTrackingSpecification()
	ResetDescription()
	ResetEnablePrimaryIpv6()
	ResetGroupSet()
	ResetInterfaceType()
	ResetIpv4PrefixCount()
	ResetIpv4Prefixes()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetIpv6PrefixCount()
	ResetIpv6Prefixes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateIpAddress()
	ResetPrivateIpAddresses()
	ResetSecondaryPrivateIpAddressCount()
	ResetSourceDestCheck()
	ResetTags()
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

// The jsii proxy struct for Ec2NetworkInterface
type jsiiProxy_Ec2NetworkInterface struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2NetworkInterface) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) ConnectionTrackingSpecification() Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference {
	var returns Ec2NetworkInterfaceConnectionTrackingSpecificationOutputReference
	_jsii_.Get(
		j,
		"connectionTrackingSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) ConnectionTrackingSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTrackingSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) EnablePrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) EnablePrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) GroupSet() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) GroupSetInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv4PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv4Prefixes() Ec2NetworkInterfaceIpv4PrefixesList {
	var returns Ec2NetworkInterfaceIpv4PrefixesList
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv4PrefixesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6Addresses() Ec2NetworkInterfaceIpv6AddressesList {
	var returns Ec2NetworkInterfaceIpv6AddressesList
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6Prefixes() Ec2NetworkInterfaceIpv6PrefixesList {
	var returns Ec2NetworkInterfaceIpv6PrefixesList
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Ipv6PrefixesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrimaryIpv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrimaryPrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPrivateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrivateIpAddresses() Ec2NetworkInterfacePrivateIpAddressesList {
	var returns Ec2NetworkInterfacePrivateIpAddressesList
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrivateIpAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SecondaryPrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SecondaryPrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SecondaryPrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) Tags() Ec2NetworkInterfaceTagsList {
	var returns Ec2NetworkInterfaceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInterface) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface awscc_ec2_network_interface} Resource.
func NewEc2NetworkInterface(scope constructs.Construct, id *string, config *Ec2NetworkInterfaceConfig) Ec2NetworkInterface {
	_init_.Initialize()

	if err := validateNewEc2NetworkInterfaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInterface{}

	_jsii_.Create(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface awscc_ec2_network_interface} Resource.
func NewEc2NetworkInterface_Override(e Ec2NetworkInterface, scope constructs.Construct, id *string, config *Ec2NetworkInterfaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetEnablePrimaryIpv6(val interface{}) {
	if err := j.validateSetEnablePrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrimaryIpv6",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetGroupSet(val *[]*string) {
	if err := j.validateSetGroupSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetIpv4PrefixCount(val *float64) {
	if err := j.validateSetIpv4PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixCount",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetIpv6PrefixCount(val *float64) {
	if err := j.validateSetIpv6PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixCount",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetSecondaryPrivateIpAddressCount(val *float64) {
	if err := j.validateSetSecondaryPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInterface)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

// Generates CDKTF code for importing a Ec2NetworkInterface resource upon running "cdktf plan <stack-name>".
func Ec2NetworkInterface_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2NetworkInterface_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
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
func Ec2NetworkInterface_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2NetworkInterface_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2NetworkInterface_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2NetworkInterface_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2NetworkInterface_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2NetworkInterface_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2NetworkInterface_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2NetworkInterface.Ec2NetworkInterface",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInterface) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInterface) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutConnectionTrackingSpecification(value *Ec2NetworkInterfaceConnectionTrackingSpecification) {
	if err := e.validatePutConnectionTrackingSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectionTrackingSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutIpv4Prefixes(value interface{}) {
	if err := e.validatePutIpv4PrefixesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv4Prefixes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutIpv6Addresses(value interface{}) {
	if err := e.validatePutIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Addresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutIpv6Prefixes(value interface{}) {
	if err := e.validatePutIpv6PrefixesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Prefixes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutPrivateIpAddresses(value interface{}) {
	if err := e.validatePutPrivateIpAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateIpAddresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetConnectionTrackingSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectionTrackingSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetEnablePrimaryIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetEnablePrimaryIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetGroupSet() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupSet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv4PrefixCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4PrefixCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv4Prefixes() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4Prefixes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv6PrefixCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6PrefixCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetIpv6Prefixes() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Prefixes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetPrivateIpAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetSecondaryPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetSecondaryPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2NetworkInterface) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInterface) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

