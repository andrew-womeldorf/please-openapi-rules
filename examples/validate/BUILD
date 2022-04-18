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
    linters = "redocly",
    srcs = [
        "petstore.yaml",
    ],
    entrypoint = "petstore.yaml",
)

openapi_generate(
    name = "petstore#gen",
    spec = ":spec",
    config = ":petstore#config",
    generator = "go-server",
    out_dirs = ["go"],
)

go_library(
    name = "validate",
    srcs = [
            ":petstore#gen",
            "service.go",
    ],
    deps = ["//third_party/go:mux"],
)

go_binary(
    name = "example",
    srcs = ["main.go"],
    deps = [":validate"],
)
