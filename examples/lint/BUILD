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
    out = "single_list.sh",
    entrypoint = "petstore.yaml",
    linters = ["redocly"],
)

openapi_lint(
    name = "single_dict",
    srcs = [
        "petstore.yaml",
    ],
    out = "single_dict.sh",
    entrypoint = "petstore.yaml",
    linters = {
        "redocly": [
            "--format",
            "json",
        ],
    },
)

openapi_lint(
    name = "spectral",
    srcs = [
        "petstore.yaml",
        "spectral-different-name.json",
    ],
    out = "spectral-ruleset-overwrite.sh",
    entrypoint = "petstore.yaml",
    linters = {
        "spectral": [
            "--ruleset",
            "$(location spectral-different-name.json)",
        ],
    },
)

openapi_lint(
    name = "all_linters",
    srcs = [
        "petstore.yaml",
        "spectral.json",
    ],
    out = "all_linters.sh",
    entrypoint = "petstore.yaml",
)
