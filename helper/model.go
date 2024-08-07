package helper

import (
	"database/sql"
	"strconv"
	"uaspw2/exception"
	"uaspw2/models/entity"
	"uaspw2/models/web/response"
)

func ToUserResponse(user entity.User) response.UserResponse {
	return response.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []entity.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserWithProfileResponse(user entity.UserWithProfile) response.UserWithProfileResponse {
	return response.UserWithProfileResponse{
		Id:        user.Id,
		Username:  user.Username,
		Profile:   user.Profile,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserWithProfileResponses(users []entity.UserWithProfile) []response.UserWithProfileResponse {
	var userResponses []response.UserWithProfileResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserWithProfileResponse(user))
	}
	return userResponses
}

func ToUserProfileResponse(userProfile entity.UserProfile) response.UserProfileResponse {
	return response.UserProfileResponse{
		UserId:      userProfile.UserId,
		FullName:    userProfile.FullName,
		Gender:      userProfile.Gender,
		BirthDate:   userProfile.BirthDate,
		PhoneNumber: userProfile.PhoneNumber,
		Address:     userProfile.Address,
		CreatedAt:   userProfile.CreatedAt,
		UpdatedAt:   userProfile.UpdatedAt,
	}
}

func ToUserProfileResponses(userProfiles []entity.UserProfile) []response.UserProfileResponse {
	var userProfileResponses []response.UserProfileResponse
	for _, userProfile := range userProfiles {
		userProfileResponses = append(userProfileResponses, ToUserProfileResponse(userProfile))
	}
	return userProfileResponses
}

func ToUserPhotoProfileResponse(userPhotoProfile entity.UserProfilePhoto) response.UserProfilePhotoResponse {
	return response.UserProfilePhotoResponse{
		UserId:    userPhotoProfile.UserId,
		Path:      userPhotoProfile.Path,
		CreatedAt: userPhotoProfile.CreatedAt,
		UpdatedAt: userPhotoProfile.UpdatedAt,
	}
}

func ToArticleResponse(article entity.Article) response.ArticleResponse {
	return response.ArticleResponse{
		Id:          article.Id,
		UserId:      article.UserId,
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		Author:      article.Author,
		Media:       article.Media,
		IsPublished: article.IsPublished,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
}

func ToArticleResponses(articles []entity.Article) []response.ArticleResponse {
	var articleResponses []response.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, ToArticleResponse(article))
	}
	return articleResponses
}

func ToLikeResponse(like entity.Like) response.LikeResponse {
	return response.LikeResponse{
		Id:        like.Id,
		UserId:    like.UserId,
		ArticleId: like.ArticleId,
		CreatedAt: like.CreatedAt,
		UpdatedAt: like.UpdatedAt,
	}
}

func ToLikeResponses(likes []entity.Like) []response.LikeResponse {
	var likeResponses []response.LikeResponse
	for _, like := range likes {
		likeResponses = append(likeResponses, ToLikeResponse(like))
	}
	return likeResponses
}

func ToCommentResponse(comment entity.Comment) response.CommentResponse {
	return response.CommentResponse{
		Id:        comment.Id,
		UserId:    comment.UserId,
		ArticleId: comment.ArticleId,
		Comment:   comment.Comment,
		Author:    comment.Author,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func ToCommentResponses(comments []entity.Comment) []response.CommentResponse {
	var commentResponses []response.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, ToCommentResponse(comment))
	}
	return commentResponses
}

func ToIntFromParams(params string) int {
	id, err := strconv.Atoi(params)
	if err != nil {
		panic(exception.NewInvalidParameter("Invalid parameter. Excepted a number but received string"))
	}
	return id
}

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
