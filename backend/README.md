# Backend

## oapi-codegen

The OpenAPI specification is found in `backend/api` and should **ALWAYS**
be validated against the swagger editor after changes are made.

For generating go server stubs, oapi-codegen should be run from within the
backend folder to avoid weird behavior.  Codegen should be run any time there
are changes to the OpenAPI specification.

To codegen, do the following while working dir is `~/../project_dir/backend/`:

`scripts/update_oapi.sh`

This is destructive and will overwrite the existing `api.gen.go` in
`~/../proj_dir/backend/internal/services`

### Codegen Help

Reach out to [Nathan Jodoin](nathan@jodoin.io) by email or discord with questions.
