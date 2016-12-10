package ovh

import (
	"fmt"
	"time"
)

// Cart is a go representation of Cart instance
type Cart struct {
	CartID      string     `json:"cartId,omitempty"`
	Expire      *time.Time `json:"expire,omitempty"`
	Description string     `json:"description,omitempty"`
	ReadOnly    bool       `json:"readOnly,omitempty"`
	Items       []int      `json:"items,omitempty"`
}

// CartCreateReq defines the fields for a Cart creation
type CartCreateReq struct {
	Expire        *time.Time `json:"expire,omitempty"`
	Description   string     `json:"description,omitempty"`
	OVHSubsidiary string     `json:"ovhSubsidiary,omitempty"`
}

// CartUpdateReq defines the fields for a Cart creation
type CartUpdateReq struct {
	Expire      *time.Time `json:"expire,omitempty"`
	Description string     `json:"description,omitempty"`
}

// CartList list all your cart
func (c *Client) CartList() ([]Cart, error) {
	var ids []string
	e := c.OVHClient.Get("/order/cart", &ids)
	carts := []Cart{}
	for _, id := range ids {
		carts = append(carts, Cart{CartID: id})
	}
	return carts, e
}

// CartInfo retrieve all infos of one of your cart
func (c *Client) CartInfo(cartID string) (*Cart, error) {
	cart := &Cart{}
	err := c.OVHClient.Get(fmt.Sprintf("/order/cart/%s", cartID), cart)
	return cart, err
}

// CreateCart create a new cart
func (c *Client) CreateCart(cartCreateReq CartCreateReq) (*Cart, error) {
	cart := &Cart{}
	e := c.OVHClient.Post("/order/cart", cartCreateReq, cart)
	return cart, e
}

// UpdateCart update a cart
func (c *Client) UpdateCart(cartID string, cartUpdateReq CartUpdateReq) (*Cart, error) {
	cart := &Cart{}
	e := c.OVHClient.Put(fmt.Sprintf("/order/cart/%s", cartID), cartUpdateReq, cart)
	return cart, e
}

// DeleteCart delete a cart
func (c *Client) DeleteCart(cartID string) error {
	e := c.OVHClient.Delete(fmt.Sprintf("/order/cart/%s", cartID), nil)
	return e
}

// AssignCart assign to connected user a cart
func (c *Client) AssignCart(cartID string) error {
	e := c.OVHClient.Post(fmt.Sprintf("/order/cart/%s/assign", cartID), nil, nil)
	return e
}
