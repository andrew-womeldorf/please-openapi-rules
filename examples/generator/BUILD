subinclude("//build_defs:openapi")

openapi_generate(
    name = "petstore#gen",
    srcs = [
        "go-server.json",
        "petstore.yaml",
    ],
    spec = "petstore.yaml",
    generator = "go-server",
    config = "go-server.json",
    out_dirs = ["go"],
)

go_library(
    name = "generator",
    srcs = [":petstore#gen"],
    deps = [
            "//third_party/go:mux",
    ],
)

go_binary(
    name = "example",
    srcs = ["main.go"],
    deps = [":generator"],
)
