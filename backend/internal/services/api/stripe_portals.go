package api

import (
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"
	"open-tutor/stripe_client"

	"github.com/stripe/stripe-go/v81"
)

func (t *OpenTutor) ViewBillingPortal(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var customerId string
	err := db.GetDB().QueryRow("SELECT stripe_customer_id FROM users WHERE user_id = $1", authInfo.UserID).Scan(&customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}

	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerId),
		ReturnURL: stripe.String("http://localhost:5173/"),
	}

	session, err := stripe_client.GetClient().BillingPortalSessions.New(params)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error generating billing portal redirect: %v", err))
		return
	}

	http.Redirect(w, r, session.URL, http.StatusSeeOther)
}

func (t *OpenTutor) TutorIdVerification(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var accountID string
	err := db.GetDB().QueryRow("SELECT stripe_account_id FROM tutors WHERE user_id = $1", authInfo.UserID).Scan(&accountID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}

	// params := &stripe.IdentityVerificationSessionParams{
	// 	Type:      stripe.String("document"),
	// 	ReturnURL: stripe.String("http://localhost:5173/"),
	// 	Metadata: map[string]string{
	// 		"account_id": accountID,
	// 	},
	// }

	// params.Options = &stripe.IdentityVerificationSessionOptionsParams{
	// 	Document: &stripe.IdentityVerificationSessionOptionsDocumentParams{
	// 		AllowedTypes: []*string{
	// 			stripe.String("driving_license"),
	// 			stripe.String("passport"),
	// 			stripe.String("id_card"),
	// 		},
	// 		RequireMatchingSelfie: stripe.Bool(true),
	// 	},
	// }

	// params.SetStripeAccount(accountID)

	// session, err := stripe_client.GetClient().IdentityVerificationSessions.New(params)
	// if err != nil {
	// 	sendError(w, http.StatusInternalServerError, fmt.Sprintf("error creating verification session: %v", err))
	// }

	params := &stripe.AccountLinkParams{
		Account:    stripe.String(accountID),
		RefreshURL: stripe.String("http://localhost:5173/"),
		ReturnURL:  stripe.String("http://localhost:5173/"),
		Type:       stripe.String("account_onboarding"),
	}

	accountLink, err := stripe_client.GetClient().AccountLinks.New(params)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error generating account onboarding portal redirect: %v", err))
	}

	http.Redirect(w, r, accountLink.URL, http.StatusSeeOther)

	// params := &stripe.BillingPortalSessionParams{
	// 	Customer:  stripe.String(customerId),
	// 	ReturnURL: stripe.String("http://localhost:5173/"),
	// }

	// session, err := stripe_client.GetClient().BillingPortalSessions.New(params)
	// if err != nil {
	// 	sendError(w, http.StatusInternalServerError, fmt.Sprintf("error generating billing portal redirect: %v", err))
	// 	return
	// }

	// http.Redirect(w, r, session.URL, http.StatusSeeOther)
}
