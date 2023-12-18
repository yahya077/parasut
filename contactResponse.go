package parasut

type ContactResponse struct {
	Data []Contact `json:"data"`
	HasResponseMeta
	HasLinks
}

type Contact struct {
	Data
	HasAttributes
	HasRelationships
}

type ContactAttributes struct {
	Balance     float64 `json:"balance"`
	TrlBalance  float64 `json:"trl_balance"`
	UsdBalance  float64 `json:"usd_balance"`
	EurBalance  float64 `json:"eur_balance"`
	GbpBalance  float64 `json:"gbp_balance"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	ShortName   string  `json:"short_name"`
	ContactType string  `json:"contact_type"`
	TaxOffice   string  `json:"tax_office"`
	TaxNumber   string  `json:"tax_number"`
	District    string  `json:"district"`
	PostalCode  string  `json:"postal_code"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	Fax         string  `json:"fax"`
	IsAbroad    bool    `json:"is_abroad"`
	Archived    bool    `json:"archived"`
	Iban        string  `json:"iban"`
	AccountType string  `json:"account_type"`
	Untrackable bool    `json:"untrackable"`
}
