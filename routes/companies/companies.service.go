package companies

import (
	"github.com/rs/xid"
)

var companies = []Company{
	{ID: "1", Name: "Apple", CEO: "Tim Cook", Revenue: "USD 274.5 billion"},
	{ID: "2", Name: "Samsung", CEO: "Kim Hyun-suk", Revenue: "USD 197.7 billion"},
	{ID: "3", Name: "Microsoft", CEO: "Satya Nadella", Revenue: "USD 143 billion"},
}

func New(newCompany Company) {
	newCompany.ID = xid.New().String()
	companies = append(companies, newCompany)
}

func FindById(id string) *Company {
	for _, company := range companies {
		if company.ID == id {
			return &company
		}
	}
	return nil
}

func UpdateById(id string, updatedCompany Company) (*Company, bool) {
	for i, company := range companies {
		if company.ID == id {
			companies[i].Name = updatedCompany.Name
			companies[i].CEO = updatedCompany.CEO
			companies[i].Revenue = updatedCompany.Revenue
			return &companies[i], true
		}
	}

	return nil, false
}

func Delete(id string) bool {
	for i, company := range companies {
		if company.ID == id {
			companies = append(companies[:i], companies[i+1:]...)
			return true
		}
	}

	return false
}
