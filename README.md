# OpenAPI Rules for Please.Build

This repo is a plugin that provides OpenAPI rules for the
[Please](https://please.build) build system.

Follow the example set by projects in
[github.com/please-build](https://github.com/please-build), and
[docs](https://please.build/config.html#plugindefinition).

## Basic Usage

Include this plugin in your project:

```python
# BUILD
plugin_repo(
    name = "please-openapi",
    owner = "andrew-womeldorf",
    revision = "<Some git tag, commit, or other reference>",
)
```

Use it in a `BUILD` file:

```python
# some_dir/BUILD
subinclude("///please-openapi//build_defs:openapi")

openapi_generate(
    name = "spec",
    srcs = ["openapi.yaml"],
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
