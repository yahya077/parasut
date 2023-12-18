package parasut

import (
	"fmt"
	"github.com/WEG-Technology/room"
	"os"
)

type ContactIndexParams struct {
	Name        string `url:"filter[name],omitempty"`
	Email       string `url:"filter[email],omitempty"`
	TaxNumber   string `url:"filter[tax_number],omitempty"`
	TaxOffice   string `url:"filter[tax_office],omitempty"`
	City        string `url:"filter[city],omitempty"`
	AccountType string `url:"filter[account_type],omitempty"`
	HasPagination
	HasInclude
}

type ContactPostParams struct {
	HasInclude
}

func ContactIndex(params ContactIndexParams) room.IRequest {
	r, e := room.NewGetRequest(
		room.WithEndPoint(fmt.Sprintf("%s/%s/%s", v4Endpoint, os.Getenv(CompanyIDEnv), contactEndpoint)),
		room.WithQuery(room.NewQuery(params)),
		room.WithDto(&ContactResponse{}),
	)

	if e != nil {
		panic(e)
	}

	return r
}

func ContactCreateBasicPurchaseBill(params ContactPostParams) room.IRequest {
	r, e := room.NewGetRequest(
		room.WithEndPoint(fmt.Sprintf("%s/%s/%s", v4Endpoint, os.Getenv(CompanyIDEnv), contactCreateBasicPurchaseBillEndpoint)),
		room.WithBody(params),
		room.WithDto(&ContactResponse{}),
	)

	if e != nil {
		panic(e)
	}

	return r
}
