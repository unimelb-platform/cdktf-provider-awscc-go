package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint awscc_mediapackage_origin_endpoint}.
type MediapackageOriginEndpoint interface {
	cdktf.TerraformResource
	Arn() *string
	Authorization() MediapackageOriginEndpointAuthorizationOutputReference
	AuthorizationInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelId() *string
	SetChannelId(val *string)
	ChannelIdInput() *string
	CmafPackage() MediapackageOriginEndpointCmafPackageOutputReference
	CmafPackageInput() interface{}
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
	DashPackage() MediapackageOriginEndpointDashPackageOutputReference
	DashPackageInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HlsPackage() MediapackageOriginEndpointHlsPackageOutputReference
	HlsPackageInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManifestName() *string
	SetManifestName(val *string)
	ManifestNameInput() *string
	MssPackage() MediapackageOriginEndpointMssPackageOutputReference
	MssPackageInput() interface{}
	// The tree node.
	Node() constructs.Node
	Origination() *string
	SetOrigination(val *string)
	OriginationInput() *string
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
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	StartoverWindowSecondsInput() *float64
	Tags() MediapackageOriginEndpointTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	TimeDelaySecondsInput() *float64
	Url() *string
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	WhitelistInput() *[]*string
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
	PutAuthorization(value *MediapackageOriginEndpointAuthorization)
	PutCmafPackage(value *MediapackageOriginEndpointCmafPackage)
	PutDashPackage(value *MediapackageOriginEndpointDashPackage)
	PutHlsPackage(value *MediapackageOriginEndpointHlsPackage)
	PutMssPackage(value *MediapackageOriginEndpointMssPackage)
	PutTags(value interface{})
	ResetAuthorization()
	ResetCmafPackage()
	ResetDashPackage()
	ResetDescription()
	ResetHlsPackage()
	ResetManifestName()
	ResetMssPackage()
	ResetOrigination()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartoverWindowSeconds()
	ResetTags()
	ResetTimeDelaySeconds()
	ResetWhitelist()
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

// The jsii proxy struct for MediapackageOriginEndpoint
type jsiiProxy_MediapackageOriginEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Authorization() MediapackageOriginEndpointAuthorizationOutputReference {
	var returns MediapackageOriginEndpointAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) AuthorizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) CmafPackage() MediapackageOriginEndpointCmafPackageOutputReference {
	var returns MediapackageOriginEndpointCmafPackageOutputReference
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) CmafPackageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) DashPackage() MediapackageOriginEndpointDashPackageOutputReference {
	var returns MediapackageOriginEndpointDashPackageOutputReference
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) DashPackageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) HlsPackage() MediapackageOriginEndpointHlsPackageOutputReference {
	var returns MediapackageOriginEndpointHlsPackageOutputReference
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) HlsPackageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) ManifestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) MssPackage() MediapackageOriginEndpointMssPackageOutputReference {
	var returns MediapackageOriginEndpointMssPackageOutputReference
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) MssPackageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Origination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) OriginationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) StartoverWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) StartoverWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Tags() MediapackageOriginEndpointTagsList {
	var returns MediapackageOriginEndpointTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) TimeDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpoint) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint awscc_mediapackage_origin_endpoint} Resource.
func NewMediapackageOriginEndpoint(scope constructs.Construct, id *string, config *MediapackageOriginEndpointConfig) MediapackageOriginEndpoint {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpoint{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint awscc_mediapackage_origin_endpoint} Resource.
func NewMediapackageOriginEndpoint_Override(m MediapackageOriginEndpoint, scope constructs.Construct, id *string, config *MediapackageOriginEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetManifestName(val *string) {
	if err := j.validateSetManifestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetOrigination(val *string) {
	if err := j.validateSetOriginationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"origination",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetStartoverWindowSeconds(val *float64) {
	if err := j.validateSetStartoverWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startoverWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetTimeDelaySeconds(val *float64) {
	if err := j.validateSetTimeDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpoint)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
}

// Generates CDKTF code for importing a MediapackageOriginEndpoint resource upon running "cdktf plan <stack-name>".
func MediapackageOriginEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediapackageOriginEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
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
func MediapackageOriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackageOriginEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediapackageOriginEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackageOriginEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MediapackageOriginEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackageOriginEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediapackageOriginEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutAuthorization(value *MediapackageOriginEndpointAuthorization) {
	if err := m.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutCmafPackage(value *MediapackageOriginEndpointCmafPackage) {
	if err := m.validatePutCmafPackageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCmafPackage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutDashPackage(value *MediapackageOriginEndpointDashPackage) {
	if err := m.validatePutDashPackageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDashPackage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutHlsPackage(value *MediapackageOriginEndpointHlsPackage) {
	if err := m.validatePutHlsPackageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHlsPackage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutMssPackage(value *MediapackageOriginEndpointMssPackage) {
	if err := m.validatePutMssPackageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMssPackage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetAuthorization() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetCmafPackage() {
	_jsii_.InvokeVoid(
		m,
		"resetCmafPackage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetDashPackage() {
	_jsii_.InvokeVoid(
		m,
		"resetDashPackage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetHlsPackage() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsPackage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetManifestName() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetMssPackage() {
	_jsii_.InvokeVoid(
		m,
		"resetMssPackage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetOrigination() {
	_jsii_.InvokeVoid(
		m,
		"resetOrigination",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetStartoverWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetStartoverWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetTimeDelaySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeDelaySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ResetWhitelist() {
	_jsii_.InvokeVoid(
		m,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

