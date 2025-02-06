# Tutoring Site - UNT Capstone 2024

## Devcontainer Instructions

You should typically only have to do this on a devcontainer rebuild.
If any changes are made to the .devcontainer folder, a rebuild is
probably necessary, otherwise ask backend team.

Once the devcontainer is set up, you should be able to skip the
setup steps below and just open the devcontainer when prompted when you load
VSCode.

1. Open VSCode
2. Open project in container dialog box should show in bottom right corner.
Or View > Command Palette... > Dev Containers: Reopen in Container.
3. Wait for the devcontainer to start. At the end of building/rebuilding,
the devcontainer will prompt to you "Hit Enter" when the build script is
done. **DO THIS**!!

If you don't acknowledge and close the setup script,
the necessary port forwarding may not activate correctly.
4. Open and set up pgAdmin: go to [localhost:5050](http://localhost:5050)
Use the email and password below to log in to PGAdmin.

```
Username: dev@opentutor.io
Password: password  
```

Now right click Servers, then find "Register > Server..."

Use these credentials to register the development postgres database in pgAdmin.

```
Host Name: postgres
Database: opentutor
Username: dev
Password: password
```