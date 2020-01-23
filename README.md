# ipn
--
    import "github.com/jackdoe/gin-ipn"

gin friendly paypal ipn listener Example:

    package main

    import (
    	"log"

    	"github.com/gin-gonic/gin"
    	"github.com/jackdoe/gin-ipn/ipn"
    )

    func main() {
    	r := gin.Default()

    	ipn.Listener(r, "/ipn/:paymentID", func(c *gin.Context, err error, body string, n *ipn.Notification) error {
    		if err != nil {
    			panic(err)
    		}

    		if n.TestIPN {
    			log.Printf("test")
    		}

    		log.Printf("notification: %v", n)

    		return nil
    	})
    }

## Usage

```go
const (
	ReasonChargeback = "Chargeback Settlement"
	ReasonAdmin      = "Admin reversal"
	ReasonRefund     = "Refund"
)
```
Negative amount reasons

```go
const LiveIPNEndpoint = "https://www.paypal.com/cgi-bin/webscr"
```
LiveIPNEndpoint contains the notification verification URL

```go
const SandboxIPNEndpoint = "https://ipnpb.sandbox.paypal.com/cgi-bin/webscr"
```
SandboxIPNEndpoint is the Sandbox notification verification URL

#### func  Listener

```go
func Listener(g *gin.Engine, path string, cb func(c *gin.Context, err error, body string, n *Notification) error)
```
Listener creates a PayPal listener. if err is set in cb, PayPal will resend the
notification at some future point.

#### type Notification

```go
type Notification struct {
	TxnType          string `schema:"txn_type"`
	TxnID            string `schema:"txn_id"`
	Business         string `schema:"business"`
	Custom           string `schema:"custom"`
	ParentTxnID      string `schema:"parent_txn_id"`
	ReceiptID        string `schema:"receipt_id"`
	RecieverEmail    string `schema:"receiver_email"`
	RecieverID       string `schema:"receiver_id"`
	Resend           bool   `schema:"resend"`
	ResidenceCountry string `schema:"residence_country"`
	TestIPN          bool   `schema:"test_ipn"`
	ItemName         string `schema:"item_name"`
	ItemNumber       string `schema:"item_number"`

	//Buyer address information
	AddressCountry     string `schema:"address_country"`
	AddressCity        string `schema:"address_city"`
	AddressCountryCode string `schema:"address_country_code"`
	AddressName        string `schema:"address_name"`
	AddressState       string `schema:"address_state"`
	AddressStatus      string `schema:"address_status"`
	AddressStreet      string `schema:"address_street"`
	AddressZip         string `schema:"address_zip"`

	//Misc buyer info
	ContactPhone      string `schema:"contact_phone"`
	FirstName         string `schema:"first_name"`
	LastName          string `schema:"last_name"`
	PayerBusinessName string `schema:"payer_business_name"`
	PayerEmail        string `schema:"payer_email"`
	PayerID           string `schema:"payer_id"`
	PayerStatus       string `schema:"payer_status"`

	AuthAmount string `schema:"auth_amount"`
	AuthExpire string `schema:"auth_exp"`
	AuthIfD    string `schema:"auth_id"`
	AuthStatus string `schema:"auth_status"`
	Invoice    string `schema:"invoice"`

	//Payment amount
	Currency string  `schema:"mc_currency"`
	Fee      float64 `schema:"mc_fee"`
	Gross    float64 `schema:"mc_gross"`

	PaymentDate   Time          `schema:"payment_date"`
	PaymentStatus PaymentStatus `schema:"payment_status"`
	PaymentType   PaymentType   `schema:"payment_type"`
	PendingReason PendingReason `schema:"pending_reason"`

	//ReasonCode is populated if the payment is negative
	ReasonCode string `schema:"reason_code"`

	Memo string `schema:"memo"`
}
```

Notification is sent from PayPal to our application. See
https://developer.paypal.com/docs/classic/ipn/integration-guide/IPNandPDTVariables
for more info

#### func  ReadNotification

```go
func ReadNotification(vals url.Values) *Notification
```
ReadNotification reads a notification from an //IPN request

#### func (*Notification) CustomerInfo

```go
func (n *Notification) CustomerInfo() string
```

#### func (*Notification) JSON

```go
func (n *Notification) JSON() (string, error)
```
CustomerInfo returns a nicely formatted customer info string

#### type PaymentStatus

```go
type PaymentStatus string
```

PaymentStatus represents the status of a payment

```go
var (
	PaymentStatusCanceledReversal PaymentStatus = "Canceled_Reversal"
	PaymentStatusCompleted        PaymentStatus = "Completed"
	PaymentStatusCreated          PaymentStatus = "Created"
	PaymentStatusDenied           PaymentStatus = "Denied"
	PaymentStatusExpired          PaymentStatus = "Expired"
	PaymentStatusFailed           PaymentStatus = "Failed"
	PaymentStatusPending          PaymentStatus = "Pending"
	PaymentStatusReversed         PaymentStatus = "Reversed"
	PaymentStatusProcessed        PaymentStatus = "Processed"
	PaymentStatusVoided           PaymentStatus = "Voided"
)
```
Payment statuses

#### type PaymentType

```go
type PaymentType string
```

PaymentType represents the type of a payment

```go
var (
	PaymentTypeEcheck  PaymentType = "echeck"
	PaymentTypeInstant PaymentType = "instant"
)
```
Payment Types

#### type PendingReason

```go
type PendingReason string
```

PendingReason represents the reason the payment is pending

```go
var (
	PendingReasonAddress             PendingReason = "address"
	PendingReasonAuthorization       PendingReason = "authorization"
	PendingReasonDelayedDisbursement PendingReason = "delayed_disbursement"
	PendingReasonEcheck              PendingReason = "echeck"
	PendingResasonIntl               PendingReason = "intl"
	PendingReasonMultiCurrency       PendingReason = "multi_currency"
	PendingReasonOrder               PendingReason = "order"
	PendingReasonPaymentReview       PendingReason = "paymentreview"
	PendingReasonRegulatoryReview    PendingReason = "regulatory_review"
	PendingReasonUnilateral          PendingReason = "unilateral"
	PendingReasonUpgrade             PendingReason = "upgrade"
	PendingReasonVerify              PendingReason = "verify"
	PendingReasonOther               PendingReason = "other"
)
```
Pending reasons

#### type Time

```go
type Time struct {
	Time *time.Time
}
```


#### func (*Time) UnmarshalText

```go
func (t *Time) UnmarshalText(text []byte) (err error)
```
