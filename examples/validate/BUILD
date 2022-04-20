subinclude("//build_defs:openapi")

filegroup(
    name = "spec",
    srcs = ["petstore.yaml"],
)

filegroup(
    name = "petstore#config",
    srcs = ["go-server.json"],
)

openapi_lint(
    name = "lint",
    srcs = [
        "petstore.yaml",
    ],
    entrypoint = "petstore.yaml",
    linters = "redocly",
)

openapi_generate(
    name = "petstore#gen",
    config = ":petstore#config",
    generator = "go-server",
    out_dirs = ["go"],
    spec = ":spec",
)

go_library(
    name = "validate",
    srcs = [
        "api_spec.go",
        "service.go",
        ":petstore#gen",
    ],
    resources = [":spec"],
    deps = ["//third_party/go:mux"],
)

go_binary(
    name = "example",
    srcs = ["main.go"],
    deps = [":validate"],
)

openapi_test(
    name = "test",
    sandbox = True,
    spec_url = "http://localhost:8080/openapi.yaml",
    webserver = ":example",
)
