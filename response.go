package parasut

import "time"

type Data struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type HasRelationships struct {
	Relationships struct {
		HasCategoryRelation
		HasContactPortalRelation
		HasContactPeopleRelation
		HasCompanyRelation
		HasTagsRelation
	} `json:"relationships"`
}

type HasRelationShip struct {
	Data Data `json:"data,omitempty"`
}

type HasMultipleRelationShip struct {
	Data []Data `json:"data,omitempty"`
}

type HasCategoryRelation struct {
	Category HasRelationShip `json:"category,omitempty"`
}

type HasContactPortalRelation struct {
	ContactPortal HasRelationShip `json:"contact_portal,omitempty"`
}

type HasCompanyRelation struct {
	Company HasRelationShip `json:"company,omitempty"`
}

type HasContactPeopleRelation struct {
	ContactPeople HasMultipleRelationShip `json:"contact_people,omitempty"`
}

type HasTagsRelation struct {
	Tags HasMultipleRelationShip `json:"tags,omitempty"`
}

type HasListMeta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
}

type HasAttributes struct {
	Attributes struct {
		CreatedAt    time.Time     `json:"created_at"`
		UpdatedAt    time.Time     `json:"updated_at"`
		CategoryType string        `json:"category_type"`
		Name         string        `json:"name"`
		FullPath     []interface{} `json:"full_path"`
		BgColor      string        `json:"bg_color"`
		TextColor    string        `json:"text_color"`
	} `json:"attributes"`
}

type HasLinks struct {
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

type HasResponseMeta struct {
	Meta struct {
		CurrentPage                          int    `json:"current_page"`
		TotalPages                           int    `json:"total_pages"`
		TotalCount                           int    `json:"total_count"`
		PerPage                              int    `json:"per_page"`
		PayableTotal                         string `json:"payable_total"`
		CollectibleTotal                     string `json:"collectible_total"`
		SupplierPaymentsDueInThirtyDaysTotal string `json:"supplier_payments_due_in_thirty_days_total"`
		CustomerPaymentsDueInThirtyDaysTotal string `json:"customer_payments_due_in_thirty_days_total"`
		ExportUrl                            string `json:"export_url"`
	} `json:"meta"`
}
