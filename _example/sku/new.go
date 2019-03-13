// Digota <http://digota.com> - eCommerce microservice
// Copyright (c) 2018 Yaron Sumel <yaron@digota.com>
//
// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/digota/digota/payment/paymentpb"
	"github.com/digota/digota/sdk"
	"github.com/digota/digota/sku/skupb"
	"github.com/icrowley/fake"
	"golang.org/x/net/context"
)

func main() {

	c, err := sdk.NewClient("localhost:8082", &sdk.ClientOpt{
		InsecureSkipVerify: true,
		ServerName:         "server.merryworld.org",
		CaCrt:              "../../cert/out/ca.merryworld.org.crt",
		Crt:                "../../cert/out/client.merryworld.org.crt",
		Key:                "../../cert/out/client.merryworld.org.key",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	rand.Seed(time.Now().UnixNano())

	// Charge amount
	log.Println(skupb.NewSkuServiceClient(c).New(context.Background(), &skupb.NewRequest{
		Name:      fake.Brand(),
		Active:    true,
		Price:     uint64(rand.Int31n(10001)),
		Currency:  paymentpb.Currency_EUR,
		Parent: "708208c9-efe5-4ff7-bf70-4f888a80126d",
		//Metadata: map[string]string{
		//	"key": "val",
		//},
		Image: "http://sadf.124.com",
		PackageDimensions: &skupb.PackageDimensions{
			Weight: 1.1,
			Length: 1.2,
			Height: 1.3,
			Width:  1.4,
		},
		Inventory: &skupb.Inventory{
			Quantity: 1111,
			Type:     skupb.Inventory_Finite,
		},
		Attributes: map[string]string{
			"color": "red",
		},
	}))

}
