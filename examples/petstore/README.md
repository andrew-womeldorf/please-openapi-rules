Using the `openapi_generator` rule to generate a go server. This rule outputs a
`filegroup`, which can then be used in a `go_library` rule.

Note that in the default `petstore.yaml` that's generally used, there's a POST
endpoint with a file upload. The `go-server` generator doesn't appear to be
importing the `os` module, and I'm not yet famliiar enough to fix it, so I just
removed that endpoint for this example.
