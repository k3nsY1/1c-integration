package odoo

import (
	"fmt"
)

// ResPartner represents res.partner model.
type ResPartner struct {
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omptempty"`
	ActiveLimit                   *Bool      `xmlrpc:"active_limit,omptempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ActualAddress                 *Many2One  `xmlrpc:"actual_address,omptempty"`
	AdditionalInfo                *String    `xmlrpc:"additional_info,omptempty"`
	AttorneyDate                  *Time      `xmlrpc:"attorney_date,omptempty"`
	AttorneyNumber                *String    `xmlrpc:"attorney_number,omptempty"`
	Bank                          *Many2One  `xmlrpc:"bank,omptempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omptempty"`
	Bin                           *String    `xmlrpc:"bin,omptempty"`
	BlockingStage                 *Float     `xmlrpc:"blocking_stage,omptempty"`
	Branches                      *String    `xmlrpc:"branches,omptempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanPublish                    *Bool      `xmlrpc:"can_publish,omptempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omptempty"`
	CertificationsCompanyCount    *Int       `xmlrpc:"certifications_company_count,omptempty"`
	CertificationsCount           *Int       `xmlrpc:"certifications_count,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	City                          *String    `xmlrpc:"city,omptempty"`
	CodesKatoId                   *Many2One  `xmlrpc:"codes_kato_id,omptempty"`
	Color                         *Int       `xmlrpc:"color,omptempty"`
	Comment                       *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyModels                 *Relation  `xmlrpc:"company_models,omptempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omptempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omptempty"`
	Concretness                   *String    `xmlrpc:"concretness,omptempty"`
	ConferenceHalls               *Relation  `xmlrpc:"conference_halls,omptempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omptempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omptempty"`
	CountConHalls                 *String    `xmlrpc:"count_con_halls,omptempty"`
	CountryCode                   *String    `xmlrpc:"country_code,omptempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omptempty"`
	Crdate                        *Int       `xmlrpc:"crdate,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                        *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	Customer                      *Int       `xmlrpc:"customer,omptempty"`
	CustomerRank                  *Int       `xmlrpc:"customer_rank,omptempty"`
	Date                          *Time      `xmlrpc:"date,omptempty"`
	DateOfRegistration            *Time      `xmlrpc:"date_of_registration,omptempty"`
	Debit                         *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omptempty"`
	Director                      *String    `xmlrpc:"director,omptempty"`
	DirectorInitials              *String    `xmlrpc:"director_initials,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	DueAmount                     *Float     `xmlrpc:"due_amount,omptempty"`
	Email                         *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omptempty"`
	EmailNormalized               *String    `xmlrpc:"email_normalized,omptempty"`
	Employee                      *Bool      `xmlrpc:"employee,omptempty"`
	EnableCreditLimit             *Bool      `xmlrpc:"enable_credit_limit,omptempty"`
	EventCount                    *Int       `xmlrpc:"event_count,omptempty"`
	FollowupStatus                *Selection `xmlrpc:"followup_status,omptempty"`
	FullNameKz                    *String    `xmlrpc:"full_name_kz,omptempty"`
	FullNameRu                    *String    `xmlrpc:"full_name_ru,omptempty"`
	Function                      *String    `xmlrpc:"function,omptempty"`
	GroupByAb                     *Selection `xmlrpc:"group_by_ab,omptempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	Hospital                      *Many2One  `xmlrpc:"hospital,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	Iin                           *String    `xmlrpc:"iin,omptempty"`
	Image1024                     *String    `xmlrpc:"image_1024,omptempty"`
	Image128                      *String    `xmlrpc:"image_128,omptempty"`
	Image1920                     *String    `xmlrpc:"image_1920,omptempty"`
	Image256                      *String    `xmlrpc:"image_256,omptempty"`
	Image512                      *String    `xmlrpc:"image_512,omptempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omptempty"`
	IndexDate                     *String    `xmlrpc:"index_date,omptempty"`
	IndividualType                *Selection `xmlrpc:"individual_type,omptempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omptempty"`
	Inn                           *String    `xmlrpc:"inn,omptempty"`
	Instructor                    *Bool      `xmlrpc:"instructor,omptempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceList                   *Relation  `xmlrpc:"invoice_list,omptempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsAgentField                  *Bool      `xmlrpc:"is_Agent_field,omptempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omptempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omptempty"`
	IsGos                         *String    `xmlrpc:"is_gos,omptempty"`
	IsPatient                     *Bool      `xmlrpc:"is_patient,omptempty"`
	IsPhysician                   *Bool      `xmlrpc:"is_physician,omptempty"`
	IsPublished                   *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	IsSingleOrg                   *Int       `xmlrpc:"is_single_org,omptempty"`
	Jitsi                         *String    `xmlrpc:"jitsi,omptempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omptempty"`
	KatoIds                       *Many2One  `xmlrpc:"kato_ids,omptempty"`
	KatoList                      *Many2One  `xmlrpc:"kato_list,omptempty"`
	KatoName                      *String    `xmlrpc:"kato_name,omptempty"`
	KrpCode                       *Int       `xmlrpc:"krp_code,omptempty"`
	KrpText                       *String    `xmlrpc:"krp_text,omptempty"`
	KseCode                       *Int       `xmlrpc:"kse_code,omptempty"`
	L10NCaPst                     *String    `xmlrpc:"l10n_ca_pst,omptempty"`
	L10NCaPstField                *Bool      `xmlrpc:"l10n_ca_pst_field,omptempty"`
	Lang                          *Selection `xmlrpc:"lang,omptempty"`
	LangField                     *Bool      `xmlrpc:"lang_field,omptempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	LastUpdateDate                *Time      `xmlrpc:"last_update_date,omptempty"`
	LastWebsiteSoId               *Many2One  `xmlrpc:"last_website_so_id,omptempty"`
	LegalAddress                  *Many2One  `xmlrpc:"legal_address,omptempty"`
	MarkAssocWithDisab            *Int       `xmlrpc:"mark_assoc_with_disab,omptempty"`
	MarkNationalCompany           *Int       `xmlrpc:"mark_national_company,omptempty"`
	MarkNaturalMonopoly           *Int       `xmlrpc:"mark_natural_monopoly,omptempty"`
	MarkPatronymicProducer        *Int       `xmlrpc:"mark_patronymic_producer,omptempty"`
	MarkPatronymicSupplyer        *Int       `xmlrpc:"mark_patronymic_supplyer,omptempty"`
	MarkResident                  *Int       `xmlrpc:"mark_resident,omptempty"`
	MarkSmallEmployer             *Int       `xmlrpc:"mark_small_employer,omptempty"`
	MarkStateMonopoly             *Int       `xmlrpc:"mark_state_monopoly,omptempty"`
	MarkWorldCompany              *Int       `xmlrpc:"mark_world_company,omptempty"`
	MeetingCount                  *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                    *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mobile                        *String    `xmlrpc:"mobile,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	NameField                     *Many2One  `xmlrpc:"name_field,omptempty"`
	NameKz                        *String    `xmlrpc:"name_kz,omptempty"`
	NameRu                        *String    `xmlrpc:"name_ru,omptempty"`
	NextReminderDate              *Time      `xmlrpc:"next_reminder_date,omptempty"`
	NumberReg                     *String    `xmlrpc:"number_reg,omptempty"`
	OkedList                      *Int       `xmlrpc:"oked_list,omptempty"`
	OkedName                      *Many2One  `xmlrpc:"oked_name,omptempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityCountIds           *Relation  `xmlrpc:"opportunity_count_ids,omptempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	Organizer                     *Int       `xmlrpc:"organizer,omptempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omptempty"`
	ParentSubject                 *String    `xmlrpc:"parent_subject,omptempty"`
	PartnerGid                    *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omptempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	Phone                         *String    `xmlrpc:"phone,omptempty"`
	PhoneBlacklisted              *Bool      `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneSanitized                *String    `xmlrpc:"phone_sanitized,omptempty"`
	PickingWarn                   *Selection `xmlrpc:"picking_warn,omptempty"`
	PickingWarnMsg                *String    `xmlrpc:"picking_warn_msg,omptempty"`
	Pid                           *Int       `xmlrpc:"pid,omptempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyPurchaseCurrencyId    *Many2One  `xmlrpc:"property_purchase_currency_id,omptempty"`
	PropertyStockCustomer         *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier         *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	PurchaseIds                   *Relation  `xmlrpc:"purchase_ids,omptempty"`
	PurchaseOrderCount            *Int       `xmlrpc:"purchase_order_count,omptempty"`
	PurchaseWarn                  *Selection `xmlrpc:"purchase_warn,omptempty"`
	PurchaseWarnMsg               *String    `xmlrpc:"purchase_warn_msg,omptempty"`
	Qvazi                         *Int       `xmlrpc:"qvazi,omptempty"`
	Ref                           *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	RefKopfCode                   *String    `xmlrpc:"ref_kopf_code,omptempty"`
	RefKopfCodeFull               *String    `xmlrpc:"ref_kopf_code_full,omptempty"`
	RefKopfCodeModel              *Many2One  `xmlrpc:"ref_kopf_code_model,omptempty"`
	RefKopfCodeSelection          *Selection `xmlrpc:"ref_kopf_code_selection,omptempty"`
	Regdate                       *Time      `xmlrpc:"regdate,omptempty"`
	RegisteredBy                  *String    `xmlrpc:"registered_by,omptempty"`
	RegistrationNumber            *Int       `xmlrpc:"registration_number,omptempty"`
	Representative                *Bool      `xmlrpc:"representative,omptempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omptempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	Self                          *Many2One  `xmlrpc:"self,omptempty"`
	Series                        *String    `xmlrpc:"series,omptempty"`
	SessionIds                    *Relation  `xmlrpc:"session_ids,omptempty"`
	Shortcut                      *String    `xmlrpc:"shortcut,omptempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omptempty"`
	SlideChannelCompanyCount      *Int       `xmlrpc:"slide_channel_company_count,omptempty"`
	SlideChannelCount             *Int       `xmlrpc:"slide_channel_count,omptempty"`
	SlideChannelIds               *Relation  `xmlrpc:"slide_channel_ids,omptempty"`
	Sms                           *String    `xmlrpc:"sms,omptempty"`
	Speciality                    *Many2One  `xmlrpc:"speciality,omptempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                        *String    `xmlrpc:"street,omptempty"`
	Street2                       *String    `xmlrpc:"street2,omptempty"`
	Supplier                      *Int       `xmlrpc:"supplier,omptempty"`
	SupplierInvoiceCount          *Int       `xmlrpc:"supplier_invoice_count,omptempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omptempty"`
	SystemId                      *Int       `xmlrpc:"system_id,omptempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omptempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omptempty"`
	Telegram                      *String    `xmlrpc:"telegram,omptempty"`
	Title                         *Many2One  `xmlrpc:"title,omptempty"`
	TotalDue                      *Float     `xmlrpc:"total_due,omptempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omptempty"`
	TotalOverdue                  *Float     `xmlrpc:"total_overdue,omptempty"`
	Trust                         *Selection `xmlrpc:"trust,omptempty"`
	Type                          *Selection `xmlrpc:"type,omptempty"`
	TypeCompany                   *Selection `xmlrpc:"type_company,omptempty"`
	TypeOfCountry                 *Selection `xmlrpc:"type_of_country,omptempty"`
	TypeSupplier                  *Int       `xmlrpc:"type_supplier,omptempty"`
	Tz                            *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omptempty"`
	Unp                           *String    `xmlrpc:"unp,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                           *String    `xmlrpc:"vat,omptempty"`
	VerificationAgent             *Bool      `xmlrpc:"verification_agent,omptempty"`
	VisitorIds                    *Relation  `xmlrpc:"visitor_ids,omptempty"`
	WarningStage                  *Float     `xmlrpc:"warning_stage,omptempty"`
	Website                       *String    `xmlrpc:"website,omptempty"`
	WebsiteDescription            *String    `xmlrpc:"website_description,omptempty"`
	WebsiteField                  *Bool      `xmlrpc:"website_field,omptempty"`
	WebsiteId                     *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription        *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords           *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg              *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle              *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription       *String    `xmlrpc:"website_short_description,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	Whatsapp                      *String    `xmlrpc:"whatsapp,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	Year                          *Int       `xmlrpc:"year,omptempty"`
	Zip                           *String    `xmlrpc:"zip,omptempty"`
}

