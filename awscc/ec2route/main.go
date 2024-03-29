package ec2route

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.ec2Route.Ec2Route",
		reflect.TypeOf((*Ec2Route)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "carrierGatewayId", GoGetter: "CarrierGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "carrierGatewayIdInput", GoGetter: "CarrierGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "coreNetworkArn", GoGetter: "CoreNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "coreNetworkArnInput", GoGetter: "CoreNetworkArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlockInput", GoGetter: "DestinationCidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationIpv6CidrBlock", GoGetter: "DestinationIpv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationIpv6CidrBlockInput", GoGetter: "DestinationIpv6CidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPrefixListId", GoGetter: "DestinationPrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPrefixListIdInput", GoGetter: "DestinationPrefixListIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressOnlyInternetGatewayId", GoGetter: "EgressOnlyInternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "egressOnlyInternetGatewayIdInput", GoGetter: "EgressOnlyInternetGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayIdInput", GoGetter: "GatewayIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdInput", GoGetter: "InstanceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayId", GoGetter: "LocalGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayIdInput", GoGetter: "LocalGatewayIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayId", GoGetter: "NatGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayIdInput", GoGetter: "NatGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceIdInput", GoGetter: "NetworkInterfaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCarrierGatewayId", GoMethod: "ResetCarrierGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoreNetworkArn", GoMethod: "ResetCoreNetworkArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationCidrBlock", GoMethod: "ResetDestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationIpv6CidrBlock", GoMethod: "ResetDestinationIpv6CidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationPrefixListId", GoMethod: "ResetDestinationPrefixListId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressOnlyInternetGatewayId", GoMethod: "ResetEgressOnlyInternetGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGatewayId", GoMethod: "ResetGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceId", GoMethod: "ResetInstanceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalGatewayId", GoMethod: "ResetLocalGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNatGatewayId", GoMethod: "ResetNatGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterfaceId", GoMethod: "ResetNetworkInterfaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransitGatewayId", GoMethod: "ResetTransitGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcEndpointId", GoMethod: "ResetVpcEndpointId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcPeeringConnectionId", GoMethod: "ResetVpcPeeringConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableIdInput", GoGetter: "RouteTableIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayIdInput", GoGetter: "TransitGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointIdInput", GoGetter: "VpcEndpointIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPeeringConnectionId", GoGetter: "VpcPeeringConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPeeringConnectionIdInput", GoGetter: "VpcPeeringConnectionIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2Route{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.ec2Route.Ec2RouteConfig",
		reflect.TypeOf((*Ec2RouteConfig)(nil)).Elem(),
	)
}
