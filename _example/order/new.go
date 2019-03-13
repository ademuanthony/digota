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
	"fmt"
	"github.com/digota/digota/order/orderpb"
	"github.com/digota/digota/payment/paymentpb"
	"github.com/digota/digota/sdk"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {

	c, err := sdk.NewClient("localhost:8082", &sdk.ClientOpt{
		InsecureSkipVerify: true,
		ServerName:         "server.merryworld.org",
		CaCrt:              "cert/out/ca.merryworld.org.crt",
		Crt:                "cert/out/client.merryworld.org.crt",
		Key:                "cert/out/client.merryworld.org.key",
	})

	if err != nil {
		dir, _ := os.Getwd()
		fmt.Println(dir)
		panic(err)
	}

	defer c.Close()

	// Create new order
	o, err := orderpb.NewOrderServiceClient(c).New(context.Background(), &orderpb.NewRequest{
		Currency: paymentpb.Currency_NGN,
		Items: []*orderpb.OrderItem{
			{
				Parent:   "d3854c1a-628a-4988-985e-682085e7256c",
				Quantity: 2,
				Type:     orderpb.OrderItem_Sku,
			},
			{
				Parent:   "80e1b81e-e9d9-4208-9792-42e47168f0f5",
				Quantity: 2,
				Type:     orderpb.OrderItem_Sku,
			},
			//{
			//	Parent:   "480e53bf-b409-4a34-8c74-13786b35ae11",
			//	Quantity: 1,
			//	Type:     orderpb.OrderItem_sku,
			//},
			//{
			//	Parent:   "480e53bf-b409-4a34-8c74-13786b35ae11",
			//	Quantity: 1,
			//	Type:     orderpb.OrderItem_sku,
			//},
			{
				Amount:      -1000,
				Description: "on the fly discount without parent",
				Currency:    paymentpb.Currency_NGN,
				Type:        orderpb.OrderItem_Discount,
			},
			{
				Amount:      1000,
				Description: "Tax (Included)",
				Currency:    paymentpb.Currency_NGN,
				Type:        orderpb.OrderItem_Tax,
			},
		},
		Email: "yaron@digota.com",
		UserId:"root",
		Shipping: &orderpb.Shipping{
			Name:  "Yaron Sumel",
			Phone: "+972 000 000 000",
			Address: &orderpb.Shipping_Address{
				Line1:      "Loren ipsum",
				City:       "San Jose",
				Country:    "USA",
				Line2:      "",
				PostalCode: "12345",
				State:      "CA",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println(o.GetId())

}
