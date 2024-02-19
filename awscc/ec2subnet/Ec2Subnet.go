package ec2subnet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2subnet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet awscc_ec2_subnet}.
type Ec2Subnet interface {
	cdktf.TerraformResource
	AssignIpv6AddressOnCreation() interface{}
	SetAssignIpv6AddressOnCreation(val interface{})
	AssignIpv6AddressOnCreationInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	AvailabilityZoneIdInput() *string
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlock() *string
	SetCidrBlock(val *string)
	CidrBlockInput() *string
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
	EnableDns64() interface{}
	SetEnableDns64(val interface{})
	EnableDns64Input() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Ipv4IpamPoolId() *string
	SetIpv4IpamPoolId(val *string)
	Ipv4IpamPoolIdInput() *string
	Ipv4NetmaskLength() *float64
	SetIpv4NetmaskLength(val *float64)
	Ipv4NetmaskLengthInput() *float64
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockInput() *string
	Ipv6CidrBlocks() *[]*string
	SetIpv6CidrBlocks(val *[]*string)
	Ipv6CidrBlocksInput() *[]*string
	Ipv6IpamPoolId() *string
	SetIpv6IpamPoolId(val *string)
	Ipv6IpamPoolIdInput() *string
	Ipv6Native() interface{}
	SetIpv6Native(val interface{})
	Ipv6NativeInput() interface{}
	Ipv6NetmaskLength() *float64
	SetIpv6NetmaskLength(val *float64)
	Ipv6NetmaskLengthInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MapPublicIpOnLaunch() interface{}
	SetMapPublicIpOnLaunch(val interface{})
	MapPublicIpOnLaunchInput() interface{}
	NetworkAclAssociationId() *string
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	SetOutpostArn(val *string)
	OutpostArnInput() *string
	PrivateDnsNameOptionsOnLaunch() Ec2SubnetPrivateDnsNameOptionsOnLaunchOutputReference
	PrivateDnsNameOptionsOnLaunchInput() interface{}
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
	SubnetId() *string
	Tags() Ec2SubnetTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutPrivateDnsNameOptionsOnLaunch(value *Ec2SubnetPrivateDnsNameOptionsOnLaunch)
	PutTags(value interface{})
	ResetAssignIpv6AddressOnCreation()
	ResetAvailabilityZone()
	ResetAvailabilityZoneId()
	ResetCidrBlock()
	ResetEnableDns64()
	ResetIpv4IpamPoolId()
	ResetIpv4NetmaskLength()
	ResetIpv6CidrBlock()
	ResetIpv6CidrBlocks()
	ResetIpv6IpamPoolId()
	ResetIpv6Native()
	ResetIpv6NetmaskLength()
	ResetMapPublicIpOnLaunch()
	ResetOutpostArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateDnsNameOptionsOnLaunch()
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

// The jsii proxy struct for Ec2Subnet
type jsiiProxy_Ec2Subnet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2Subnet) AssignIpv6AddressOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) AssignIpv6AddressOnCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) EnableDns64() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) EnableDns64Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv4IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv4IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv4NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv4NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6Native() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Native",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6NativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6NativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Ipv6NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) MapPublicIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) MapPublicIpOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) NetworkAclAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAclAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) OutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) PrivateDnsNameOptionsOnLaunch() Ec2SubnetPrivateDnsNameOptionsOnLaunchOutputReference {
	var returns Ec2SubnetPrivateDnsNameOptionsOnLaunchOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptionsOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) PrivateDnsNameOptionsOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptionsOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) Tags() Ec2SubnetTagsList {
	var returns Ec2SubnetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Subnet) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet awscc_ec2_subnet} Resource.
func NewEc2Subnet(scope constructs.Construct, id *string, config *Ec2SubnetConfig) Ec2Subnet {
	_init_.Initialize()

	if err := validateNewEc2SubnetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Subnet{}

	_jsii_.Create(
		"awscc.ec2Subnet.Ec2Subnet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet awscc_ec2_subnet} Resource.
func NewEc2Subnet_Override(e Ec2Subnet, scope constructs.Construct, id *string, config *Ec2SubnetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Subnet.Ec2Subnet",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetAssignIpv6AddressOnCreation(val interface{}) {
	if err := j.validateSetAssignIpv6AddressOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetEnableDns64(val interface{}) {
	if err := j.validateSetEnableDns64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns64",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv4IpamPoolId(val *string) {
	if err := j.validateSetIpv4IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv4NetmaskLength(val *float64) {
	if err := j.validateSetIpv4NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv6IpamPoolId(val *string) {
	if err := j.validateSetIpv6IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv6Native(val interface{}) {
	if err := j.validateSetIpv6NativeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Native",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetIpv6NetmaskLength(val *float64) {
	if err := j.validateSetIpv6NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetMapPublicIpOnLaunch(val interface{}) {
	if err := j.validateSetMapPublicIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetOutpostArn(val *string) {
	if err := j.validateSetOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostArn",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2Subnet)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a Ec2Subnet resource upon running "cdktf plan <stack-name>".
func Ec2Subnet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2Subnet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2Subnet.Ec2Subnet",
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
func Ec2Subnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Subnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Subnet.Ec2Subnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Subnet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Subnet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Subnet.Ec2Subnet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Subnet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Subnet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Subnet.Ec2Subnet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2Subnet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2Subnet.Ec2Subnet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2Subnet) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2Subnet) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2Subnet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Subnet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Subnet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Subnet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Subnet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Subnet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Subnet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Subnet) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Subnet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Subnet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2Subnet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Subnet) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Subnet) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2Subnet) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Subnet) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2Subnet) PutPrivateDnsNameOptionsOnLaunch(value *Ec2SubnetPrivateDnsNameOptionsOnLaunch) {
	if err := e.validatePutPrivateDnsNameOptionsOnLaunchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateDnsNameOptionsOnLaunch",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Subnet) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetAssignIpv6AddressOnCreation() {
	_jsii_.InvokeVoid(
		e,
		"resetAssignIpv6AddressOnCreation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetAvailabilityZoneId() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZoneId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		e,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetEnableDns64() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableDns64",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv4IpamPoolId() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4IpamPoolId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv4NetmaskLength() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4NetmaskLength",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv6IpamPoolId() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6IpamPoolId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv6Native() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Native",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetIpv6NetmaskLength() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6NetmaskLength",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetMapPublicIpOnLaunch() {
	_jsii_.InvokeVoid(
		e,
		"resetMapPublicIpOnLaunch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetOutpostArn() {
	_jsii_.InvokeVoid(
		e,
		"resetOutpostArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetPrivateDnsNameOptionsOnLaunch() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsNameOptionsOnLaunch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Subnet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Subnet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

