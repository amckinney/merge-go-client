// This file was auto-generated by Fern from our API Definition.

package ats

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type LinkedAccountsListRequest struct {
	// Options: ('hris', 'ats', 'accounting', 'ticketing', 'crm', 'mktg', 'filestorage')
	//
	// - `hris` - hris
	// - `ats` - ats
	// - `accounting` - accounting
	// - `ticketing` - ticketing
	// - `crm` - crm
	// - `mktg` - mktg
	// - `filestorage` - filestorage
	Category *LinkedAccountsListRequestCategory `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return linked accounts associated with the given email address.
	EndUserEmailAddress *string `json:"-"`
	// If provided, will only return linked accounts associated with the given organization name.
	EndUserOrganizationName *string `json:"-"`
	// If provided, will only return linked accounts associated with the given origin ID.
	EndUserOriginId *string `json:"-"`
	// Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
	EndUserOriginIds *string `json:"-"`
	Id               *string `json:"-"`
	// Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
	Ids *string `json:"-"`
	// If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
	IncludeDuplicates *bool `json:"-"`
	// If provided, will only return linked accounts associated with the given integration name.
	IntegrationName *string `json:"-"`
	// If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
	IsTestAccount *string `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// Filter by status. Options: `COMPLETE`, `INCOMPLETE`, `RELINK_NEEDED`
	Status *string `json:"-"`
}

type LinkedAccountsListRequestCategory uint

const (
	LinkedAccountsListRequestCategoryAccounting LinkedAccountsListRequestCategory = iota + 1
	LinkedAccountsListRequestCategoryAts
	LinkedAccountsListRequestCategoryCrm
	LinkedAccountsListRequestCategoryFilestorage
	LinkedAccountsListRequestCategoryHris
	LinkedAccountsListRequestCategoryMktg
	LinkedAccountsListRequestCategoryTicketing
)

func (l LinkedAccountsListRequestCategory) String() string {
	switch l {
	default:
		return strconv.Itoa(int(l))
	case LinkedAccountsListRequestCategoryAccounting:
		return "accounting"
	case LinkedAccountsListRequestCategoryAts:
		return "ats"
	case LinkedAccountsListRequestCategoryCrm:
		return "crm"
	case LinkedAccountsListRequestCategoryFilestorage:
		return "filestorage"
	case LinkedAccountsListRequestCategoryHris:
		return "hris"
	case LinkedAccountsListRequestCategoryMktg:
		return "mktg"
	case LinkedAccountsListRequestCategoryTicketing:
		return "ticketing"
	}
}

func (l LinkedAccountsListRequestCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", l.String())), nil
}

func (l *LinkedAccountsListRequestCategory) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "accounting":
		value := LinkedAccountsListRequestCategoryAccounting
		*l = value
	case "ats":
		value := LinkedAccountsListRequestCategoryAts
		*l = value
	case "crm":
		value := LinkedAccountsListRequestCategoryCrm
		*l = value
	case "filestorage":
		value := LinkedAccountsListRequestCategoryFilestorage
		*l = value
	case "hris":
		value := LinkedAccountsListRequestCategoryHris
		*l = value
	case "mktg":
		value := LinkedAccountsListRequestCategoryMktg
		*l = value
	case "ticketing":
		value := LinkedAccountsListRequestCategoryTicketing
		*l = value
	}
	return nil
}
