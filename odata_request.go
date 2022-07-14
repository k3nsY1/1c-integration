package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	go_odoo "github.com/skilld-labs/go-odoo"
)

type Catalog struct {
	OdataMetadata string  `json:"odata.metadata"`
	Values        []Value `json:"value"`
}

type Value struct {
	RefKey                                   string        `json:"Ref_Key"`
	DataVersion                              string        `json:"DataVersion"`
	DeletionMark                             bool          `json:"DeletionMark"`
	ParentKey                                string        `json:"Parent_Key"`
	IsFolder                                 bool          `json:"IsFolder"`
	Code                                     string        `json:"Code"`
	Description                              string        `json:"Description"`
	HeadCounterpartyKey                      string        `json:"ГоловнойКонтрагент_Key"`
	DateEvidence                             string        `json:"ДатаСвидетельстваПоНДС"`
	IdentityDocument                         string        `json:"ДокументУдостоверяющийЛичность"`
	IdentityCode                             string        `json:"ИдентификационныйКодЛичности"`
	IndividualEnterprenuerLayerPrivateNotary bool          `json:"ИндивидуальныйПредпринимательАдвокатЧастныйНотариус"`
	СBE                                      string        `json:"КБЕ"`
	OKPOCode                                 string        `json:"КодПоОКПО"`
	Comment                                  string        `json:"Комментарий"`
	Fullname                                 string        `json:"НаименованиеПолное"`
	TaxRegistrationNumberCountryOfResidency  string        `json:"НомерНалоговойРегистрацииВСтранеРезидентства"`
	VATCertificateNumber                     string        `json:"НомерСвидетельстваПоНДС"`
	PrimaryContactPersonKey                  string        `json:"ОсновноеКонтактноеЛицо_Key"`
	MainBankAccountKey                       string        `json:"ОсновнойБанковскийСчет_Key"`
	MainContractOfTheCounterpartyKey         string        `json:"ОсновнойДоговорКонтрагента_Key"`
	RNN                                      string        `json:"РНН"`
	SeriesCertificatesVAT                    string        `json:"СерияСвидетельстваПоНДС"`
	SIK                                      string        `json:"СИК"`
	CountryResidenceKey                      string        `json:"СтранаРезидентства_Key"`
	IndividualKey                            string        `json:"ФизЛицо_Key"`
	Entity                                   string        `json:"ЮрФизЛицо"`
	DetailsOfTheParentOrganization           bool          `json:"УказыватьРеквизитыГоловнойОрганизацииВСчетеФактуре"`
	GovernmentAgency                         bool          `json:"ГосударственноеУчреждение"`
	SmallOutlet                              bool          `json:"МалаяТорговаяТочка"`
	AdditionalDetails                        []interface{} `json:"ДополнительныеРеквизиты"`
	Predefined                               bool          `json:"Predefined"`
	PredefinedDataName                       string        `json:"PredefinedDataName"`
	ParentNavigationLinkURL                  string        `json:"Parent@navigationLinkUrl,omitempty"`
	HeadCounterparty                         string        `json:"ГоловнойКонтрагент@navigationLinkUrl,omitempty"`
	MainContractOfTheCounterparty            string        `json:"ОсновнойДоговорКонтрагента@navigationLinkUrl,omitempty"`
	CountryResidence                         string        `json:"СтранаРезидентства@navigationLinkUrl,omitempty"`
	Individual                               string        `json:"ФизЛицо@navigationLinkUrl,omitempty"`
	MainBankAccount                          string        `json:"ОсновнойБанковскийСчет@navigationLinkUrl,omitempty"`
}

func RequestOdata() []byte {
	url := "url/odata/standard.odata/Catalog_Контрагенты?$format=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.SetBasicAuth("login", "password")
	res, err := client.Do(req)
	if err != nil {
		log.Panic()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic()
	}

	return body
}