// ResPartners represents array of res.partner model.
type ResPartners []ResPartner

// ResPartnerModel is the odoo model name.
const ResPartnerModel = "res.partner"

// Many2One convert ResPartner to *Many2One.
func (rp *ResPartner) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartner(rp *ResPartner) (int64, error) {
	return c.Create(ResPartnerModel, rp)
}

// UpdateResPartner updates an existing res.partner record.
func (c *Client) UpdateResPartner(rp *ResPartner) error {
	return c.UpdateResPartners([]int64{rp.Id.Get()}, rp)
}

// UpdateResPartners updates existing res.partner records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateResPartners(ids []int64, rp *ResPartner) error {
	return c.Update(ResPartnerModel, ids, rp)
}

// DeleteResPartner deletes an existing res.partner record.
func (c *Client) DeleteResPartner(id int64) error {
	return c.DeleteResPartners([]int64{id})
}

// DeleteResPartners deletes existing res.partner records.
func (c *Client) DeleteResPartners(ids []int64) error {
	return c.Delete(ResPartnerModel, ids)
}

// GetResPartner gets res.partner existing record.
func (c *Client) GetResPartner(id int64) (*ResPartner, error) {
	rps, err := c.GetResPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner not found", id)
}

// GetResPartners gets res.partner existing records.
func (c *Client) GetResPartners(ids []int64) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.Read(ResPartnerModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartner finds res.partner record by querying it with criteria.
func (c *Client) FindResPartner(criteria *Criteria) (*ResPartner, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("res.partner was not found")
}

// FindResPartners finds res.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartners(criteria *Criteria, options *Options) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerId finds record id by querying it with criteria.
func (c *Client) FindResPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner was not found")
}
