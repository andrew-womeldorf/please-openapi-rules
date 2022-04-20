subinclude("//build_defs:openapi")

filegroup(
    name = "petstore#config",
    srcs = ["go-server.json"],
)

filegroup(
    name = "petstore#spec",
    srcs = ["petstore.yaml"],
)

openapi_generate(
    name = "petstore#gen",
    config = ":petstore#config",
    generator = "go-server",
    out_dirs = ["go"],
    spec = ":petstore#spec",
)

go_library(
    name = "generator",
    srcs = [":petstore#gen"],
    deps = ["//third_party/go:mux"],
)

go_binary(
    name = "example",
    srcs = ["main.go"],
    deps = [":generator"],
)
