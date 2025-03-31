# Open Tutor Deployment Guide

## Overview

This document aims to guide you through deploying the tech stack required to run the Open Tutor application.

## Assumptions

This guide assumes you are deploying this project on a system that can run Docker. If you're running Windows, you will typically need to enable CPU virtualization or Hyper-V in your BIOS settings.

## Dependencies

You must install the latest versio of the following base dependencies to run this project:

[Docker](https://docs.docker.com/engine/install/) (containerization)

## Constraints

You may change the ports that the services run on, however by default you should ensure that the following ports are not occupied by any other servers:

- 5173 (web)
- 8080 (api)
- 5432 (postgres host exposure)
- 5050 (pgadmin dashboard) [optional]

## Description of Artifacts

The artifact we've submitted contains the full source code for OpenTutor so far. Everything you need to run the application will be contained within this archive.

## Data Creation

The following information will need to be configured before Open Tutor can be used.

### Environment files

The following files should be duplicated accordingly:

- `backend/.env.example` duplicated as `backend/.env`
- `frontend/.env.example` duplicated as `frontend/.env`

### Zoom application

To automatically create and manage Zoom meetings, Open Tutor needs an API key for a Zoom application.

1. Navigate to [Zoom Marketplace](https://marketplace.zoom.us/)
2. Sign up for a new account or log into your existing account
3. Go to the "Develop" dropdown in the top right, click "Build App"
4. Select "Server to Server OAuth App" from the list, click "Create"
5. Name it "Open Tutor", click "Create"
6. Navigate to the "Activation" tab on the left side of the window
7. Fill out the requisite information to activate your application
8. Navigate to the "Scopes" tab on the left side of the window
9. Add the following scopes:

- `meeting:write:meeting:admin`
- `meeting:write:meeting:master`

10. Navigate to the "App Credentials" tab on the left side of the window
11. Copy the credentials accordingly into the `backend/.env` file:

- Account ID -> `ZOOM_ACCOUNT_ID`
- Client ID -> `ZOOM_CLIENT_ID`
- Client Secret -> `ZOOM_CLIENT_SECRET`

### Stripe application

For tutors to receive payouts and for students to submit payments, Open Tutor utilizes the Stripe platform. Stripe manages user-to-user payouts and handles legal compliance for such transactions.

1. Navigate to [Stripe](https://stripe.com/)
2. Sign up for a new account
3. Fill out business information as requested
4. Select "Build a platform or marketplace"
5. Click "Go To Dashboard"
6. Click "Get your live account" at the top of the screen
7. Verify email as requested
8. Click "Connect" on the left sidebar
9. Click "Connect overview"
10. Click the ... in the top right, select "Settings"
11. Click "Customize" under Branding
12. Fill out information under "Onboarding experience" as requested
13. Navigate back to [Stripe dashboard](https://dashboard.stripe.com/)
14. Click "Complete your profile" in the bottom right corner
15. Fill out business information as requested (if testing out, do Individual)
16. Enter personal information as requested, continue when done
17. For "Industry" select "Education" -> "Other educational services". Fill out other information on page as requested. Continue when done.
18. Fill out information on Public details tab, continue when done.
19. Link an empty bank account to attach to payout system. This will be the account that receives payments and pays out after a buffer period.
20. Set up authenticator appropriately, continue when done
21. Enter "Educational Services" into "Select your product category", continue
22. Select "Agree and submit"
23. In the bottom right corner, select "Complete Connect setup"
24. Select "Sellers will collect payments directly"
25. Select "On-Demand services"
26. Select "Onboarding hosted by Stripe"
27. Select "Stripe Dashboard" for "Where will sellers manage their account"
28. Review information and click Continue
29. Go to the "Home" tab on the left panel
30. Copy the API keys into the `backend/.env` file accordingly:

- Publishable key -> `STRIPE_PK`
- Secret key -> `STRIPE_SK`

## Admin Credentials

The credentials for the pgadmin dashboard are located under the `pgadmin` service in `docker-compose.yml` by default. You may reconfigure these.

## Deployment Process

Complete the following procedure to start the software stack:

1. Open a new command shell
2. Navigate to the `open-tutor` project directory
3. Run `docker compose up -d`
4. Open Tutor should now be running!
5. Run `docker compose logs -f SERVICE_NAME` to view live logs from a service. Valid service names include:

- `api`
- `web`
- `postgres`
