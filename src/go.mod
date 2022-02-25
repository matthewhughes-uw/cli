module github.com/opslevel/cli

go 1.16

require (
	github.com/creasty/defaults v1.5.1
	github.com/go-git/go-git/v5 v5.4.2
	github.com/gosimple/slug v1.12.0
	github.com/manifoldco/promptui v0.9.0
	github.com/opslevel/opslevel-go v0.4.5-0.20220225194148-bc75f2c38d93
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// Uncomment for local development
// replace github.com/opslevel/opslevel-go => ../../opslevel-go
