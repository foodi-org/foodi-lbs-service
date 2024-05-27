// Code generated by goctl. DO NOT EDIT.
// Source: foodiLBS.proto

package geo

import (
	"context"

	"github.com/foodi-org/foodi-lbs-server/github.com/foodi-org/foodi-lbs-server"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeliverDemoReply        = foodi_lbs_server.DeliverDemoReply
	DeliverDemoRequest      = foodi_lbs_server.DeliverDemoRequest
	DistReply               = foodi_lbs_server.DistReply
	DistRequest             = foodi_lbs_server.DistRequest
	GeoAddReply             = foodi_lbs_server.GeoAddReply
	GeoAddRequest           = foodi_lbs_server.GeoAddRequest
	GeoLocation             = foodi_lbs_server.GeoLocation
	HashReply               = foodi_lbs_server.HashReply
	HashRequest             = foodi_lbs_server.HashRequest
	Position                = foodi_lbs_server.Position
	PositionReply           = foodi_lbs_server.PositionReply
	PositionRequest         = foodi_lbs_server.PositionRequest
	RadioMemberRequest      = foodi_lbs_server.RadioMemberRequest
	RadioMemberStoreRequest = foodi_lbs_server.RadioMemberStoreRequest
	RadioOrderOption        = foodi_lbs_server.RadioOrderOption
	RadioOrderReply         = foodi_lbs_server.RadioOrderReply
	RadioOrderRequest       = foodi_lbs_server.RadioOrderRequest
	RadioOrderStoreRequest  = foodi_lbs_server.RadioOrderStoreRequest

	Geo interface {
		GeoAdd(ctx context.Context, in *GeoAddRequest, opts ...grpc.CallOption) (*GeoAddReply, error)
		GeoPosition(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReply, error)
		GeoHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashReply, error)
		GeoDist(ctx context.Context, in *DistRequest, opts ...grpc.CallOption) (*DistReply, error)
		RadioOrder(ctx context.Context, in *RadioOrderRequest, opts ...grpc.CallOption) (*RadioOrderReply, error)
		RadioOrderStore(ctx context.Context, in *RadioOrderStoreRequest, opts ...grpc.CallOption) (*RadioOrderReply, error)
		RadioMember(ctx context.Context, in *RadioMemberRequest, opts ...grpc.CallOption) (*RadioOrderReply, error)
		RadioMemberStore(ctx context.Context, in *RadioMemberStoreRequest, opts ...grpc.CallOption) (*RadioOrderReply, error)
	}

	defaultGeo struct {
		cli zrpc.Client
	}
)

func NewGeo(cli zrpc.Client) Geo {
	return &defaultGeo{
		cli: cli,
	}
}

func (m *defaultGeo) GeoAdd(ctx context.Context, in *GeoAddRequest, opts ...grpc.CallOption) (*GeoAddReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.GeoAdd(ctx, in, opts...)
}

func (m *defaultGeo) GeoPosition(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.GeoPosition(ctx, in, opts...)
}

func (m *defaultGeo) GeoHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.GeoHash(ctx, in, opts...)
}

func (m *defaultGeo) GeoDist(ctx context.Context, in *DistRequest, opts ...grpc.CallOption) (*DistReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.GeoDist(ctx, in, opts...)
}

func (m *defaultGeo) RadioOrder(ctx context.Context, in *RadioOrderRequest, opts ...grpc.CallOption) (*RadioOrderReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.RadioOrder(ctx, in, opts...)
}

func (m *defaultGeo) RadioOrderStore(ctx context.Context, in *RadioOrderStoreRequest, opts ...grpc.CallOption) (*RadioOrderReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.RadioOrderStore(ctx, in, opts...)
}

func (m *defaultGeo) RadioMember(ctx context.Context, in *RadioMemberRequest, opts ...grpc.CallOption) (*RadioOrderReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.RadioMember(ctx, in, opts...)
}

func (m *defaultGeo) RadioMemberStore(ctx context.Context, in *RadioMemberStoreRequest, opts ...grpc.CallOption) (*RadioOrderReply, error) {
	client := foodi_lbs_server.NewGeoClient(m.cli.Conn())
	return client.RadioMemberStore(ctx, in, opts...)
}
