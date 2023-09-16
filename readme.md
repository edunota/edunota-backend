# edunota backend


# main project layout

```bash
├── api
│   └── # OpenAPI/Swagger specs, JSON schema files, protocol definition files.
├── cmd
│   ├── main.go # entry file(s) for project
├── configs
│   └── # config files
├── database https://github.com/edunota/edunota-database # submodule
├── githooks
│   └── # git hooks here
├── go.mod 
├── go.sum
├── guideline https://github.com/edunota/edunota-guideline # submodule
├── infrastructure https://github.com/edunota/edunota-infrastructure # submodule
├── internal
│   ├── app # all applications within this project
│   │   ├── readme.md
│   │   └── rest # rest api
│   │       ├── handlers # handlers root 
│   │       │   ├── note # note handlers
│   │       │   └── user # user handler 
│   │       │       ├── handler.user.go # handler file 
│   │       │       └── handler.user_test.go # handler unit test file
│   │       └── rest.go # entry point for rest api
│   ├── modules # root domain
│   │   ├── module.note # domain/note
│   │   ├── module.user # domain/user
│   │   └── readme.md
│   ├── readme.md
│   └── storage # all storage implemantations here
│       ├── postgres.go # postgres connection 
│       └── readme.md
├── logs # log files
├── readme.md
└── scripts # script files
    └── tree.sh
```