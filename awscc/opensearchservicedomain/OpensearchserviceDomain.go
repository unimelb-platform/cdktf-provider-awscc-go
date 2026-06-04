package opensearchservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/opensearchservicedomain/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain awscc_opensearchservice_domain}.
type OpensearchserviceDomain interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	AdvancedOptions() *map[string]*string
	SetAdvancedOptions(val *map[string]*string)
	AdvancedOptionsInput() *map[string]*string
	AdvancedSecurityOptions() OpensearchserviceDomainAdvancedSecurityOptionsOutputReference
	AdvancedSecurityOptionsInput() interface{}
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() OpensearchserviceDomainClusterConfigOutputReference
	ClusterConfigInput() interface{}
	CognitoOptions() OpensearchserviceDomainCognitoOptionsOutputReference
	CognitoOptionsInput() interface{}
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
	DomainArn() *string
	DomainEndpoint() *string
	DomainEndpointOptions() OpensearchserviceDomainDomainEndpointOptionsOutputReference
	DomainEndpointOptionsInput() interface{}
	DomainEndpoints() cdktf.StringMap
	DomainEndpointV2() *string
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() OpensearchserviceDomainEbsOptionsOutputReference
	EbsOptionsInput() interface{}
	EncryptionAtRestOptions() OpensearchserviceDomainEncryptionAtRestOptionsOutputReference
	EncryptionAtRestOptionsInput() interface{}
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdentityCenterOptions() OpensearchserviceDomainIdentityCenterOptionsOutputReference
	IdentityCenterOptionsInput() interface{}
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() OpensearchserviceDomainLogPublishingOptionsMap
	LogPublishingOptionsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryptionOptions() OpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference
	NodeToNodeEncryptionOptionsInput() interface{}
	OffPeakWindowOptions() OpensearchserviceDomainOffPeakWindowOptionsOutputReference
	OffPeakWindowOptionsInput() interface{}
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
	ServiceSoftwareOptions() OpensearchserviceDomainServiceSoftwareOptionsOutputReference
	SkipShardMigrationWait() interface{}
	SetSkipShardMigrationWait(val interface{})
	SkipShardMigrationWaitInput() interface{}
	SnapshotOptions() OpensearchserviceDomainSnapshotOptionsOutputReference
	SnapshotOptionsInput() interface{}
	SoftwareUpdateOptions() OpensearchserviceDomainSoftwareUpdateOptionsOutputReference
	SoftwareUpdateOptionsInput() interface{}
	Tags() OpensearchserviceDomainTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcOptions() OpensearchserviceDomainVpcOptionsOutputReference
	VpcOptionsInput() interface{}
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
	PutAdvancedSecurityOptions(value *OpensearchserviceDomainAdvancedSecurityOptions)
	PutClusterConfig(value *OpensearchserviceDomainClusterConfig)
	PutCognitoOptions(value *OpensearchserviceDomainCognitoOptions)
	PutDomainEndpointOptions(value *OpensearchserviceDomainDomainEndpointOptions)
	PutEbsOptions(value *OpensearchserviceDomainEbsOptions)
	PutEncryptionAtRestOptions(value *OpensearchserviceDomainEncryptionAtRestOptions)
	PutIdentityCenterOptions(value *OpensearchserviceDomainIdentityCenterOptions)
	PutLogPublishingOptions(value interface{})
	PutNodeToNodeEncryptionOptions(value *OpensearchserviceDomainNodeToNodeEncryptionOptions)
	PutOffPeakWindowOptions(value *OpensearchserviceDomainOffPeakWindowOptions)
	PutSnapshotOptions(value *OpensearchserviceDomainSnapshotOptions)
	PutSoftwareUpdateOptions(value *OpensearchserviceDomainSoftwareUpdateOptions)
	PutTags(value interface{})
	PutVpcOptions(value *OpensearchserviceDomainVpcOptions)
	ResetAccessPolicies()
	ResetAdvancedOptions()
	ResetAdvancedSecurityOptions()
	ResetClusterConfig()
	ResetCognitoOptions()
	ResetDomainEndpointOptions()
	ResetDomainName()
	ResetEbsOptions()
	ResetEncryptionAtRestOptions()
	ResetEngineVersion()
	ResetIdentityCenterOptions()
	ResetIpAddressType()
	ResetLogPublishingOptions()
	ResetNodeToNodeEncryptionOptions()
	ResetOffPeakWindowOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkipShardMigrationWait()
	ResetSnapshotOptions()
	ResetSoftwareUpdateOptions()
	ResetTags()
	ResetVpcOptions()
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

