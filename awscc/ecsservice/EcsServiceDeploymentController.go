package ecsservice


type EcsServiceDeploymentController struct {
	// The deployment controller type to use.
	//
	// There are three deployment controller types available:
	//   + ECS The rolling update (ECS) deployment type involves replacing the current running version of the container with the latest version. The number of containers Amazon ECS adds or removes from the service during a rolling update is controlled by adjusting the minimum and maximum number of healthy tasks allowed during a service deployment, as specified in the DeploymentConfiguration. + CODE_DEPLOY The blue/green (CODE_DEPLOY) deployment type uses the blue/green deployment model powered by , which allows you to verify a new deployment of a service before sending production traffic to it. + EXTERNAL The external (EXTERNAL) deployment type enables you to use any third-party deployment controller for full control over the deployment process for an Amazon ECS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_service#type EcsService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

