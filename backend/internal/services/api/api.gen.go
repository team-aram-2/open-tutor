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

// CreateMeetingBody defines model for CreateMeetingBody.
type CreateMeetingBody struct {
	EndAt     *time.Time          `json:"endAt,omitempty"`
	StartAt   *time.Time          `json:"startAt,omitempty"`
	StudentId *openapi_types.UUID `json:"studentId,omitempty"`
}

// Conversation defines model for Conversation.
type Conversation struct {
	Id    *openapi_types.UUID   `json:"id,omitempty"`
	Users *[]openapi_types.UUID `json:"users,omitempty"`
}

// ErrorModel defines model for ErrorModel.
type ErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Meeting defines model for Meeting.
type Meeting struct {
	EndAt        time.Time          `json:"endAt"`
	Id           openapi_types.UUID `json:"id"`
	StartAt      time.Time          `json:"startAt"`
	StudentId    openapi_types.UUID `json:"studentId"`
	TutorId      openapi_types.UUID `json:"tutorId"`
	ZoomHostLink *string            `json:"zoomHostLink,omitempty"`
	ZoomJoinLink *string            `json:"zoomJoinLink,omitempty"`
}

// Message defines model for Message.
type Message struct {
	// MessageAttachments Array of message attachments
	MessageAttachments *[]MessageAttachment `json:"MessageAttachments,omitempty"`

	// ConversationId Unique identifier for the conversation for the message.
	ConversationId openapi_types.UUID `json:"conversationId"`

	// Id Unique identifier for the message.
	Id openapi_types.UUID `json:"id"`

	// Message Message content.
	Message string `json:"message"`

	// OriginId Unique identifier for the originID for the message.
	OriginId openapi_types.UUID `json:"originId"`
	SentOn   time.Time          `json:"sentOn"`
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

// ProtoMessage defines model for ProtoMessage.
type ProtoMessage struct {
	// ConversationId Unique identifier for the conversation.
	ConversationId openapi_types.UUID `json:"conversationId"`

	// Message Message content.
	Message string `json:"message"`

	// OriginId Unique identifier for the ID of the sender.
	OriginId openapi_types.UUID `json:"originId"`
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
	PasswordHash  *string             `json:"passwordHash,omitempty"`
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
	PasswordHash  *string             `json:"passwordHash,omitempty"`
	SignedUpAt    *time.Time          `json:"signedUpAt,omitempty"`
	UserId        openapi_types.UUID  `json:"userId"`
}

// UserLogin Payload for user logins
type UserLogin struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// UserSignup Payload for user signups
type UserSignup struct {
	Email     openapi_types.Email `json:"email"`
	FirstName *string             `json:"first_name,omitempty"`
	LastName  *string             `json:"last_name,omitempty"`
	Password  string              `json:"password"`
}

// CreateConversationJSONBody defines parameters for CreateConversation.
type CreateConversationJSONBody = []openapi_types.UUID

// GetRatingByIdParams defines parameters for GetRatingById.
type GetRatingByIdParams struct {
	UserType *GetRatingByIdParamsUserType `form:"userType,omitempty" json:"userType,omitempty"`
}

// GetRatingByIdParamsUserType defines parameters for GetRatingById.
type GetRatingByIdParamsUserType string

// UserLoginJSONRequestBody defines body for UserLogin for application/json ContentType.
type UserLoginJSONRequestBody = UserLogin

// UserRegisterFormdataRequestBody defines body for UserRegister for application/x-www-form-urlencoded ContentType.
type UserRegisterFormdataRequestBody = UserSignup

// CreateConversationJSONRequestBody defines body for CreateConversation for application/json ContentType.
type CreateConversationJSONRequestBody = CreateConversationJSONBody

// CreateMeetingJSONRequestBody defines body for CreateMeeting for application/json ContentType.
type CreateMeetingJSONRequestBody = CreateMeetingBody

// CreateMessageJSONRequestBody defines body for CreateMessage for application/json ContentType.
type CreateMessageJSONRequestBody = ProtoMessage

// UpdateMessageByIdJSONRequestBody defines body for UpdateMessageById for application/json ContentType.
type UpdateMessageByIdJSONRequestBody = Message

// CreateMessageAttachmentJSONRequestBody defines body for CreateMessageAttachment for application/json ContentType.
type CreateMessageAttachmentJSONRequestBody = MessageAttachment

// PostRatingJSONRequestBody defines body for PostRating for application/json ContentType.
type PostRatingJSONRequestBody = Rating

// SignUpAsTutorJSONRequestBody defines body for SignUpAsTutor for application/json ContentType.
type SignUpAsTutorJSONRequestBody = Tutor

// UpdateUserByIdJSONRequestBody defines body for UpdateUserById for application/json ContentType.
type UpdateUserByIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Log in as an existing user
	// (POST /auth/login)
	UserLogin(w http.ResponseWriter, r *http.Request)
	// Sign up as a new user
	// (POST /auth/register)
	UserRegister(w http.ResponseWriter, r *http.Request)
	// Create a new conversation
	// (POST /conversation)
	CreateConversation(w http.ResponseWriter, r *http.Request)
	// Get all messages in the conversation via conversationId
	// (GET /conversation/messages/{conversationId})
	GetMessagesByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID)
	// Get all conversations the user is a member in
	// (GET /conversation/user/{userId})
	GetConversationsByUserId(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID)
	// Get all userIds in the conversation via conversationId
	// (GET /conversation/{conversationId})
	GetUsersByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID)
	// Create a new meeting
	// (POST /meeting)
	CreateMeeting(w http.ResponseWriter, r *http.Request)
	// Get meetings for user
	// (GET /meetings)
	GetMeetings(w http.ResponseWriter, r *http.Request)
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