func RequestByGuid(model string, id string) []byte {
	url := "url" + model + "('" + id + "')?$format=json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("login", "password")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func GetCatalogFrom1c() {
	// Запрос в 1С что получит список план счетов
	result := RequestOdata()
	catalog := &Catalog{}

	err := json.Unmarshal(result, &catalog)
	if err != nil {
		fmt.Println(err, "Error Unmarshal")
	}

	for _, value := range catalog.Values {

		// Поиск по названию в Odoo
		resPartner, err := ClientPm.FindResPartner(go_odoo.NewCriteria().Add("display_name", "ilike", value.Fullname))
		if err != nil {
			fmt.Println(err, ReplacePrefix(value.Fullname), "-----")
			// CreateIfNotExsist(value) // создать если нет в Odoo
		}

		fmt.Println(resPartner.Name, "=====")
		// (мапа ключ значение) для обновления поля (где ключ'string' это название поля которое мы хотим изменить а значение'interface{}' то что мы меняем )
		updates := make(map[string]interface{})
		if value.CountryResidenceKey != "00000000-0000-0000-0000-000000000000" {
			countryRefKey := RequestByGuid("Catalog_КлассификаторСтранМира", value.CountryResidenceKey)

			countryKey := &ValueCountry{}
			err = json.Unmarshal(countryRefKey, &countryKey)
			if err != nil {
				fmt.Println(err, "Error Unmarshal")
				return
			}

			updates["country_code"] = countryKey.CodeAlfa3
		}
		if value.ParentKey != "00000000-0000-0000-0000-000000000000" {
			result := RequestByGuid("Catalog_Контрагенты", value.ParentKey)
			parent := &Value{}

			err = json.Unmarshal(result, &parent)
			if err != nil {
				fmt.Println(err, "Error Unmarshal")
			} else if parent.Description == "Юридические лица" {

				updates["l10n_ca_pst"] = value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT
				updates["bin"] = value.IdentityCode
				updates["company_type"] = "company"
				updates["is_company"] = true
			} else if parent.Description == "Физические лица" {

				updates["name_kz"] = value.IdentityDocument
				updates["l10n_ca_pst"] = value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT
				updates["iin"] = value.IdentityCode
				updates["company_type"] = "person"
				updates["is_company"] = false
			} else if parent.Description == "Учредители" {

				updates["individual_type"] = "employee"
				updates["name_kz"] = value.IdentityDocument
				updates["l10n_ca_pst"] = value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT
				updates["iin"] = value.IdentityCode
				updates["company_type"] = "person"
				updates["is_company"] = false
			} else if parent.Description == "ГПХ" {

				updates["individual_type"] = "civil_contract"
				updates["name_kz"] = value.IdentityDocument
				updates["l10n_ca_pst"] = value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT
				updates["iin"] = value.IdentityCode
				updates["company_type"] = "person"
				updates["is_company"] = false
			} else if parent.Description == "Розничная выручка" {
				updates["individual_type"] = "other"
				updates["company_type"] = "person"
				updates["is_company"] = false
			}
		}
		if value.MainBankAccountKey != "00000000-0000-0000-0000-000000000000" {
			result := RequestByGuid("Catalog_БанковскиеСчета", value.MainBankAccountKey)

			bankAccount := &Value{}
			err := json.Unmarshal(result, &bankAccount)
			if err != nil {
				fmt.Println(err, "Error Unmarshal")
			}
			fmt.Println(bankAccount.Description)
			regName := regexp.MustCompile(`"([^\"]+)"`)
			name := regName.FindString(bankAccount.Description)
			regBic := regexp.MustCompile(`[A-Z,0-9]{20}`)
			accNumber := regBic.FindString(bankAccount.Description)

			bank, err := ClientPm.FindResBank(go_odoo.NewCriteria().Add("name", "ilike", ReplacePrefix(name)))
			if err != nil {
				fmt.Println("ERRR", err)
			}

			resPartnerBank := &go_odoo.ResPartnerBank{
				AccNumber: go_odoo.NewString(accNumber),
				BankBic:   go_odoo.NewString(bank.Bic.Get()),
				BankId:    go_odoo.NewMany2One(bank.Id.Get(), ""),
				BankName:  go_odoo.NewString(bank.Name.Get()),
				PartnerId: go_odoo.NewMany2One(resPartner.Id.Get(), ""),
			}

			newId, err := ClientPm.CreateResPartnerBank(resPartnerBank)
			if err != nil {
				fmt.Println("EROOR", err)
			}
			fmt.Println("ID", newId)
		}
		err = ClientPm.Update("res.partner", []int64{resPartner.Id.Get()}, updates)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ReplacePrefix(text string) string {

	text = strings.ReplaceAll(text, "АО", "")
	text = strings.ReplaceAll(text, "ИП", "")
	text = strings.ReplaceAll(text, "ООО", "")
	text = strings.ReplaceAll(text, `"`, "")
	text = strings.ReplaceAll(text, "РГУ", "")

	return text
}

func CreateIfNotExsist(value Value) {

	aa := go_odoo.ResPartner{

		Name:        go_odoo.NewString(value.Fullname),
		DisplayName: go_odoo.NewString(value.Fullname),
	}
	if value.CountryResidenceKey != "00000000-0000-0000-0000-000000000000" {
		countryRefKey := RequestByGuid("Catalog_КлассификаторСтранМира", value.CountryResidenceKey)

		countryKey := &ValueCountry{}
		err := json.Unmarshal(countryRefKey, &countryKey)
		if err != nil {
			fmt.Println(err, "Error Unmarshal")
			return
		}

		aa.CountryCode = go_odoo.NewString(countryKey.CodeAlfa3)
	}
	if value.ParentKey != "00000000-0000-0000-0000-000000000000" {
		result := RequestByGuid("Catalog_Контрагенты", value.ParentKey)
		parent := &Value{}

		err := json.Unmarshal(result, &parent)
		if err != nil {
			fmt.Println(err, "Error Unmarshal")
		} else if parent.Description == "Юридические лица" {

			aa.L10NCaPst = go_odoo.NewString(value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT)
			aa.Bin = go_odoo.NewString(value.IdentityCode)
			aa.CompanyType = go_odoo.NewSelection("company")
			aa.IsCompany = go_odoo.NewBool(true)
		} else if parent.Description == "Физические лица" {

			aa.NameKz = go_odoo.NewString(value.IdentityCode)
			aa.L10NCaPst = go_odoo.NewString(value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT)
			aa.Iin = go_odoo.NewString(value.IdentityCode)
			aa.CompanyType = go_odoo.NewSelection("person")
			aa.IsCompany = go_odoo.NewBool(false)
		} else if parent.Description == "Учредители" {

			aa.IndividualType = go_odoo.NewSelection("employee")
			aa.NameKz = go_odoo.NewString(value.IdentityDocument)
			aa.L10NCaPst = go_odoo.NewString(value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT)
			aa.Iin = go_odoo.NewString(value.IdentityCode)
			aa.CompanyType = go_odoo.NewSelection("person")

		} else if parent.Description == "ГПХ" {

			aa.IndividualType = go_odoo.NewSelection("individual_type")
			aa.NameKz = go_odoo.NewString(value.IdentityCode)
			aa.L10NCaPst = go_odoo.NewString(value.DateEvidence + value.VATCertificateNumber + value.SeriesCertificatesVAT)
			aa.Iin = go_odoo.NewString(value.IdentityCode)
			aa.CompanyType = go_odoo.NewSelection("person")
			aa.IsCompany = go_odoo.NewBool(false)
		} else if parent.Description == "Розничная выручка" {
			aa.IndividualType = go_odoo.NewSelection("other")
			aa.CompanyType = go_odoo.NewSelection("person")
			aa.IsCompany = go_odoo.NewBool(false)
		}
	}
	if value.MainBankAccountKey != "00000000-0000-0000-0000-000000000000" {
		result := RequestByGuid("Catalog_БанковскиеСчета", value.MainBankAccountKey)

		bankAccount := &Value{}
		err := json.Unmarshal(result, &bankAccount)
		if err != nil {
			fmt.Println(err, "Error Unmarshal")
		}
		fmt.Println(bankAccount.Description)
		regName := regexp.MustCompile(`"([^\"]+)"`)
		name := regName.FindString(bankAccount.Description)
		regBic := regexp.MustCompile(`[A-Z,0-9]{20}`)
		accNumber := regBic.FindString(bankAccount.Description)

		bank, err := ClientPm.FindResBank(go_odoo.NewCriteria().Add("name", "ilike", ReplacePrefix(name)))
		if err != nil {
			fmt.Println("ERRR", err)
		}

		resPartnerBank := &go_odoo.ResPartnerBank{
			AccNumber: go_odoo.NewString(accNumber),
			BankBic:   go_odoo.NewString(bank.Bic.Get()),
			BankId:    go_odoo.NewMany2One(bank.Id.Get(), ""),
			BankName:  go_odoo.NewString(bank.Name.Get()),
			PartnerId: go_odoo.NewMany2One(aa.Id.Get(), ""),
		}

		newId, err := ClientPm.CreateResPartnerBank(resPartnerBank)
		if err != nil {
			fmt.Println("EROOR", err)
		}
		fmt.Println("ID", newId)

	}
}
