package charge

import (
	"github.com/Invoiced/invoiced-go"
	"strconv"
)

const endpoint = "/charges"

type Client struct {
	Api *invoiced.Api
}

func (c *Client) Create(request *invoiced.ChargeRequest) (*invoiced.PaymentClient, error) {
	payment := new(invoiced.PaymentClient)
	err := c.Api.Create(endpoint, request, payment)
	return payment, err
}

func (c *Client) Refund(chargeId int64, request *invoiced.RefundRequest) (*invoiced.Refund, error) {
	endpoint2 := endpoint + "/" + strconv.FormatInt(chargeId, 10) + "/refunds"
	refund := new(invoiced.Refund)
	err := c.Api.Create(endpoint2, request, refund)
	return refund, err
}