// UserLogin operation middleware
func (siw *ServerInterfaceWrapper) UserLogin(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UserLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UserRegister operation middleware
func (siw *ServerInterfaceWrapper) UserRegister(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UserRegister(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateConversation operation middleware
func (siw *ServerInterfaceWrapper) CreateConversation(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateConversation(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMessagesByConversationId operation middleware
func (siw *ServerInterfaceWrapper) GetMessagesByConversationId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "conversationId" -------------
	var conversationId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "conversationId", r.PathValue("conversationId"), &conversationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "conversationId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMessagesByConversationId(w, r, conversationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetConversationsByUserId operation middleware
func (siw *ServerInterfaceWrapper) GetConversationsByUserId(w http.ResponseWriter, r *http.Request) {

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
		siw.Handler.GetConversationsByUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUsersByConversationId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersByConversationId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "conversationId" -------------
	var conversationId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "conversationId", r.PathValue("conversationId"), &conversationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "conversationId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersByConversationId(w, r, conversationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

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

// GetMeetings operation middleware
func (siw *ServerInterfaceWrapper) GetMeetings(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, GitHubOAuthScopes, []string{"read:user"})

	ctx = context.WithValue(ctx, GoogleOAuthScopes, []string{"openid"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMeetings(w, r)
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

	m.HandleFunc("POST "+options.BaseURL+"/auth/login", wrapper.UserLogin)
	m.HandleFunc("POST "+options.BaseURL+"/auth/register", wrapper.UserRegister)
	m.HandleFunc("POST "+options.BaseURL+"/conversation", wrapper.CreateConversation)
	m.HandleFunc("GET "+options.BaseURL+"/conversation/messages/{conversationId}", wrapper.GetMessagesByConversationId)
	m.HandleFunc("GET "+options.BaseURL+"/conversation/user/{userId}", wrapper.GetConversationsByUserId)
	m.HandleFunc("GET "+options.BaseURL+"/conversation/{conversationId}", wrapper.GetUsersByConversationId)
	m.HandleFunc("POST "+options.BaseURL+"/meeting", wrapper.CreateMeeting)
	m.HandleFunc("GET "+options.BaseURL+"/meetings", wrapper.GetMeetings)
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
	m.HandleFunc("DELETE "+options.BaseURL+"/user/{userId}", wrapper.DeleteUserById)
	m.HandleFunc("GET "+options.BaseURL+"/user/{userId}", wrapper.GetUserById)
	m.HandleFunc("PUT "+options.BaseURL+"/user/{userId}", wrapper.UpdateUserById)

	return m
}
