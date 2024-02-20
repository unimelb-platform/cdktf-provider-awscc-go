package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitions struct {
	// The image used to start a container. This string is passed directly to the Docker daemon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#image EcsTaskDefinition#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// The name of a container. Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#command EcsTaskDefinition#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#cpu EcsTaskDefinition#cpu}.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#credential_specs EcsTaskDefinition#credential_specs}.
	CredentialSpecs *[]*string `field:"optional" json:"credentialSpecs" yaml:"credentialSpecs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#depends_on EcsTaskDefinition#depends_on}.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#disable_networking EcsTaskDefinition#disable_networking}.
	DisableNetworking interface{} `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#dns_search_domains EcsTaskDefinition#dns_search_domains}.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#dns_servers EcsTaskDefinition#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#docker_labels EcsTaskDefinition#docker_labels}.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#docker_security_options EcsTaskDefinition#docker_security_options}.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#entry_point EcsTaskDefinition#entry_point}.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#environment EcsTaskDefinition#environment}
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The list of one or more files that contain the environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#environment_files EcsTaskDefinition#environment_files}
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#essential EcsTaskDefinition#essential}.
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#extra_hosts EcsTaskDefinition#extra_hosts}.
	ExtraHosts interface{} `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#firelens_configuration EcsTaskDefinition#firelens_configuration}.
	FirelensConfiguration *EcsTaskDefinitionContainerDefinitionsFirelensConfiguration `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// The health check command and associated configuration parameters for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#health_check EcsTaskDefinition#health_check}
	HealthCheck *EcsTaskDefinitionContainerDefinitionsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#hostname EcsTaskDefinition#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#interactive EcsTaskDefinition#interactive}.
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#links EcsTaskDefinition#links}.
	Links *[]*string `field:"optional" json:"links" yaml:"links"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#linux_parameters EcsTaskDefinition#linux_parameters}.
	LinuxParameters *EcsTaskDefinitionContainerDefinitionsLinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#log_configuration EcsTaskDefinition#log_configuration}.
	LogConfiguration *EcsTaskDefinitionContainerDefinitionsLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the memory specified here, the container is killed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#memory EcsTaskDefinition#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#memory_reservation EcsTaskDefinition#memory_reservation}.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#mount_points EcsTaskDefinition#mount_points}.
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Port mappings allow containers to access ports on the host container instance to send or receive traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#port_mappings EcsTaskDefinition#port_mappings}
	PortMappings interface{} `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#privileged EcsTaskDefinition#privileged}.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#pseudo_terminal EcsTaskDefinition#pseudo_terminal}.
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#readonly_root_filesystem EcsTaskDefinition#readonly_root_filesystem}.
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#repository_credentials EcsTaskDefinition#repository_credentials}.
	RepositoryCredentials *EcsTaskDefinitionContainerDefinitionsRepositoryCredentials `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#resource_requirements EcsTaskDefinition#resource_requirements}.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#secrets EcsTaskDefinition#secrets}.
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#start_timeout EcsTaskDefinition#start_timeout}.
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#stop_timeout EcsTaskDefinition#stop_timeout}.
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#system_controls EcsTaskDefinition#system_controls}.
	SystemControls interface{} `field:"optional" json:"systemControls" yaml:"systemControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#ulimits EcsTaskDefinition#ulimits}.
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#user EcsTaskDefinition#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#volumes_from EcsTaskDefinition#volumes_from}.
	VolumesFrom interface{} `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#working_directory EcsTaskDefinition#working_directory}.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

