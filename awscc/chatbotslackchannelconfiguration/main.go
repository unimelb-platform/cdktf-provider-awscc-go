package chatbotslackchannelconfiguration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfiguration",
		reflect.TypeOf((*ChatbotSlackChannelConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configurationName", GoGetter: "ConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "configurationNameInput", GoGetter: "ConfigurationNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailPolicies", GoGetter: "GuardrailPolicies"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailPoliciesInput", GoGetter: "GuardrailPoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleArn", GoGetter: "IamRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleArnInput", GoGetter: "IamRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevel", GoGetter: "LoggingLevel"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevelInput", GoGetter: "LoggingLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetGuardrailPolicies", GoMethod: "ResetGuardrailPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingLevel", GoMethod: "ResetLoggingLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnsTopicArns", GoMethod: "ResetSnsTopicArns"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserRoleRequired", GoMethod: "ResetUserRoleRequired"},
			_jsii_.MemberProperty{JsiiProperty: "slackChannelId", GoGetter: "SlackChannelId"},
			_jsii_.MemberProperty{JsiiProperty: "slackChannelIdInput", GoGetter: "SlackChannelIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "slackWorkspaceId", GoGetter: "SlackWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "slackWorkspaceIdInput", GoGetter: "SlackWorkspaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicArns", GoGetter: "SnsTopicArns"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicArnsInput", GoGetter: "SnsTopicArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userRoleRequired", GoGetter: "UserRoleRequired"},
			_jsii_.MemberProperty{JsiiProperty: "userRoleRequiredInput", GoGetter: "UserRoleRequiredInput"},
		},
		func() interface{} {
			j := jsiiProxy_ChatbotSlackChannelConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.chatbotSlackChannelConfiguration.ChatbotSlackChannelConfigurationConfig",
		reflect.TypeOf((*ChatbotSlackChannelConfigurationConfig)(nil)).Elem(),
	)
}
