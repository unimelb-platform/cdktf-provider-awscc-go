package invoicinginvoiceunit


type InvoicingInvoiceUnitRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/invoicing_invoice_unit#linked_accounts InvoicingInvoiceUnit#linked_accounts}.
	LinkedAccounts *[]*string `field:"required" json:"linkedAccounts" yaml:"linkedAccounts"`
}

