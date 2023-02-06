package main

import (
	"context"
	video "tiktok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, createVideoReq *video.CreateVideoReq) (resp *video.CreateVideoResp, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, getPublishListReq *video.GetPublishListReq) (resp *video.GetPublishListResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, getVideosReq *video.GetVideosReq) (resp *video.GetVideosResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideoById implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoById(ctx context.Context, getVideoByIdReq *video.GetVideoByIdReq) (resp *video.GetVideoByIdResp, err error) {
	// TODO: Your code here...
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, publishReq *video.PublishReq) (resp *video.PublishResp, err error) {
	// TODO: Your code here...
	return
}
