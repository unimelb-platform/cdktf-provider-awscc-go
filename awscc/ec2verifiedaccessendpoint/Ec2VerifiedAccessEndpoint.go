package ec2verifiedaccessendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2verifiedaccessendpoint/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint awscc_ec2_verified_access_endpoint}.
type Ec2VerifiedAccessEndpoint interface {
	cdktf.TerraformResource
	ApplicationDomain() *string
	SetApplicationDomain(val *string)
	ApplicationDomainInput() *string
	AttachmentType() *string
	SetAttachmentType(val *string)
	AttachmentTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceValidationDomain() *string
	DomainCertificateArn() *string
	SetDomainCertificateArn(val *string)
	DomainCertificateArnInput() *string
	EndpointDomain() *string
	EndpointDomainPrefix() *string
	SetEndpointDomainPrefix(val *string)
	EndpointDomainPrefixInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastUpdatedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerOptions() Ec2VerifiedAccessEndpointLoadBalancerOptionsOutputReference
	LoadBalancerOptionsInput() interface{}
	NetworkInterfaceOptions() Ec2VerifiedAccessEndpointNetworkInterfaceOptionsOutputReference
	NetworkInterfaceOptionsInput() interface{}
	// The tree node.
	Node() constructs.Node
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	PolicyDocumentInput() *string
	PolicyEnabled() interface{}
	SetPolicyEnabled(val interface{})
	PolicyEnabledInput() interface{}
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SseSpecification() Ec2VerifiedAccessEndpointSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	Status() *string
	Tags() Ec2VerifiedAccessEndpointTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VerifiedAccessEndpointId() *string
	VerifiedAccessGroupId() *string
	SetVerifiedAccessGroupId(val *string)
	VerifiedAccessGroupIdInput() *string
	VerifiedAccessInstanceId() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLoadBalancerOptions(value *Ec2VerifiedAccessEndpointLoadBalancerOptions)
	PutNetworkInterfaceOptions(value *Ec2VerifiedAccessEndpointNetworkInterfaceOptions)
	PutSseSpecification(value *Ec2VerifiedAccessEndpointSseSpecification)
	PutTags(value interface{})
	ResetDescription()
	ResetLoadBalancerOptions()
	ResetNetworkInterfaceOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDocument()
	ResetPolicyEnabled()
	ResetSecurityGroupIds()
	ResetSseSpecification()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2VerifiedAccessEndpoint
type jsiiProxy_Ec2VerifiedAccessEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) ApplicationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) ApplicationDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) AttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) DeviceValidationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceValidationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) DomainCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) DomainCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) EndpointDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) EndpointDomainPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) EndpointDomainPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) LoadBalancerOptions() Ec2VerifiedAccessEndpointLoadBalancerOptionsOutputReference {
	var returns Ec2VerifiedAccessEndpointLoadBalancerOptionsOutputReference
	_jsii_.Get(
		j,
		"loadBalancerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) LoadBalancerOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) NetworkInterfaceOptions() Ec2VerifiedAccessEndpointNetworkInterfaceOptionsOutputReference {
	var returns Ec2VerifiedAccessEndpointNetworkInterfaceOptionsOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) NetworkInterfaceOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) PolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) PolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) SseSpecification() Ec2VerifiedAccessEndpointSseSpecificationOutputReference {
	var returns Ec2VerifiedAccessEndpointSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) Tags() Ec2VerifiedAccessEndpointTagsList {
	var returns Ec2VerifiedAccessEndpointTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) VerifiedAccessEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) VerifiedAccessGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) VerifiedAccessGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint) VerifiedAccessInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessInstanceId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint awscc_ec2_verified_access_endpoint} Resource.
func NewEc2VerifiedAccessEndpoint(scope constructs.Construct, id *string, config *Ec2VerifiedAccessEndpointConfig) Ec2VerifiedAccessEndpoint {
	_init_.Initialize()

	if err := validateNewEc2VerifiedAccessEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VerifiedAccessEndpoint{}

	_jsii_.Create(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint awscc_ec2_verified_access_endpoint} Resource.
func NewEc2VerifiedAccessEndpoint_Override(e Ec2VerifiedAccessEndpoint, scope constructs.Construct, id *string, config *Ec2VerifiedAccessEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetApplicationDomain(val *string) {
	if err := j.validateSetApplicationDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationDomain",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetDomainCertificateArn(val *string) {
	if err := j.validateSetDomainCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainCertificateArn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetEndpointDomainPrefix(val *string) {
	if err := j.validateSetEndpointDomainPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointDomainPrefix",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetPolicyEnabled(val interface{}) {
	if err := j.validateSetPolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyEnabled",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessEndpoint)SetVerifiedAccessGroupId(val *string) {
	if err := j.validateSetVerifiedAccessGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAccessGroupId",
		val,
	)
}

// Generates CDKTF code for importing a Ec2VerifiedAccessEndpoint resource upon running "cdktf plan <stack-name>".
func Ec2VerifiedAccessEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
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
func Ec2VerifiedAccessEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VerifiedAccessEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VerifiedAccessEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2VerifiedAccessEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2VerifiedAccessEndpoint.Ec2VerifiedAccessEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) PutLoadBalancerOptions(value *Ec2VerifiedAccessEndpointLoadBalancerOptions) {
	if err := e.validatePutLoadBalancerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancerOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) PutNetworkInterfaceOptions(value *Ec2VerifiedAccessEndpointNetworkInterfaceOptions) {
	if err := e.validatePutNetworkInterfaceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaceOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) PutSseSpecification(value *Ec2VerifiedAccessEndpointSseSpecification) {
	if err := e.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetLoadBalancerOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancerOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetNetworkInterfaceOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaceOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetPolicyDocument() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicyDocument",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetPolicyEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicyEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

