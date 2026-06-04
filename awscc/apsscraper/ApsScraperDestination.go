package apsscraper


type ApsScraperDestination struct {
	// Configuration for Amazon Managed Prometheus metrics destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#amp_configuration ApsScraper#amp_configuration}
	AmpConfiguration *ApsScraperDestinationAmpConfiguration `field:"optional" json:"ampConfiguration" yaml:"ampConfiguration"`
}

