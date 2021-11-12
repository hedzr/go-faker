package main

//go:generate swag init --output ./swaggerdocs

import (
	_ "github.com/%REPOSITORY%/cli/app/cli/app/swaggerdocs"

	_ "github.com/alecthomas/template"
	_ "github.com/swaggo/swag"
)
