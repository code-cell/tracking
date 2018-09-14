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
	sums := make(map[string]float32, 0)
	var sumTotal float32

	for _, hour := range hours {
		if _, found := sums[hour.Project]; !found {
			sums[hour.Project] = 0
		}
		sums[hour.Project] += hour.Hours
		sumTotal += hour.Hours
	}

	rows := make([]*templateRow, 0)

	for projectKey, sum := range sums {
		project := client.FindProject(projectKey)
		if project == nil {
			log.Fatalf("Client %v doesn't have project %v", client.Key, projectKey)
		}

		rows = append(rows, &templateRow{
			Description: project.Name,
			Hours:       sum,
			Rate:        formatCurrency(invoice.Rate),
			Total:       formatCurrency(sum * invoice.Rate),
		})
	}

	vat := float32(0.0)

	data := &templateData{
		SubmitDate:    formatTime(invoice.To),
		CompanyName:   from.Name,
		From:          linesWithBR(from.BillingInfo),
		For:           linesWithBR(client.BillingInfo),
		InvoiceNumber: invoice.Number,
		Due:           formatTime(invoice.To),
		VatRate:       0,
		Vat:           formatCurrency(vat),
		Total:         formatCurrency((sumTotal * invoice.Rate) + vat),
		Rows:          rows,
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
