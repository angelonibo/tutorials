// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/orders/orders.proto

package mu_micro_book_srv_orders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Orders service

type OrdersService interface {
	New(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type ordersService struct {
	c    client.Client
	name string
}

func NewOrdersService(name string, c client.Client) OrdersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mu.micro.book.srv.orders"
	}
	return &ordersService{
		c:    c,
		name: name,
	}
}

func (c *ordersService) New(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.New", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) GetOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.GetOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orders service

type OrdersHandler interface {
	New(context.Context, *Request, *Response) error
	GetOrder(context.Context, *Request, *Response) error
}

func RegisterOrdersHandler(s server.Server, hdlr OrdersHandler, opts ...server.HandlerOption) error {
	type orders interface {
		New(ctx context.Context, in *Request, out *Response) error
		GetOrder(ctx context.Context, in *Request, out *Response) error
	}
	type Orders struct {
		orders
	}
	h := &ordersHandler{hdlr}
	return s.Handle(s.NewHandler(&Orders{h}, opts...))
}

type ordersHandler struct {
	OrdersHandler
}

func (h *ordersHandler) New(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.New(ctx, in, out)
}

func (h *ordersHandler) GetOrder(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.GetOrder(ctx, in, out)
}
