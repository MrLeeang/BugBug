package controller

import (
	"BugBug/controller/views"

	"github.com/gin-gonic/gin"
)

// MakeRouter 是路由
func MakeRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	// 用户信息
	v1.GET("/user/:userId/info", views.ActionUserInfo)
	// 用户登录
	v1.POST("/user/login", views.ActionUserLogin)
	// 用户列表
	v1.GET("/user", views.AuthHandler(), views.ActionUserList)
	// 修改用户信息
	v1.PUT("/user/update", views.AuthHandler(), views.ActionUpdateUserInfo)
	// 发表评论
	v1.POST("/comment/create", views.AuthHandler(), views.ActionCreateComment)
	// 删除评论
	v1.DELETE("/comment/:postCommentID/delete", views.AuthHandler(), views.ActionDeleteComment)
	// 帖子评论列表
	v1.GET("/post/:postID/comments", views.AuthHandler(), views.ActionPostCommentList)
	// 评论得回复列表
	v1.GET("/comment/:postCommentID/replys", views.AuthHandler(), views.ActionCommentReplyList)
	// 点赞
	v1.POST("/vote/post", views.AuthHandler(), views.ActionVotePost)
	// 取消点赞
	v1.DELETE("/vote/cancel", views.AuthHandler(), views.ActionVoteCancel)
	// 采纳
	v1.POST("/adopt/post", views.AuthHandler(), views.ActionAdoptPost)
	// 用户采纳的帖子
	v1.GET("/adopt/:userId/posts", views.AuthHandler(), views.ActionUserAdoptList)
	// 上传token
	v1.GET("/upload/token", views.AuthHandler(), views.ActionUploadToken)

	// 发布帖子
	v1.POST("/post/publish", views.AuthHandler(), views.ActionPostPublish)
	// 用户的帖子列表
	v1.GET("/user/:userId/posts", views.ActionUserPostList)
	// 推荐的帖子列表
	v1.GET("/post/:postID", views.ActionRecommendPostList)
	// 帖子详情
	v1.GET("/post/:postID/info", views.ActionPostInfo)

	// messages
	// 用户帖子点赞列表
	v1.GET("/information/get_vote_list", views.AuthHandler(), views.GetVoteListByUserPost)
	// 用户帖子采纳列表
	v1.GET("/information/get_adopt_list", views.AuthHandler(), views.GetVoteListByUserPost)
	// 我得采纳列表 get_adopt_list
	v1.GET("/user/:userId", views.AuthHandler(), views.GetAdoptPostListByUser)
	// 我的评论列表
	v1.GET("/vote/msgcomm", views.AuthHandler(), views.ActionCommentListByUser)
	// 未读消息 redis

	// 暂未开发的功能
	// 关注
	v1.POST("/user/follow/:userId", views.AuthHandler(), views.ActionAddUserFollow)
	// 取关
	v1.DELETE("/user/follow/:userId", views.AuthHandler(), views.ActionDelUserFollow)
	// 关注列表
	v1.GET("/user/:userId/follows", views.ActionUserFollowList)
	// 粉丝列表
	v1.GET("/user/:userId/fans", views.ActionUserFansList)
	// 删除粉丝
	v1.DELETE("/user/fans/:userId", views.AuthHandler(), views.ActionDelUserFans)
}
