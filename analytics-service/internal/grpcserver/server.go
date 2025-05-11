package grpcserver

import (
	"analytics-service/internal/storage"
	"analytics-service/proto"
	"context"
	"time"
)

type Server struct {
	proto.UnimplementedAnalyticsServer
	Store *storage.Storage
}

func (s *Server) LogVisit(ctx context.Context, req *proto.VisitRequest) (*proto.VisitResponse, error) {
	v := storage.Visit{
		Shortcode: req.Shortcode,
		IP:        req.Ip,
		UserAgent: req.UserAgent,
		VisitedAt: time.Unix(req.Timestamp, 0),
	}
	err := s.Store.SaveVisit(v)
	if err != nil {
		return nil, err
	}
	return &proto.VisitResponse{Status: "ok"}, nil
}
