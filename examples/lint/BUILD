subinclude("//build_defs:openapi")

openapi_lint(
    name = "string_string",
    srcs = ["petstore.yaml"],
    entrypoint = "petstore.yaml",
    linters = "redocly",
)

openapi_lint(
    name = "single_list",
    srcs = ["petstore.yaml"],
    entrypoint = "petstore.yaml",
    linters = ["redocly"],
    out = "single_list.sh"
)

openapi_lint(
    name = "single_dict",
    srcs = [
        "petstore.yaml"
    ],
    entrypoint = "petstore.yaml",
    linters = {
        "redocly": [
            "--format", "json",
        ],
    },
    out = "single_dict.sh"
)
