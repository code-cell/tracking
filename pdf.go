package main

import (
	"bytes"
	"log"
	"text/template"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gobuffalo/packr"
)

//go:generate packr

type templateRow struct {
	Description string
	Hours       float32
	Rate        string
	Total       string
}

type templateData struct {
	SubmitDate    string
	CompanyName   string
	From          string
	For           string
	InvoiceNumber string
	Due           string
	VatRate       float32
	Vat           string
	Total         string
	Rows          []*templateRow
}

func generateInvoice(output string, from *Client, invoice *Invoice, client *Client, hours []*Hour) {
	var sumHours float32
	for _, hour := range hours {
		sumHours += hour.Hours
	}

	vat := float32(0.0)

	data := &templateData{
		SubmitDate:    formatTime(invoice.To),
		CompanyName:   from.Key,
		From:          linesWithBR(from.BillingInfo),
		For:           linesWithBR(client.BillingInfo),
		InvoiceNumber: invoice.Number,
		Due:           formatTime(invoice.To),
		VatRate:       0,
		Vat:           formatCurrency(vat),
		Total:         formatCurrency((sumHours * invoice.Rate) + vat),
		Rows: []*templateRow{
			{
				Description: client.Name,
				Hours:       sumHours,
				Rate:        formatCurrency(invoice.Rate),
				Total:       formatCurrency(sumHours * invoice.Rate),
			},
		},
	}

	box := packr.NewBox("./templates")

	t := template.New("invoice")
	t, err := t.Parse(box.String("simple.html"))
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(40)
	pdfg.MarginLeft.Set(40)
	pdfg.MarginRight.Set(40)

	page := wkhtmltopdf.NewPageReader(&buf)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(output)
	if err != nil {
		log.Fatal(err)
	}
}
