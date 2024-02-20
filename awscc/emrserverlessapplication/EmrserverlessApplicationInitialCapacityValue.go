package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#worker_configuration EmrserverlessApplication#worker_configuration}.
	WorkerConfiguration *EmrserverlessApplicationInitialCapacityValueWorkerConfiguration `field:"required" json:"workerConfiguration" yaml:"workerConfiguration"`
	// Initial count of workers to be initialized when an Application is started.
	//
	// This count will be continued to be maintained until the Application is stopped
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#worker_count EmrserverlessApplication#worker_count}
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
}

