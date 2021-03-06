package harvest

import (
	// "fmt"

	"github.com/polarsquad/harvest/config"
	"github.com/polarsquad/harvest/structs"
)

// VERSION v0.0.2
const VERSION = "v0.0.2"

// Harvest creates the struct for the API, User and Entries
type Harvest struct {
	API         *structs.API
	User        *User
	Project     string
	TimeEntries *TimeEntries
}

type api structs.API

type TimeEntries structs.TimeEntries

type Entries []structs.Entries

// User ...
type User structs.User

// Links ...
type Links structs.Links

// GetTimeEntriesParams ...
type GetTimeEntriesParams struct {
	From   string `url:"from"`
	To     string `url:"to"`
	UserID int64  `url:"user_id"`
	// ClientID  string `url:"client_id"`
	// ProjectID string `url:"project_id"`
	// IsBilled  bool   `url:"is_billed"`
	// Page      int16  `url:"page"`
	PerPage int8 `url:"per_page"`
}

type API structs.API

// Init methot initializes the data structure needed for Harvest
func Init(conf *config.Config) *Harvest {
	a := &structs.API{
		AuthToken: conf.API.AuthToken,
		AccountID: conf.API.AccountID,
	}

	e := &TimeEntries{}

	H := &Harvest{
		API:         a,
		User:        &User{},
		Project:     "",
		TimeEntries: e,
	}

	// API.AccountID = conf.API.AccountID
	// API.AuthToken = conf.API.AuthToken
	// API.BaseURL = conf.API.BaseURL
	// h := &Harvest{
	// 	User:     "Mika",
	// 	Projects: "Client",
	// 	Entries:  []TimeEntries{},
	// }

	return H
}
