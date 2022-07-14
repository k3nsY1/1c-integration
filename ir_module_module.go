package main

type IrModuleModule struct {
	ID               int           `json:"id"`
	Shortdesc        string        `json:"shortdesc"`
	Name             string        `json:"name"`
	Author           string        `json:"author"`
	Website          string        `json:"website"`
	InstalledVersion string        `json:"installed_version"`
	State            string        `json:"state"`
	CategoryID       []interface{} `json:"category_id"`
}
