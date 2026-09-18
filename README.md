## Title: 

** Intelligent collaborative workspace for teams

## Overview

NoxOrbit is a Go-based backend for a collaborative workspace. The product is designed to bring projects, tasks, decisions, team context, and intelligent insights into one place.

The project is currently in active development. The documentation describes the planned product and architecture alongside the functionality that is being implemented.


## Features:

workspaces
projects
tasks
comments
notifications
AI insights

## Tech stack:

Go
PostgreSQL
HTTP router
Gmail SMTP

## Run locally:

go mod download
set env vars
go run ./backend/cmd/server (for now)

## Project structure:

backend/cmd/server
backend/config
backend/internal/...

## Documentation:

docs/prd.md
docs/architecture.md
docs/database.md
docs/api.md
docs/roadmap.md
