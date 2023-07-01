package tracing

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func PropagateRequestIDToContext(ctx context.Context) context.Context {
	requestMetadata := FromContext(ctx)
	if requestMetadata == nil {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, DefaultContextKeyRequestID, requestMetadata.RequestID)
}

func PropagateRequestIDToHeader(ctx context.Context, outGoingHeader *http.Header) {
	if ginCtx, ok := ctx.(*gin.Context); ok && ginCtx != nil && ginCtx.Request != nil && ginCtx.Request.Context() != nil {
		ctx = ginCtx.Request.Context()
	}
	requestMetadata := FromContext(ctx)
	if requestMetadata == nil {
		return
	}
	outGoingHeader.Set(DefaultContextKeyRequestID, requestMetadata.RequestID)
}
