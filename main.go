package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var (
	clientsFile  string
	invoicesFile string
	hoursFile    string
	fromCompany  string
	invoiceNum   string
	output       string
)

func main() {
	flag.StringVar(&clientsFile, "clients", "clients.md", "Markdown file containing the list of clients")
	flag.StringVar(&invoicesFile, "invoices", "invoices.md", "Markdown file containing the list of invoices")
	flag.StringVar(&hoursFile, "hours", "hours.md", "Markdown file containing the list of hours")
	flag.StringVar(&fromCompany, "from", "", "Company generating the invoice")
	flag.StringVar(&invoiceNum, "i", "", "Invoice to generate")
	flag.StringVar(&output, "o", "", "Output file")
	flag.Parse()

	if invoiceNum == "" {
		log.Fatal("Please, specify an invoice to generate")
	}

	if output == "" {
		log.Fatal("Please, specify an output file")
	}

	if fromCompany == "" {
		log.Fatal("Please, specify a from company")
	}

	clientsRaw, err := ioutil.ReadFile(clientsFile)
	if err != nil {
		log.Fatal(err)
	}
	clients := ParseClients(string(clientsRaw))

	invoicesRaw, err := ioutil.ReadFile(invoicesFile)
	if err != nil {
		log.Fatal(err)
	}
	invoices := ParseInvoices(string(invoicesRaw))

	hoursRaw, err := ioutil.ReadFile(hoursFile)
	if err != nil {
		log.Fatal(err)
	}
	hours := ParseHours(string(hoursRaw))

	var invoice *Invoice
	for _, i := range invoices {
		if i.Number == invoiceNum {
			invoice = i
			break
		}
	}
	if invoice == nil {
		log.Fatal("Invoice not found")
	}

	var client *Client
	for _, c := range clients {
		if c.Key == invoice.Client {
			client = c
			break
		}
	}
	if client == nil {
		log.Fatal("Client not found")
	}

	var from *Client
	for _, c := range clients {
		if c.Key == fromCompany {
			from = c
			break
		}
	}
	if client == nil {
		log.Fatal("Client not found")
	}

	invoiceHours := make([]*Hour, 0)
	for _, h := range hours {
		if invoice.Contains(h) {
			invoiceHours = append(invoiceHours, h)
		}
	}

	generateInvoice(output, from, invoice, client, invoiceHours)
}
