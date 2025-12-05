package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()

	fmt.Println("Requesting drinks 🍹🍺☕")
	res, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	fmt.Println("Available drinks:")
	for _, d := range res.Drinks {
		if d.Price == 0 {
			fmt.Printf("\t> id:%d  name:\"%s\"  description:\"%s\"\n", d.Id, d.Name, d.Description)
		} else {
			fmt.Printf("\t> id:%d  name:\"%s\"  price:%.0f  description:\"%s\"\n", d.Id, d.Name, d.Price, d.Description)
		}
	}

	orderAndPrint := func(id int32, name string, count int32) {
		fmt.Printf("\t> Ordering: %d x %s\n", count, name)
		_, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{DrinkId: id, Count: count})
	}

	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	orderAndPrint(1, "Spritzer", 2)
	orderAndPrint(2, "Beer", 2)
	orderAndPrint(3, "Coffee", 2)

	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	orderAndPrint(1, "Spritzer", 6)
	orderAndPrint(2, "Beer", 6)
	orderAndPrint(3, "Coffee", 6)

	fmt.Println("Getting the bill 💹💹💹")
	ordersRes, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	for _, item := range ordersRes.Orders {
		fmt.Printf("\t> Total: %d x %s\n", item.Amount, item.Drink.Name)
	}

	return nil
}