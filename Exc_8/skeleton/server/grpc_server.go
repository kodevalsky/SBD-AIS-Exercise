package server

import (
	"context"
	"exc8/pb"
	"net"
	"sync"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
}

func StartGrpcServer() error {
	srv := grpc.NewServer()
	grpcService := &GRPCService{}
	pb.RegisterOrderServiceServer(srv, grpcService)
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

var (
	drinksStore = map[int32]*pb.Drink{
		1: {Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
		2: {Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
		3: {Id: 3, Name: "Coffee", Price: 0, Description: "Mifare isn't that secure"},
	}
	orderStore = make(map[int32]int32)
	mu         sync.Mutex
)

func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.GetDrinksResponse, error) {
	var list []*pb.Drink
	for i := int32(1); i <= 3; i++ {
		if d, ok := drinksStore[i]; ok {
			list = append(list, d)
		}
	}
	return &pb.GetDrinksResponse{Drinks: list}, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*emptypb.Empty, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := drinksStore[req.DrinkId]; exists {
		orderStore[req.DrinkId] += req.Count
	}
	return &emptypb.Empty{}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	var items []*pb.OrderItem
	for i := int32(1); i <= 3; i++ {
		count := orderStore[i]
		if count > 0 {
			items = append(items, &pb.OrderItem{
				Drink:  drinksStore[i],
				Amount: count,
			})
		}
	}
	return &pb.GetOrdersResponse{Orders: items}, nil
}