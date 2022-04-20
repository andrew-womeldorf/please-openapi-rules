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
    out = "openapi.json",
    entrypoint = "petstore.yaml",
    extra_flags = [
        "--format",
        "json",
    ],
)

filegroup(
    name = "spec",
    srcs = glob(["**/*.yaml"]),
)
