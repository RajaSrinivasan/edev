module gitlab.com/RajaSrinivasan/edev/server

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/google/uuid v1.1.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.2
	gitlab.com/RajaSrinivasan/edev/tools v0.0.0-00010101000000-000000000000
)

replace gitlab.com/RajaSrinivasan/edev/tools => ../tools
