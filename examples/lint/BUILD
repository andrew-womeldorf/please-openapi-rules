subinclude("//build_defs:openapi")

openapi_lint(
    name = "single_string",
    srcs = [
        "petstore.yaml",
        "spectral.json",
    ],
    entrypoint = "petstore.yaml",
    linters = "spectral",
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

openapi_lint(
    name = "spectral",
    srcs = [
        "petstore.yaml",
        "spectral-different-name.json",
    ],
    entrypoint = "petstore.yaml",
    linters = {
        "spectral": [
            "--ruleset", "$(location spectral-different-name.json)",
        ],
    },
    out = "spectral-ruleset-overwrite.sh"
)

openapi_lint(
    name = "all_linters",
    srcs = [
        "petstore.yaml",
        "spectral.json",
    ],
    entrypoint = "petstore.yaml",
    out = "all_linters.sh",
)
