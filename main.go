package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/code-cell/tracking/data"
)

var (
	clientsFile  string
	invoicesFile string
	hoursFile    string
	fromFile     string
	invoiceNum   string
	output       string
)

func main() {
	flag.StringVar(&clientsFile, "clients", "clients.yaml", "YAML file containing the list of clients")
	flag.StringVar(&invoicesFile, "invoices", "invoices.yaml", "YAML file containing the list of invoices")
	flag.StringVar(&hoursFile, "hours", "hours.yaml", "YAML file containing the list of hours")
	flag.StringVar(&fromFile, "from", "from.yaml", "YAML file containing the info about the company generating the invoice")
	flag.StringVar(&invoiceNum, "i", "", "Invoice number to generate")
	flag.StringVar(&output, "o", "", "Output file (default to <invoice_number>.pdf)")
	flag.Parse()

	if invoiceNum == "" {
		log.Fatal("Please, specify an invoice to generate")
	}

	if output == "" {
		output = fmt.Sprintf("%v.pdf", invoiceNum)
	}

	clientsRaw, err := ioutil.ReadFile(clientsFile)
	if err != nil {
		log.Fatal(err)
	}
	clients := data.ParseClients(string(clientsRaw))

	invoicesRaw, err := ioutil.ReadFile(invoicesFile)
	if err != nil {
		log.Fatal(err)
	}
	invoices := data.ParseInvoices(string(invoicesRaw))

	hoursRaw, err := ioutil.ReadFile(hoursFile)
	if err != nil {
		log.Fatal(err)
	}
	hours := data.ParseHours(string(hoursRaw))

	fromRaw, err := ioutil.ReadFile(fromFile)
	if err != nil {
		log.Fatal(err)
	}
	from := data.ParseFrom(string(fromRaw))

	var invoice *data.Invoice
	for _, i := range invoices {
		if i.Number == invoiceNum {
			invoice = i
			break
		}
	}
	if invoice == nil {
		log.Fatal("Invoice not found")
	}

	var client *data.Client
	for _, c := range clients {
		if c.Key == invoice.Client {
			client = c
			break
		}
	}
	if client == nil {
		log.Fatal("Client not found")
	}

	invoiceHours := make([]*data.Hour, 0)
	for _, h := range hours {
		if h.Client == client.Key && invoice.Contains(h) {
			invoiceHours = append(invoiceHours, h)
		}
	}

	generateInvoice(output, from, invoice, client, invoiceHours)
}
