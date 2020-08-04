package themap

import (
	//"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGooglePay(t *testing.T) {

	payload := []byte(`{
  "protocolVersion":"ECv2",
  "signature":"MEQCIH6Q4OwQ0jAceFEkGF0JID6sJNXxOEi4r+mA7biRxqBQAiAondqoUpU/bdsrAOpZIsrHQS9nwiiNwOrr24RyPeHA0Q\u003d\u003d",
  "intermediateSigningKey":{
    "signedKey": "{\"keyExpiration\":\"1542323393147\",\"keyValue\":\"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/1+3HBVSbdv+j7NaArdgMyoSAM43yRydzqdg1TxodSzA96Dj4Mc1EiKroxxunavVIvdxGnJeFViTzFvzFRxyCw\\u003d\\u003d\"}",
    "signatures": ["MEYCIQCO2EIi48s8VTH+ilMEpoXLFfkxAwHjfPSCVED/QDSHmQIhALLJmrUlNAY8hDQRV/y1iKZGsWpeNmIP+z+tCQHQxP0v"]
  },
  "signedMessage":"{\"tag\":\"jpGz1F1Bcoi/fCNxI9n7Qrsw7i7KHrGtTf3NrRclt+U\\u003d\",\"ephemeralPublicKey\":\"BJatyFvFPPD21l8/uLP46Ta1hsKHndf8Z+tAgk+DEPQgYTkhHy19cF3h/bXs0tWTmZtnNm+vlVrKbRU9K8+7cZs\\u003d\",\"encryptedMessage\":\"mKOoXwi8OavZ\"}"
}`)

	reply := `{
    "Success": true,
    "OrderId": "GooglePayTestOrder-001",
    "Amount": 300,
    "ErrCode": ""
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/Pay",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetTerm("123")
	err := trans.GooglePay(300, payload)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

}

//The ExampleGooglePay method executes test GooglePay payment
// func ExampleGooglePay() {
//
// 	payload := []byte(`{
//   "protocolVersion":"ECv2",
//   "signature":"MEQCIH6Q4OwQ0jAceFEkGF0JID6sJNXxOEi4r+mA7biRxqBQAiAondqoUpU/bdsrAOpZIsrHQS9nwiiNwOrr24RyPeHA0Q\u003d\u003d",
//   "intermediateSigningKey":{
//     "signedKey": "{\"keyExpiration\":\"1542323393147\",\"keyValue\":\"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/1+3HBVSbdv+j7NaArdgMyoSAM43yRydzqdg1TxodSzA96Dj4Mc1EiKroxxunavVIvdxGnJeFViTzFvzFRxyCw\\u003d\\u003d\"}",
//     "signatures": ["MEYCIQCO2EIi48s8VTH+ilMEpoXLFfkxAwHjfPSCVED/QDSHmQIhALLJmrUlNAY8hDQRV/y1iKZGsWpeNmIP+z+tCQHQxP0v"]
//   },
//   "signedMessage":"{\"tag\":\"jpGz1F1Bcoi/fCNxI9n7Qrsw7i7KHrGtTf3NrRclt+U\\u003d\",\"ephemeralPublicKey\":\"BJatyFvFPPD21l8/uLP46Ta1hsKHndf8Z+tAgk+DEPQgYTkhHy19cF3h/bXs0tWTmZtnNm+vlVrKbRU9K8+7cZs\\u003d\",\"encryptedMessage\":\"mKOoXwi8OavZ\"}"
// }`)
//
// 	pay := New("TestTerminal", "GooglePayTestOrder-001")
// 	err := pay.GooglePay(300, payload) // Create session for 3.00RUB
// 	if err != nil {
// 		fmt.Printf("Error occurred: %v", err)
// 	}
//
// 	fmt.Printf("%v", pay.Success)
// 	// Output: true
// }
