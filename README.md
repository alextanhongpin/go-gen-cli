# go-gen-cli

Code generation utility for golang. Avoid copy-pasting, create quality reusable boilerplate.

```bash
$ go get github.com/alextanhongpin/go-gen-cli/cmd/gen
```

```bash
# To view help
$ gen

# Generate config file. If the config already exists, it will skip the step.
$ gen init

# Generate pkg for a given template. If the generated file already exists, it
# will skip the step.
$ gen [g|generate] [-t|--template] <template> <pkg>
$ gen g -t domain user_service

# Remove generated pkg for a given template.
$ gen [c|clear] [-t|--template] <template> <pkg>
$ gen c -t domain user_service
```

## Configuration

Minimal `gen.yaml`:
```yaml
version: "0.1"
templates:
- name: my_action
  actions:
  - name: my_controller
    path: templates/controller.tmpl:/pkg/{{kebabcase .Pkg}}/controller.go
```

The variable `.Pkg` comes from the argument that is passed in when calling the CLI:

```bash
# gen g -t <template-name> <pkg>
$ gen g -t my_action hello
```

## Basic Template

This is how a basic template will look like (`controller.tmpl`):

```tmpl
package {{ snakecase .Pkg }}

type {{ pascalcase .Pkg }}Controller {
}
func New{{ pascalcase .Pkg }}Controller() *{{ pascalcase .Pkg }}Controller {
    return &{{ pascalcase .Pkg }}Controller{}
}

func main() {
    {{ camelcase .Pkg }} := New{{ pascalcase .Pkg }}Controller()
    fmt.Println({{ camelcase .Pkg }})
}
```

## Template functions

The following string case functions are available for use:
- kebabcase
- snakecase
- pascalcase
- camelcase

Other functions from `github.com/Masterminds/sprig` can be used too.

## Environment Variables

Environment variables can be passed in through the environment field, and are accessible in the `gen.yaml` as well as the templates.

	NOTE: The `path` does not accept environment variables.

```yaml
version: "0.1"
templates:
- name: my_action
  environment:
    Greeting: $GREETING
  actions:
  - name: my_controller
    path: templates/controller.tmpl:/tmp/{{.Greeting}}/controller.go
```

## Variables

To keep things dry and simple, variables can be used to provide computed values:

```yaml
version: "0.1"
templates:
- name: my_action
  actions:
  - name: my_controller
    path: templates/controller.tmpl:/tmp/{{.Pkg}}/controller.go
    variables:
      Controller: '{{ .Pkg }}Controller'
```

Values from `environment` as well as environment variables can also be used here:

```
version: "0.1"
templates:
- name: my_action
  environment:
    Greeting: $GREETING
  actions:
  - name: my_controller
    path: templates/controller.tmpl:/tmp/{{.Pkg}}/controller.go
    variables:
      Controller: '{{ pascalCase .Greeting }}Controller'
      AnotherController: '{{ pascalCase "$GREETING" }}Controller'
```
The template above will then look like this:
```tmpl
package {{ snakecase .Pkg }}

type {{ .Controller }} {
}
func New{{ .Controller }}() *{{ .Controller }} {
    return &{{ .Controller }}{}
}

func main() {
    {{ camelcase .Controller }} := New{{ .Controller }}()
    fmt.Println({{ camelcase .Controller }})
}
```

## Prompts

You can also include prompts and assign the values to a variable. The example below assigns the user input to the variable `Age`.

```
  prompts:
    - name: Age
      message: Enter your age
      type: input
```

## Auto-format and auto-import

If you are using this to generate go files, it will auto-import missing dependencies and apply go-fmt when generating the file.
