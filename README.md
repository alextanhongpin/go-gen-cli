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
