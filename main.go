package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Adidas White T-Shirt",
				Price: 9500.00,
				Stock: 15,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("[ERROR] Marshall Error : ", err)
	}

	fmt.Println(data)

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("[ERROR] Unmarshall Error : ", err)
	}

	for _, product := range testProducts.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetPrice())
		fmt.Println(product.GetStock())
		fmt.Println(product.GetCategory().GetName())
	}
}