// The jsii proxy struct for OpensearchserviceDomain
type jsiiProxy_OpensearchserviceDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpensearchserviceDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) AdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) AdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) AdvancedSecurityOptions() OpensearchserviceDomainAdvancedSecurityOptionsOutputReference {
	var returns OpensearchserviceDomainAdvancedSecurityOptionsOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) AdvancedSecurityOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) ClusterConfig() OpensearchserviceDomainClusterConfigOutputReference {
	var returns OpensearchserviceDomainClusterConfigOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) ClusterConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) CognitoOptions() OpensearchserviceDomainCognitoOptionsOutputReference {
	var returns OpensearchserviceDomainCognitoOptionsOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) CognitoOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainEndpointOptions() OpensearchserviceDomainDomainEndpointOptionsOutputReference {
	var returns OpensearchserviceDomainDomainEndpointOptionsOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainEndpointOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainEndpointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainEndpoints() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"domainEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainEndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EbsOptions() OpensearchserviceDomainEbsOptionsOutputReference {
	var returns OpensearchserviceDomainEbsOptionsOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EbsOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EncryptionAtRestOptions() OpensearchserviceDomainEncryptionAtRestOptionsOutputReference {
	var returns OpensearchserviceDomainEncryptionAtRestOptionsOutputReference
	_jsii_.Get(
		j,
		"encryptionAtRestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EncryptionAtRestOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) IdentityCenterOptions() OpensearchserviceDomainIdentityCenterOptionsOutputReference {
	var returns OpensearchserviceDomainIdentityCenterOptionsOutputReference
	_jsii_.Get(
		j,
		"identityCenterOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) IdentityCenterOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityCenterOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) LogPublishingOptions() OpensearchserviceDomainLogPublishingOptionsMap {
	var returns OpensearchserviceDomainLogPublishingOptionsMap
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) NodeToNodeEncryptionOptions() OpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference {
	var returns OpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) NodeToNodeEncryptionOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) OffPeakWindowOptions() OpensearchserviceDomainOffPeakWindowOptionsOutputReference {
	var returns OpensearchserviceDomainOffPeakWindowOptionsOutputReference
	_jsii_.Get(
		j,
		"offPeakWindowOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) OffPeakWindowOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"offPeakWindowOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) ServiceSoftwareOptions() OpensearchserviceDomainServiceSoftwareOptionsOutputReference {
	var returns OpensearchserviceDomainServiceSoftwareOptionsOutputReference
	_jsii_.Get(
		j,
		"serviceSoftwareOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SkipShardMigrationWait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipShardMigrationWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SkipShardMigrationWaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipShardMigrationWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SnapshotOptions() OpensearchserviceDomainSnapshotOptionsOutputReference {
	var returns OpensearchserviceDomainSnapshotOptionsOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SnapshotOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SoftwareUpdateOptions() OpensearchserviceDomainSoftwareUpdateOptionsOutputReference {
	var returns OpensearchserviceDomainSoftwareUpdateOptionsOutputReference
	_jsii_.Get(
		j,
		"softwareUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) SoftwareUpdateOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softwareUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) Tags() OpensearchserviceDomainTagsList {
	var returns OpensearchserviceDomainTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) VpcOptions() OpensearchserviceDomainVpcOptionsOutputReference {
	var returns OpensearchserviceDomainVpcOptionsOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomain) VpcOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain awscc_opensearchservice_domain} Resource.
func NewOpensearchserviceDomain(scope constructs.Construct, id *string, config *OpensearchserviceDomainConfig) OpensearchserviceDomain {
	_init_.Initialize()

	if err := validateNewOpensearchserviceDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchserviceDomain{}

	_jsii_.Create(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain awscc_opensearchservice_domain} Resource.
func NewOpensearchserviceDomain_Override(o OpensearchserviceDomain, scope constructs.Construct, id *string, config *OpensearchserviceDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetAccessPolicies(val *string) {
	if err := j.validateSetAccessPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomain)SetSkipShardMigrationWait(val interface{}) {
	if err := j.validateSetSkipShardMigrationWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipShardMigrationWait",
		val,
	)
}

// Generates CDKTF code for importing a OpensearchserviceDomain resource upon running "cdktf plan <stack-name>".
func OpensearchserviceDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOpensearchserviceDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
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
func OpensearchserviceDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchserviceDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpensearchserviceDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchserviceDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpensearchserviceDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpensearchserviceDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpensearchserviceDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.opensearchserviceDomain.OpensearchserviceDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutAdvancedSecurityOptions(value *OpensearchserviceDomainAdvancedSecurityOptions) {
	if err := o.validatePutAdvancedSecurityOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdvancedSecurityOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutClusterConfig(value *OpensearchserviceDomainClusterConfig) {
	if err := o.validatePutClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putClusterConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutCognitoOptions(value *OpensearchserviceDomainCognitoOptions) {
	if err := o.validatePutCognitoOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCognitoOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutDomainEndpointOptions(value *OpensearchserviceDomainDomainEndpointOptions) {
	if err := o.validatePutDomainEndpointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDomainEndpointOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutEbsOptions(value *OpensearchserviceDomainEbsOptions) {
	if err := o.validatePutEbsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEbsOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutEncryptionAtRestOptions(value *OpensearchserviceDomainEncryptionAtRestOptions) {
	if err := o.validatePutEncryptionAtRestOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEncryptionAtRestOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutIdentityCenterOptions(value *OpensearchserviceDomainIdentityCenterOptions) {
	if err := o.validatePutIdentityCenterOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIdentityCenterOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutLogPublishingOptions(value interface{}) {
	if err := o.validatePutLogPublishingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogPublishingOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutNodeToNodeEncryptionOptions(value *OpensearchserviceDomainNodeToNodeEncryptionOptions) {
	if err := o.validatePutNodeToNodeEncryptionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNodeToNodeEncryptionOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutOffPeakWindowOptions(value *OpensearchserviceDomainOffPeakWindowOptions) {
	if err := o.validatePutOffPeakWindowOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOffPeakWindowOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutSnapshotOptions(value *OpensearchserviceDomainSnapshotOptions) {
	if err := o.validatePutSnapshotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSnapshotOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutSoftwareUpdateOptions(value *OpensearchserviceDomainSoftwareUpdateOptions) {
	if err := o.validatePutSoftwareUpdateOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSoftwareUpdateOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) PutVpcOptions(value *OpensearchserviceDomainVpcOptions) {
	if err := o.validatePutVpcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetAccessPolicies() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessPolicies",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetAdvancedOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetAdvancedOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetAdvancedSecurityOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetAdvancedSecurityOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetClusterConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetCognitoOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetCognitoOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetDomainEndpointOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainEndpointOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetDomainName() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetEbsOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetEncryptionAtRestOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionAtRestOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetIdentityCenterOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetIdentityCenterOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		o,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetNodeToNodeEncryptionOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetNodeToNodeEncryptionOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetOffPeakWindowOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetOffPeakWindowOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetSkipShardMigrationWait() {
	_jsii_.InvokeVoid(
		o,
		"resetSkipShardMigrationWait",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetSnapshotOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetSnapshotOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetSoftwareUpdateOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetSoftwareUpdateOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

