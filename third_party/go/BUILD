package(default_visibility = ["PUBLIC"])

go_module(
    name = "oauth2",
    module = "golang.org/x/oauth2",
    version = "6242fa91",
    deps = [
        ":net",
    ],
    install = [
        ".",
        "internal",
    ],
)

go_module(
    name = "net",
    module = "golang.org/x/net",
    version = "cd36cc0744dd",
    install = ["context/ctxhttp"]
)

go_module(
    name = "mux",
    module = "github.com/gorilla/mux",
    version = "v1.8.0",
)
