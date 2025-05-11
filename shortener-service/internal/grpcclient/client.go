package grpcclient

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"shortener-service/proto"
	"time"
)

type AnalyticsClient struct {
	client proto.AnalyticsClient
}

func NewAnalyticsClient(addr string) (*AnalyticsClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &AnalyticsClient{client: proto.NewAnalyticsClient(conn)}, nil
}

func (c *AnalyticsClient) SendVisit(shortcode, ip, userAgent string, url string) {
	var err error
	_, err = c.client.LogVisit(context.Background(), &proto.VisitRequest{
		Shortcode: shortcode,
		Url:       url,
		Ip:        ip,
		UserAgent: userAgent,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		log.Println(err)
	}
}
