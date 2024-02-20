package ec2verifiedaccesstrustprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2verifiedaccesstrustprovider/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider awscc_ec2_verified_access_trust_provider}.
type Ec2VerifiedAccessTrustProvider interface {
	cdktf.TerraformResource
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
	DeviceOptions() Ec2VerifiedAccessTrustProviderDeviceOptionsOutputReference
	DeviceOptionsInput() interface{}
	DeviceTrustProviderType() *string
	SetDeviceTrustProviderType(val *string)
	DeviceTrustProviderTypeInput() *string
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
	// The tree node.
	Node() constructs.Node
	OidcOptions() Ec2VerifiedAccessTrustProviderOidcOptionsOutputReference
	OidcOptionsInput() interface{}
	PolicyReferenceName() *string
	SetPolicyReferenceName(val *string)
	PolicyReferenceNameInput() *string
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
	SseSpecification() Ec2VerifiedAccessTrustProviderSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	Tags() Ec2VerifiedAccessTrustProviderTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustProviderType() *string
	SetTrustProviderType(val *string)
	TrustProviderTypeInput() *string
	UserTrustProviderType() *string
	SetUserTrustProviderType(val *string)
	UserTrustProviderTypeInput() *string
	VerifiedAccessTrustProviderId() *string
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
	PutDeviceOptions(value *Ec2VerifiedAccessTrustProviderDeviceOptions)
	PutOidcOptions(value *Ec2VerifiedAccessTrustProviderOidcOptions)
	PutSseSpecification(value *Ec2VerifiedAccessTrustProviderSseSpecification)
	PutTags(value interface{})
	ResetDescription()
	ResetDeviceOptions()
	ResetDeviceTrustProviderType()
	ResetOidcOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSseSpecification()
	ResetTags()
	ResetUserTrustProviderType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2VerifiedAccessTrustProvider
type jsiiProxy_Ec2VerifiedAccessTrustProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DeviceOptions() Ec2VerifiedAccessTrustProviderDeviceOptionsOutputReference {
	var returns Ec2VerifiedAccessTrustProviderDeviceOptionsOutputReference
	_jsii_.Get(
		j,
		"deviceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DeviceOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DeviceTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) DeviceTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) OidcOptions() Ec2VerifiedAccessTrustProviderOidcOptionsOutputReference {
	var returns Ec2VerifiedAccessTrustProviderOidcOptionsOutputReference
	_jsii_.Get(
		j,
		"oidcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) OidcOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidcOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) PolicyReferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyReferenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) PolicyReferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyReferenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) SseSpecification() Ec2VerifiedAccessTrustProviderSseSpecificationOutputReference {
	var returns Ec2VerifiedAccessTrustProviderSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) Tags() Ec2VerifiedAccessTrustProviderTagsList {
	var returns Ec2VerifiedAccessTrustProviderTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) TrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) UserTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) UserTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider) VerifiedAccessTrustProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessTrustProviderId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider awscc_ec2_verified_access_trust_provider} Resource.
func NewEc2VerifiedAccessTrustProvider(scope constructs.Construct, id *string, config *Ec2VerifiedAccessTrustProviderConfig) Ec2VerifiedAccessTrustProvider {
	_init_.Initialize()

	if err := validateNewEc2VerifiedAccessTrustProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VerifiedAccessTrustProvider{}

	_jsii_.Create(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider awscc_ec2_verified_access_trust_provider} Resource.
func NewEc2VerifiedAccessTrustProvider_Override(e Ec2VerifiedAccessTrustProvider, scope constructs.Construct, id *string, config *Ec2VerifiedAccessTrustProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetDeviceTrustProviderType(val *string) {
	if err := j.validateSetDeviceTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTrustProviderType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetPolicyReferenceName(val *string) {
	if err := j.validateSetPolicyReferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyReferenceName",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetTrustProviderType(val *string) {
	if err := j.validateSetTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustProviderType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessTrustProvider)SetUserTrustProviderType(val *string) {
	if err := j.validateSetUserTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTrustProviderType",
		val,
	)
}

// Generates CDKTF code for importing a Ec2VerifiedAccessTrustProvider resource upon running "cdktf plan <stack-name>".
func Ec2VerifiedAccessTrustProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessTrustProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
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
func Ec2VerifiedAccessTrustProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessTrustProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VerifiedAccessTrustProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessTrustProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2VerifiedAccessTrustProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2VerifiedAccessTrustProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2VerifiedAccessTrustProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2VerifiedAccessTrustProvider.Ec2VerifiedAccessTrustProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) PutDeviceOptions(value *Ec2VerifiedAccessTrustProviderDeviceOptions) {
	if err := e.validatePutDeviceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeviceOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) PutOidcOptions(value *Ec2VerifiedAccessTrustProviderOidcOptions) {
	if err := e.validatePutOidcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOidcOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) PutSseSpecification(value *Ec2VerifiedAccessTrustProviderSseSpecification) {
	if err := e.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetDeviceOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetDeviceTrustProviderType() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceTrustProviderType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetOidcOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetOidcOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ResetUserTrustProviderType() {
	_jsii_.InvokeVoid(
		e,
		"resetUserTrustProviderType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessTrustProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

