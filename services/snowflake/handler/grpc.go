package handler

import (
	"context"

	"github.com/vhall1/shorturl/lib/grpc/snowflake"
	"github.com/vhall1/shorturl/services/snowflake/types"
	"google.golang.org/grpc"
)

type SnowflakeGrpcHandler struct {
	snowflakeService types.Snowflake
	snowflake.UnimplementedSnowflakeServer
}

func NewSnowflakeGrpcHandler(grpc *grpc.Server, snowflakeService types.Snowflake) {
	h := &SnowflakeGrpcHandler{snowflakeService: snowflakeService}
	snowflake.RegisterSnowflakeServer(grpc, h)
}

func (h *SnowflakeGrpcHandler) HandleGetSnowflake(ctx context.Context, req *snowflake.GetSnowflakeRequest) *snowflake.GetSnowflakeResponse {
	id := h.snowflakeService.Generate()
	res := &snowflake.GetSnowflakeResponse{Id: id}
	return res
}
