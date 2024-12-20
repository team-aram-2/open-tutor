//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	GitHubOAuthScopes = "GitHubOAuth.Scopes"
	GoogleOAuthScopes = "GoogleOAuth.Scopes"
)

// Defines values for RatingRatingType.
const (
	RatingRatingTypeStudent RatingRatingType = "student"
	RatingRatingTypeTutor   RatingRatingType = "tutor"
)

// Defines values for GetRatingByIdParamsUserType.
const (
	GetRatingByIdParamsUserTypeStudent GetRatingByIdParamsUserType = "student"
	GetRatingByIdParamsUserTypeTutor   GetRatingByIdParamsUserType = "tutor"
)

// ErrorModel defines model for ErrorModel.
type ErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Meeting defines model for Meeting.
type Meeting struct {
	EndAt     *time.Time          `json:"endAt,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	StartAt   *time.Time          `json:"startAt,omitempty"`
	StudentId *openapi_types.UUID `json:"studentId,omitempty"`
	TutorId   *openapi_types.UUID `json:"tutorId,omitempty"`
}

// Message defines model for Message.
type Message struct {
	// MessageAttachments Array of message attachments
	MessageAttachments *[]MessageAttachment `json:"MessageAttachments,omitempty"`

	// Id Unique identifier for the message.
	Id openapi_types.UUID `json:"id"`

	// Message Message content.
	Message string `json:"message"`

	// OriginId Unique identifier for the originID for the message.
	OriginId openapi_types.UUID `json:"originId"`

	// RecipientId Unique identifier for the recipient of the message.
	RecipientId openapi_types.UUID `json:"recipientId"`
	SentOn      time.Time          `json:"sentOn"`
}

// MessageAttachment defines model for MessageAttachment.
type MessageAttachment struct {
	// Filename Name of the attachment.
	Filename string `json:"filename"`

	// Id Unique identifier for the message the attachment belongs to.
	Id openapi_types.UUID `json:"id"`

	// Mimetype Mimetype of the attachment.
	Mimetype string `json:"mimetype"`

	// Url Source of the attachment.
	Url string `json:"url"`
}

// Rating defines model for Rating.
type Rating struct {
	Comment    *string             `json:"comment,omitempty"`
	Id         *openapi_types.UUID `json:"id,omitempty"`
	MeetingId  *openapi_types.UUID `json:"meetingId,omitempty"`
	RatingType *RatingRatingType   `json:"ratingType,omitempty"`
	Ratings    *struct {
		Communication   int `json:"communication"`
		Knowledge       int `json:"knowledge"`
		Overall         int `json:"overall"`
		Professionalism int `json:"professionalism"`
		Punctuality     int `json:"punctuality"`
	} `json:"ratings,omitempty"`
	ReviewerUserId *openapi_types.UUID `json:"reviewerUserId,omitempty"`
	UserId         *openapi_types.UUID `json:"userId,omitempty"`
}

// RatingRatingType defines model for Rating.RatingType.
type RatingRatingType string

// Tutor defines model for Tutor.
type Tutor struct {
	AccountLocked *bool               `json:"accountLocked,omitempty"`
	Email         openapi_types.Email `json:"email"`
	FirstName     string              `json:"firstName"`
	LastName      string              `json:"lastName"`
	SignedUpAt    *time.Time          `json:"signedUpAt,omitempty"`
	Skills        *[]string           `json:"skills,omitempty"`
	TotalHours    *int                `json:"totalHours,omitempty"`
	UserId        openapi_types.UUID  `json:"userId"`
}

// User Base User object containing shared details needed for all users.
type User struct {
	AccountLocked *bool               `json:"accountLocked,omitempty"`
	Email         openapi_types.Email `json:"email"`
	FirstName     string              `json:"firstName"`
	LastName      string              `json:"lastName"`
	SignedUpAt    *time.Time          `json:"signedUpAt,omitempty"`
	UserId        openapi_types.UUID  `json:"userId"`
}

// GetRatingByIdParams defines parameters for GetRatingById.
type GetRatingByIdParams struct {
	UserType *GetRatingByIdParamsUserType `form:"userType,omitempty" json:"userType,omitempty"`
}

// GetRatingByIdParamsUserType defines parameters for GetRatingById.
type GetRatingByIdParamsUserType string

// CreateMeetingJSONRequestBody defines body for CreateMeeting for application/json ContentType.
type CreateMeetingJSONRequestBody = Meeting

// UpdateMeetingByIdJSONRequestBody defines body for UpdateMeetingById for application/json ContentType.
type UpdateMeetingByIdJSONRequestBody = Meeting

// CreateMessageJSONRequestBody defines body for CreateMessage for application/json ContentType.
type CreateMessageJSONRequestBody = Message

// UpdateMessageByIdJSONRequestBody defines body for UpdateMessageById for application/json ContentType.
type UpdateMessageByIdJSONRequestBody = Message

// CreateMessageAttachmentJSONRequestBody defines body for CreateMessageAttachment for application/json ContentType.
type CreateMessageAttachmentJSONRequestBody = MessageAttachment

// PostRatingJSONRequestBody defines body for PostRating for application/json ContentType.
type PostRatingJSONRequestBody = Rating

// SignUpAsTutorJSONRequestBody defines body for SignUpAsTutor for application/json ContentType.
type SignUpAsTutorJSONRequestBody = Tutor

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// UpdateUserByIdJSONRequestBody defines body for UpdateUserById for application/json ContentType.
type UpdateUserByIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new meeting
	// (POST /meeting)
	CreateMeeting(w http.ResponseWriter, r *http.Request)
	// Delete a meeting by ID
	// (DELETE /meeting/{meetingId})
	DeleteMeetingById(w http.ResponseWriter, r *http.Request, meetingId interface{})
	// Get a meeting by ID
	// (GET /meeting/{meetingId})
	GetMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID)
	// Update meeting information
	// (PUT /meeting/{meetingId})
	UpdateMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID)
	// Creates a new message
	// (POST /message)
	CreateMessage(w http.ResponseWriter, r *http.Request)
	// Delete a message by id
	// (DELETE /message/{messageId})
	DeleteMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID)
	// Get a message by id
	// (GET /message/{messageId})
	GetMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID)
	// Update a message
	// (PUT /message/{messageId})
	UpdateMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID)
	// Create a new message attachment.
	// (POST /messageAttachment)
	CreateMessageAttachment(w http.ResponseWriter, r *http.Request)
	// Delete an attachment by ID
	// (DELETE /messageAttachment/{messageAttachmentId})
	DeleteMessageAttachmentById(w http.ResponseWriter, r *http.Request, messageAttachmentId openapi_types.UUID)
	// Get an attachment by ID
	// (GET /messageAttachment/{messageAttachmentId})
	GetMessageAttachmentById(w http.ResponseWriter, r *http.Request, messageAttachmentId openapi_types.UUID)
	// Post a rating
	// (POST /rating)
	PostRating(w http.ResponseWriter, r *http.Request)
	// Get a user's rating by user ID, optionally filtering by usertype.
	// (GET /rating/{userId})
	GetRatingById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID, params GetRatingByIdParams)
	// Create Tutor Profile for User
	// (POST /tutor)
	SignUpAsTutor(w http.ResponseWriter, r *http.Request)
	// Get a tutor by ID
	// (GET /tutor/{tutorId})
	GetTutorById(w http.ResponseWriter, r *http.Request, tutorId openapi_types.UUID)
	// Create a new user
	// (POST /user)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Delete user account, maybe via settings or moderation panel
	// (DELETE /user/{userId})
	DeleteUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID)
	// Get a user by ID
	// (GET /user/{userId})
	GetUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID)
	// Update user information
	// (PUT /user/{userId})
	UpdateUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CreateMeeting operation middleware
func (siw *ServerInterfaceWrapper) CreateMeeting(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateMeeting(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteMeetingById operation middleware
func (siw *ServerInterfaceWrapper) DeleteMeetingById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "meetingId" -------------
	var meetingId interface{}

	err = runtime.BindStyledParameterWithOptions("simple", "meetingId", r.PathValue("meetingId"), &meetingId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "meetingId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteMeetingById(w, r, meetingId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMeetingById operation middleware
func (siw *ServerInterfaceWrapper) GetMeetingById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "meetingId" -------------
	var meetingId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "meetingId", r.PathValue("meetingId"), &meetingId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "meetingId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMeetingById(w, r, meetingId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateMeetingById operation middleware
func (siw *ServerInterfaceWrapper) UpdateMeetingById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "meetingId" -------------
	var meetingId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "meetingId", r.PathValue("meetingId"), &meetingId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "meetingId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateMeetingById(w, r, meetingId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateMessage operation middleware
func (siw *ServerInterfaceWrapper) CreateMessage(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateMessage(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteMessageById operation middleware
func (siw *ServerInterfaceWrapper) DeleteMessageById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "messageId" -------------
	var messageId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "messageId", r.PathValue("messageId"), &messageId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "messageId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteMessageById(w, r, messageId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMessageById operation middleware
func (siw *ServerInterfaceWrapper) GetMessageById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "messageId" -------------
	var messageId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "messageId", r.PathValue("messageId"), &messageId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "messageId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMessageById(w, r, messageId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateMessageById operation middleware
func (siw *ServerInterfaceWrapper) UpdateMessageById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "messageId" -------------
	var messageId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "messageId", r.PathValue("messageId"), &messageId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "messageId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateMessageById(w, r, messageId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateMessageAttachment operation middleware
func (siw *ServerInterfaceWrapper) CreateMessageAttachment(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateMessageAttachment(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteMessageAttachmentById operation middleware
func (siw *ServerInterfaceWrapper) DeleteMessageAttachmentById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "messageAttachmentId" -------------
	var messageAttachmentId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "messageAttachmentId", r.PathValue("messageAttachmentId"), &messageAttachmentId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "messageAttachmentId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteMessageAttachmentById(w, r, messageAttachmentId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMessageAttachmentById operation middleware
func (siw *ServerInterfaceWrapper) GetMessageAttachmentById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "messageAttachmentId" -------------
	var messageAttachmentId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "messageAttachmentId", r.PathValue("messageAttachmentId"), &messageAttachmentId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "messageAttachmentId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMessageAttachmentById(w, r, messageAttachmentId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostRating operation middleware
func (siw *ServerInterfaceWrapper) PostRating(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRating(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRatingById operation middleware
func (siw *ServerInterfaceWrapper) GetRatingById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "userId", r.PathValue("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRatingByIdParams

	// ------------- Optional query parameter "userType" -------------

	err = runtime.BindQueryParameter("form", true, false, "userType", r.URL.Query(), &params.UserType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userType", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRatingById(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// SignUpAsTutor operation middleware
func (siw *ServerInterfaceWrapper) SignUpAsTutor(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SignUpAsTutor(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetTutorById operation middleware
func (siw *ServerInterfaceWrapper) GetTutorById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tutorId" -------------
	var tutorId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "tutorId", r.PathValue("tutorId"), &tutorId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tutorId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTutorById(w, r, tutorId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteUserById operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "userId", r.PathValue("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserById operation middleware
func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "userId", r.PathValue("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateUserById operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "userId", r.PathValue("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/meeting", wrapper.CreateMeeting)
	m.HandleFunc("DELETE "+options.BaseURL+"/meeting/{meetingId}", wrapper.DeleteMeetingById)
	m.HandleFunc("GET "+options.BaseURL+"/meeting/{meetingId}", wrapper.GetMeetingById)
	m.HandleFunc("PUT "+options.BaseURL+"/meeting/{meetingId}", wrapper.UpdateMeetingById)
	m.HandleFunc("POST "+options.BaseURL+"/message", wrapper.CreateMessage)
	m.HandleFunc("DELETE "+options.BaseURL+"/message/{messageId}", wrapper.DeleteMessageById)
	m.HandleFunc("GET "+options.BaseURL+"/message/{messageId}", wrapper.GetMessageById)
	m.HandleFunc("PUT "+options.BaseURL+"/message/{messageId}", wrapper.UpdateMessageById)
	m.HandleFunc("POST "+options.BaseURL+"/messageAttachment", wrapper.CreateMessageAttachment)
	m.HandleFunc("DELETE "+options.BaseURL+"/messageAttachment/{messageAttachmentId}", wrapper.DeleteMessageAttachmentById)
	m.HandleFunc("GET "+options.BaseURL+"/messageAttachment/{messageAttachmentId}", wrapper.GetMessageAttachmentById)
	m.HandleFunc("POST "+options.BaseURL+"/rating", wrapper.PostRating)
	m.HandleFunc("GET "+options.BaseURL+"/rating/{userId}", wrapper.GetRatingById)
	m.HandleFunc("POST "+options.BaseURL+"/tutor", wrapper.SignUpAsTutor)
	m.HandleFunc("GET "+options.BaseURL+"/tutor/{tutorId}", wrapper.GetTutorById)
	m.HandleFunc("POST "+options.BaseURL+"/user", wrapper.CreateUser)
	m.HandleFunc("DELETE "+options.BaseURL+"/user/{userId}", wrapper.DeleteUserById)
	m.HandleFunc("GET "+options.BaseURL+"/user/{userId}", wrapper.GetUserById)
	m.HandleFunc("PUT "+options.BaseURL+"/user/{userId}", wrapper.UpdateUserById)

	return m
}
