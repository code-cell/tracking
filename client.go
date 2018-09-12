package main

import blackfriday "gopkg.in/russross/blackfriday.v2"

type Client struct {
	Key         string
	BillingInfo string
}

func ParseClients(markdown string) []*Client {
	parser := blackfriday.New()
	ast := parser.Parse([]byte(markdown))

	clients := make([]*Client, 0)
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering && node.Type == blackfriday.Heading {
			key := node.FirstChild.Literal
			billingInfo := node.Next.FirstChild.Literal
			clients = append(clients, &Client{
				Key:         string(key),
				BillingInfo: string(billingInfo),
			})
		}
		return blackfriday.GoToNext
	})

	return clients
}
