subinclude("//build_defs:openapi")

openapi_bundle(
    name = "yaml",
    srcs = [
        ":spec",
    ],
    entrypoint = "petstore.yaml",
)

filegroup(
    name = "spec",
    srcs = glob(["**/*.yaml"]),
)
