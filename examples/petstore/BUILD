subinclude("//build_defs:openapi")

openapi_generate(
    name = "petstore#gen",
    srcs = [
        "go-server.json",
        "petstore.yaml",
        "//examples/petstore/generator-templates/go-server:template",
    ],
    spec = "petstore.yaml",
    generator = "go-server",
    config = "go-server.json",
    extra_flags = ["-t examples/petstore/generator-templates/go-server"],
    out_dirs = ["go"],
)

go_library(
    name = "petstore",
    srcs = [":petstore#gen"],
    deps = [
            "//third_party/go:mux",
    ],
)

go_binary(
    name = "example",
    srcs = ["main.go"],
    deps = [":petstore"],
)
