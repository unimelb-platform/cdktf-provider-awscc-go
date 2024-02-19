package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lambdafunction/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function awscc_lambda_function}.
type LambdaFunction interface {
	cdktf.TerraformResource
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Code() LambdaFunctionCodeOutputReference
	CodeInput() interface{}
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	CodeSigningConfigArnInput() *string
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
	DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference
	DeadLetterConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Environment() LambdaFunctionEnvironmentOutputReference
	EnvironmentInput() interface{}
	EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference
	EphemeralStorageInput() interface{}
	FileSystemConfigs() LambdaFunctionFileSystemConfigsList
	FileSystemConfigsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	ImageConfig() LambdaFunctionImageConfigOutputReference
	ImageConfigInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Layers() *[]*string
	SetLayers(val *[]*string)
	LayersInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfig() LambdaFunctionLoggingConfigOutputReference
	LoggingConfigInput() interface{}
	MemorySize() *float64
	SetMemorySize(val *float64)
	MemorySizeInput() *float64
	// The tree node.
	Node() constructs.Node
	PackageType() *string
	SetPackageType(val *string)
	PackageTypeInput() *string
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
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	ReservedConcurrentExecutionsInput() *float64
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	RuntimeManagementConfig() LambdaFunctionRuntimeManagementConfigOutputReference
	RuntimeManagementConfigInput() interface{}
	SnapStart() LambdaFunctionSnapStartOutputReference
	SnapStartInput() interface{}
	SnapStartResponse() LambdaFunctionSnapStartResponseOutputReference
	Tags() LambdaFunctionTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	TracingConfig() LambdaFunctionTracingConfigOutputReference
	TracingConfigInput() interface{}
	VpcConfig() LambdaFunctionVpcConfigOutputReference
	VpcConfigInput() interface{}
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
	PutCode(value *LambdaFunctionCode)
	PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig)
	PutEnvironment(value *LambdaFunctionEnvironment)
	PutEphemeralStorage(value *LambdaFunctionEphemeralStorage)
	PutFileSystemConfigs(value interface{})
	PutImageConfig(value *LambdaFunctionImageConfig)
	PutLoggingConfig(value *LambdaFunctionLoggingConfig)
	PutRuntimeManagementConfig(value *LambdaFunctionRuntimeManagementConfig)
	PutSnapStart(value *LambdaFunctionSnapStart)
	PutTags(value interface{})
	PutTracingConfig(value *LambdaFunctionTracingConfig)
	PutVpcConfig(value *LambdaFunctionVpcConfig)
	ResetArchitectures()
	ResetCodeSigningConfigArn()
	ResetDeadLetterConfig()
	ResetDescription()
	ResetEnvironment()
	ResetEphemeralStorage()
	ResetFileSystemConfigs()
	ResetFunctionName()
	ResetHandler()
	ResetImageConfig()
	ResetKmsKeyArn()
	ResetLayers()
	ResetLoggingConfig()
	ResetMemorySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageType()
	ResetReservedConcurrentExecutions()
	ResetRuntime()
	ResetRuntimeManagementConfig()
	ResetSnapStart()
	ResetTags()
	ResetTimeout()
	ResetTracingConfig()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Code() LambdaFunctionCodeOutputReference {
	var returns LambdaFunctionCodeOutputReference
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference {
	var returns LambdaFunctionDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Environment() LambdaFunctionEnvironmentOutputReference {
	var returns LambdaFunctionEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference {
	var returns LambdaFunctionEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigs() LambdaFunctionFileSystemConfigsList {
	var returns LambdaFunctionFileSystemConfigsList
	_jsii_.Get(
		j,
		"fileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfig() LambdaFunctionImageConfigOutputReference {
	var returns LambdaFunctionImageConfigOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfig() LambdaFunctionLoggingConfigOutputReference {
	var returns LambdaFunctionLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeManagementConfig() LambdaFunctionRuntimeManagementConfigOutputReference {
	var returns LambdaFunctionRuntimeManagementConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeManagementConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeManagementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStart() LambdaFunctionSnapStartOutputReference {
	var returns LambdaFunctionSnapStartOutputReference
	_jsii_.Get(
		j,
		"snapStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStartResponse() LambdaFunctionSnapStartResponseOutputReference {
	var returns LambdaFunctionSnapStartResponseOutputReference
	_jsii_.Get(
		j,
		"snapStartResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Tags() LambdaFunctionTagsList {
	var returns LambdaFunctionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfig() LambdaFunctionTracingConfigOutputReference {
	var returns LambdaFunctionTracingConfigOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfig() LambdaFunctionVpcConfigOutputReference {
	var returns LambdaFunctionVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function awscc_lambda_function} Resource.
func NewLambdaFunction(scope constructs.Construct, id *string, config *LambdaFunctionConfig) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"awscc.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function awscc_lambda_function} Resource.
func NewLambdaFunction_Override(l LambdaFunction, scope constructs.Construct, id *string, config *LambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaFunction)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCodeSigningConfigArn(val *string) {
	if err := j.validateSetCodeSigningConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLayers(val *[]*string) {
	if err := j.validateSetLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetPackageType(val *string) {
	if err := j.validateSetPackageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetReservedConcurrentExecutions(val *float64) {
	if err := j.validateSetReservedConcurrentExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Generates CDKTF code for importing a LambdaFunction resource upon running "cdktf plan <stack-name>".
func LambdaFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLambdaFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lambdaFunction.LambdaFunction",
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
func LambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lambdaFunction.LambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lambdaFunction.LambdaFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lambdaFunction.LambdaFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lambdaFunction.LambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaFunction) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LambdaFunction) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LambdaFunction) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaFunction) PutCode(value *LambdaFunctionCode) {
	if err := l.validatePutCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCode",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig) {
	if err := l.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEnvironment(value *LambdaFunctionEnvironment) {
	if err := l.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEphemeralStorage(value *LambdaFunctionEphemeralStorage) {
	if err := l.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutFileSystemConfigs(value interface{}) {
	if err := l.validatePutFileSystemConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFileSystemConfigs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutImageConfig(value *LambdaFunctionImageConfig) {
	if err := l.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutLoggingConfig(value *LambdaFunctionLoggingConfig) {
	if err := l.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutRuntimeManagementConfig(value *LambdaFunctionRuntimeManagementConfig) {
	if err := l.validatePutRuntimeManagementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRuntimeManagementConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutSnapStart(value *LambdaFunctionSnapStart) {
	if err := l.validatePutSnapStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSnapStart",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTracingConfig(value *LambdaFunctionTracingConfig) {
	if err := l.validatePutTracingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutVpcConfig(value *LambdaFunctionVpcConfig) {
	if err := l.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		l,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		l,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		l,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFileSystemConfigs() {
	_jsii_.InvokeVoid(
		l,
		"resetFileSystemConfigs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFunctionName() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		l,
		"resetHandler",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		l,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		l,
		"resetLayers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		l,
		"resetMemorySize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		l,
		"resetPackageType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		l,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntimeManagementConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeManagementConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSnapStart() {
	_jsii_.InvokeVoid(
		l,
		"resetSnapStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

