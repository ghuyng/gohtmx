# Minimal working web app with Go + HTMX + Tailwind CSS

This is a minimal working web app with Go and HTMX. Use this as a starting point for your web app.

## Tech Stack

1. Go web framework: [Echo](https://echo.labstack.com/)
2. Logger: [Zap](https://pkg.go.dev/go.uber.org/zap)
3. Configuration: [Viper](https://pkg.go.dev/github.com/spf13/viper)
4. Template engine: [a-h/templ](https://pkg.go.dev/github.com/a-h/templ)
5. Frontend library: [HTMX](https://htmx.org/)
6. Frontend CSS framework: [Tailwind CSS](https://tailwindcss.com/)
7. Live reload: [Air](https://github.com/cosmtrek/air)

## Commands

- `make run`: Run the web app
- `make dev`: Run the web app with live reload
- `make build`: Build the web app
- `make test`: Run the tests
- `make lint`: Run the linter
- `make fmt`: Format the code
- `make templ-generate`: Generate code for templ
- `make css-build`: Build the CSS
- `make css-watch`: Watch the CSS

## How to Develop

1. Run `make css-watch` for tailwindcss to watch the CSS on changes
2. Update environment variables in `env/web.env`
3. Open a new terminal and run `make dev` to start the web app with live reload
