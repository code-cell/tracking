package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/code-cell/tracking/data"
)

var (
	clientsFile  string
	invoicesFile string
	hoursPattern string
	fromFile     string
	invoiceNum   string
	output       string
)

func main() {
	flag.StringVar(&clientsFile, "clients", "clients.yaml", "YAML file containing the list of clients")
	flag.StringVar(&invoicesFile, "invoices", "invoices.yaml", "YAML file containing the list of invoices")
	flag.StringVar(&hoursPattern, "hours", "hours*.yaml", "Pattern to find YAML files containing the list of hours")
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

	matches, err := filepath.Glob(hoursPattern)
	if err != nil {
		log.Fatal(err)
	}
	hours := []*data.Hour{}
	for _, match := range matches {
		hoursRaw, err := ioutil.ReadFile(match)
		if err != nil {
			log.Fatal(err)
		}
		hours = append(hours, data.ParseHours(string(hoursRaw))...)
	}

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
