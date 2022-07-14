package odoo

import (
	"fmt"
)

// PurchaseRequisition represents purchase.requisition model.
type PurchaseRequisition struct {
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsQuantityCopy              *Selection `xmlrpc:"is_quantity_copy,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	LineIds                     *Relation  `xmlrpc:"line_ids,omptempty"`
	MaintenanceType             *Many2One  `xmlrpc:"maintenance_type,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	OrderCount                  *Int       `xmlrpc:"order_count,omptempty"`
	OrderingDate                *Time      `xmlrpc:"ordering_date,omptempty"`
	Origin                      *String    `xmlrpc:"origin,omptempty"`
	PartIds                     *Relation  `xmlrpc:"part_ids,omptempty"`
	PickingTypeId               *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omptempty"`
	PurchaseIds                 *Relation  `xmlrpc:"purchase_ids,omptempty"`
	ScheduleDate                *Time      `xmlrpc:"schedule_date,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	StateBlanketOrder           *Selection `xmlrpc:"state_blanket_order,omptempty"`
	ToolIds                     *Relation  `xmlrpc:"tool_ids,omptempty"`
	TotalCount                  *Float     `xmlrpc:"total_count,omptempty"`
	TotalUnitPrice              *Float     `xmlrpc:"total_unit_price,omptempty"`
	TypeId                      *Many2One  `xmlrpc:"type_id,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	VendorId                    *Many2One  `xmlrpc:"vendor_id,omptempty"`
	WarehouseId                 *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PurchaseRequisitions represents array of purchase.requisition model.
type PurchaseRequisitions []PurchaseRequisition

// PurchaseRequisitionModel is the odoo model name.
const PurchaseRequisitionModel = "purchase.requisition"

// Many2One convert PurchaseRequisition to *Many2One.
func (pr *PurchaseRequisition) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePurchaseRequisition creates a new purchase.requisition model and returns its id.
func (c *Client) CreatePurchaseRequisition(pr *PurchaseRequisition) (int64, error) {
	return c.Create(PurchaseRequisitionModel, pr)
}

// UpdatePurchaseRequisition updates an existing purchase.requisition record.
func (c *Client) UpdatePurchaseRequisition(pr *PurchaseRequisition) error {
	return c.UpdatePurchaseRequisitions([]int64{pr.Id.Get()}, pr)
}

// UpdatePurchaseRequisitions updates existing purchase.requisition records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePurchaseRequisitions(ids []int64, pr *PurchaseRequisition) error {
	return c.Update(PurchaseRequisitionModel, ids, pr)
}

// DeletePurchaseRequisition deletes an existing purchase.requisition record.
func (c *Client) DeletePurchaseRequisition(id int64) error {
	return c.DeletePurchaseRequisitions([]int64{id})
}

// DeletePurchaseRequisitions deletes existing purchase.requisition records.
func (c *Client) DeletePurchaseRequisitions(ids []int64) error {
	return c.Delete(PurchaseRequisitionModel, ids)
}

// GetPurchaseRequisition gets purchase.requisition existing record.
func (c *Client) GetPurchaseRequisition(id int64) (*PurchaseRequisition, error) {
	prs, err := c.GetPurchaseRequisitions([]int64{id})
	if err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of purchase.requisition not found", id)
}

// GetPurchaseRequisitions gets purchase.requisition existing records.
func (c *Client) GetPurchaseRequisitions(ids []int64) (*PurchaseRequisitions, error) {
	prs := &PurchaseRequisitions{}
	if err := c.Read(PurchaseRequisitionModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseRequisition finds purchase.requisition record by querying it with criteria.
func (c *Client) FindPurchaseRequisition(criteria *Criteria) (*PurchaseRequisition, error) {
	prs := &PurchaseRequisitions{}
	if err := c.SearchRead(PurchaseRequisitionModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("purchase.requisition was not found")
}

// FindPurchaseRequisitions finds purchase.requisition records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseRequisitions(criteria *Criteria, options *Options) (*PurchaseRequisitions, error) {
	prs := &PurchaseRequisitions{}
	if err := c.SearchRead(PurchaseRequisitionModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseRequisitionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseRequisitionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PurchaseRequisitionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPurchaseRequisitionId finds record id by querying it with criteria.
func (c *Client) FindPurchaseRequisitionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseRequisitionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("purchase.requisition was not found")
}
