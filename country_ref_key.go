package main

type CountryRefKey struct {
	Odata_metadata string         `json:"odata.metadata"`
	ValueCountries []ValueCountry `json:"value"`
}

type ValueCountry struct {
	Code                          string `json:"Code"`
	DataVersion                   string `json:"DataVersion"`
	DeletionMark                  bool   `json:"DeletionMark"`
	Description                   string `json:"Description"`
	Predefined                    bool   `json:"Predefined"`
	PredefinedDataName            string `json:"PredefinedDataName"`
	RefKey                        string `json:"Ref_Key"`
	CodeAlfa2                     string `json:"КодАльфа2"`
	CodeAlfa3                     string `json:"КодАльфа3"`
	FullName                      string `json:"НаименованиеПолное"`
	MemberOfTCU_navigationLinkURL string `json:"УчастникТаможенногоСоюза@navigationLinkUrl"`
	MemberOfTCU_Key               string `json:"УчастникТаможенногоСоюза_Key"`
}
