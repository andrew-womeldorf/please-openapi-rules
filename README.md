# OpenAPI Rules for Please.Build

This repo is a plugin that provides OpenAPI rules for the
[Please](https://please.build) build system.

Follow the example set by projects in
[github.com/please-build](https://github.com/please-build), and
[docs](https://please.build/config.html#plugindefinition).

## Basic Usage

Include this plugin in your project:

```
; .plzconfig
[Plugin "openapi"]
Target = //plugins:openapi

; Note: I don't think you actually need to define the js target in
; plugins/BUILD, but you need to define the plugin here with a Target argument.
[Plugin "js"]
Target = //plugins:js
```

```python
# plugins/BUILD
plugin_repo(
    name = "openapi",
    owner = "andrew-womeldorf",
    plugin = "please-openapi",
    revision = "<Some git tag, commit, or other reference>",
)
```

Use it in a `BUILD` file:

```python
# some_dir/BUILD
subinclude("///openapi//build_defs:openapi")

filegroup(
    name = "spec",
    srcs = glob(["**/*.yaml"]),
)

openapi_bundle(
    name = "yaml",
    srcs = [":spec"],
    entrypoint = "petstore.yaml",
    out = "openapi.yaml",
)

openapi_lint(
    name = "lint",
    srcs = [
        ":yaml",
        "spectral.json",
    ],
    entrypoint = "openapi.yaml",
)

openapi_docs(
    name = "docs",
    srcs = [":yaml"],
    entrypoint = "openapi.yaml",
)

openapi_preview(
    name = "preview",
    srcs = [":yaml"],
    entrypoint = "openapi.yaml",
)

openapi_generate(
    name = "generate",
    srcs = [":yaml"],
    spec = "openapi.yaml",
    generator = "go-server",
)
```

## Configuration

This plugin can be configured via the plugins section as follows:

```
[Plugin "please-openapi"]
SomeConfig = some-value
```

### GeneratorTool (str, target)

The Generator rule uses the [OpenAPI
Generator](https://openapi-generator.tech/). See
`//third_party/java:openapi-generator-cli` for the current version being used
and an example of providing your own tool.

```
[Plugin "please-openapi"]
GeneratorTool = //some/tool:you-want
```
