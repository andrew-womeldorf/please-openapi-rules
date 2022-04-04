subinclude("//build_defs:openapi")

openapi_bundle(
    name = "yaml",
    srcs = [
        ":spec",
    ],
    entrypoint = "petstore.yaml",
)

openapi_bundle(
    name = "json",
    srcs = [
        ":spec",
    ],
    extra_flags = ["--format", "json"],
    out = "openapi.json",
    entrypoint = "petstore.yaml",
)

filegroup(
    name = "spec",
    srcs = glob(["**/*.yaml"]),
)
