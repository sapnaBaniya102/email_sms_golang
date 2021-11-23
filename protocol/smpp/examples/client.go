package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/smpp"
	"github.com/sujit-baniya/smpp/pdu"
)

func main() {


	app.Post("/sms/send", func(c *fiber.Ctx) error {
		var message smpp.Message
		c.BodyParser(&message)
		go func() {
			for i := 0; i <= 100; i++ {
				manager.Send(message)
			}
		}()

		return c.JSON(fiber.Map{
			"success": true,
			"message": "We will notify you once all messages are sent",
		})
	})

	log.Fatalln(app.Listen(":8080"))
}

var handlePDU = func(conn *smpp.Conn) {
	for {
		packet := <-conn.PDU()
		switch pd := packet.(type) {
		case *pdu.DeliverSM:
			fmt.Println(pd)
			err := conn.Send(pd.Resp())
			if err != nil {
				fmt.Println(err)
			}
		case pdu.Responsable:
			fmt.Println(pd)
			err := conn.Send(pd.Resp())
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
var manager *smpp.Manager

func startSmsManager() {
	manager, err := smpp.NewManager(smpp.Setting{
		URL: "smscsim.melroselabs.com:2775",
		Auth: smpp.Auth{
			SystemID:   "426388",
			Password:   "f23d60",
			SystemType: "",
		},
		SmppVersion:      pdu.SMPPVersion50,
		ReadTimeout:      10 * time.Second,
		WriteTimeout:     10 * time.Second,
		EnquiryInterval:  time.Minute,
		EnquiryTimeout:   time.Minute,
		MaxConnection:    2,
		UseAllConnection: true,
		Throttle:         100,
		HandlePDU:        handlePDU,
	})
	if err != nil {
		panic(err)
	}
	manager.Start()
	err = manager.HandlePDU()
	if err != nil {
		panic(err)
	}
}
