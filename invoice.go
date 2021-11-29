package coinbase

type AInvoice struct {
	API *APIClient
}

type APIInvoiceResponse struct {
	Id string `json:"id,omitempty"`
	BusinessName string `json:"business_name,omitempty"`
	Resource string `json:"resource,omitempty"`
	Code string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	CustomerEmail string `json:"customer_email,omitempty"`
	LocalPrice APIInvoiceCurrency `json:"local_price,omitempty"`
	Memo string `json:"memo,omitempty"`
	HostedUrl string `json:"hosted_url,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type APIInvoiceCurrency struct {
	Amount string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}


type APIInvoice struct {
	father *AInvoice
	Data APIInvoiceResponse `json:"data,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

type APIInvoices struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Invoices []APIInvoice `json:"invoices,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

type APIInvoiceRequest struct {
	Pagination APIPagination `json:"pagination,omitempty"`
	Data []APIInvoiceResponse `json:"data,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

func (a *AInvoice) Get(id string) (invoice APIInvoice, err error) {
	err = a.API.Fetch("GET", "/invoices/"+id, nil, &invoice)
	invoice.father = a
	return
}

func (a *APIInvoice) Refresh() (err error) {
	err = a.father.API.Fetch("GET", "/invoices/"+a.Data.Id, nil, &a.Data)
	return
}

func (a *AInvoice) List() (invoices APIInvoices, err error) {
	temp := APIInvoiceRequest{}
	err = a.API.Fetch("GET", "/invoices/", nil, &temp)
	invoices.Pagination = temp.Pagination
	invoices.Errors = temp.Errors
	for _, data := range temp.Data {
		invoices.Invoices = append(invoices.Invoices, APIInvoice{father: a, Data: data, Errors: temp.Errors})
	}
	return
}

func (a *AInvoice) Create(data interface{}) (charge APIInvoice, err error) {
	err = a.API.Fetch("POST", "/invoices/", data, &charge)
	charge.father = a
	return
}

