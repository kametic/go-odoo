package odoo

import (
	"fmt"
)

// AccountOnlineProvider represents account.online.provider model.
type AccountOnlineProvider struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	AccountOnlineJournalIds    *Relation  `xmlrpc:"account_online_journal_ids,omptempty"`
	ActionRequired             *Bool      `xmlrpc:"action_required,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	LastRefresh                *Time      `xmlrpc:"last_refresh,omptempty"`
	Message                    *String    `xmlrpc:"message,omptempty"`
	MessageChannelIds          *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost            *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread              *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter       *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	NextRefresh                *Time      `xmlrpc:"next_refresh,omptempty"`
	ProviderAccountIdentifier  *String    `xmlrpc:"provider_account_identifier,omptempty"`
	ProviderIdentifier         *String    `xmlrpc:"provider_identifier,omptempty"`
	ProviderType               *Selection `xmlrpc:"provider_type,omptempty"`
	Status                     *String    `xmlrpc:"status,omptempty"`
	StatusCode                 *Int       `xmlrpc:"status_code,omptempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
	YodleeAdditionalStatus     *String    `xmlrpc:"yodlee_additional_status,omptempty"`
	YodleeLastAttemptedRefresh *Time      `xmlrpc:"yodlee_last_attempted_refresh,omptempty"`
	YodleeNextScheduleRefresh  *Time      `xmlrpc:"yodlee_next_schedule_refresh,omptempty"`
}

// AccountOnlineProviders represents array of account.online.provider model.
type AccountOnlineProviders []AccountOnlineProvider

// AccountOnlineProviderModel is the odoo model name.
const AccountOnlineProviderModel = "account.online.provider"

// Many2One convert AccountOnlineProvider to *Many2One.
func (aop *AccountOnlineProvider) Many2One() *Many2One {
	return NewMany2One(aop.Id.Get(), "")
}

// CreateAccountOnlineProvider creates a new account.online.provider model and returns its id.
func (c *Client) CreateAccountOnlineProvider(aop *AccountOnlineProvider) (int64, error) {
	return c.Create(AccountOnlineProviderModel, aop)
}

// UpdateAccountOnlineProvider updates an existing account.online.provider record.
func (c *Client) UpdateAccountOnlineProvider(aop *AccountOnlineProvider) error {
	return c.UpdateAccountOnlineProviders([]int64{aop.Id.Get()}, aop)
}

// UpdateAccountOnlineProviders updates existing account.online.provider records.
// All records (represented by ids) will be updated by aop values.
func (c *Client) UpdateAccountOnlineProviders(ids []int64, aop *AccountOnlineProvider) error {
	return c.Update(AccountOnlineProviderModel, ids, aop)
}

// DeleteAccountOnlineProvider deletes an existing account.online.provider record.
func (c *Client) DeleteAccountOnlineProvider(id int64) error {
	return c.DeleteAccountOnlineProviders([]int64{id})
}

// DeleteAccountOnlineProviders deletes existing account.online.provider records.
func (c *Client) DeleteAccountOnlineProviders(ids []int64) error {
	return c.Delete(AccountOnlineProviderModel, ids)
}

// GetAccountOnlineProvider gets account.online.provider existing record.
func (c *Client) GetAccountOnlineProvider(id int64) (*AccountOnlineProvider, error) {
	aops, err := c.GetAccountOnlineProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	if aops != nil && len(*aops) > 0 {
		return &((*aops)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.online.provider not found", id)
}

// GetAccountOnlineProviders gets account.online.provider existing records.
func (c *Client) GetAccountOnlineProviders(ids []int64) (*AccountOnlineProviders, error) {
	aops := &AccountOnlineProviders{}
	if err := c.Read(AccountOnlineProviderModel, ids, nil, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAccountOnlineProvider finds account.online.provider record by querying it with criteria.
func (c *Client) FindAccountOnlineProvider(criteria *Criteria) (*AccountOnlineProvider, error) {
	aops := &AccountOnlineProviders{}
	if err := c.SearchRead(AccountOnlineProviderModel, criteria, NewOptions().Limit(1), aops); err != nil {
		return nil, err
	}
	if aops != nil && len(*aops) > 0 {
		return &((*aops)[0]), nil
	}
	return nil, fmt.Errorf("account.online.provider was not found")
}

// FindAccountOnlineProviders finds account.online.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineProviders(criteria *Criteria, options *Options) (*AccountOnlineProviders, error) {
	aops := &AccountOnlineProviders{}
	if err := c.SearchRead(AccountOnlineProviderModel, criteria, options, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAccountOnlineProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountOnlineProviderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountOnlineProviderId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.online.provider was not found")
}
