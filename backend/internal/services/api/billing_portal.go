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
