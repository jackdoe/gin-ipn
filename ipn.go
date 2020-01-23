// gin friendly paypal ipn listener
// Example:
//  package main
//
//  import (
//  	"log"
//
//  	"github.com/gin-gonic/gin"
//  	"github.com/jackdoe/gin-ipn/ipn"
//  )
//
//  func main() {
//  	r := gin.Default()
//
//  	ipn.Listener(r, "/ipn/:paymentID", func(c *gin.Context, err error, body string, n *ipn.Notification) error {
//  		if err != nil {
//  			panic(err)
//  		}
//
//  		if n.TestIPN {
//  			log.Printf("test")
//  		}
//
//  		log.Printf("notification: %v", n)
//
//  		return nil
//  	})
//  }
package ipn
