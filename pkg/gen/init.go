package gen

import (
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"
)

func Init(cfgPath string, dryRun bool) error {
	g := gen.New(cfgPath)
	g.SetDryRun(dryRun)

	cfg := gen.NewConfig()
	if err := g.Touch(cfgPath); err != nil {
		return err
	}

	act := gen.NewAction("my_controller", "templates/controller.tmpl:/tmp/{{.Pkg}}/controller.go")
	act.Variables = map[string]interface{}{
		"Controller":        `{{ pascalcase .Pkg }}Controller`,
		"Entity":            `{{ pascalcase .Pkg }}`,
		"Service":           `{{ pascalcase "$PKG" }}Service`,
		"CamelCaseSingular": `{{ camelcase .Pkg }}`,
		"CamelCasePlural":   `{{ camelcase .Pkg }}s`,
		"SnakeCaseSingular": `{{ snakecase .Pkg }}`,
		"SnakeCasePlural":   `{{ snakecase .Pkg }}s`,
	}

	cfg.Version = "0.1"
	cfg.Templates = append(cfg.Templates, &gen.Template{
		Name:    "my_action",
		Actions: []*gen.Action{act},
		Environment: map[string]string{
			"Greeting": "$GREETING",
		},
	})

	if err := g.WriteConfig(cfg); err != nil {
		return err
	}
	fmt.Printf("%s: config written\n", cfgPath)
	return nil
}
