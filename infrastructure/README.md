# Terraform Setup

You'll have to set mock access credentials for Localstack in the devcontainer.

1. Go to the Localstack dashboard and go to Auth Tokens.
2. Copy your auth token and run `export LOCALSTACK_AUTH_TOKEN="KEYHERE"` in the devcontainer terminal.

To apply Terraform configuration, run `cd infrastructure && npm run apply`.
