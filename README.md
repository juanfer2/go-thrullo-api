# go-app-graphql

This project if a copy of thrullo in this

- Run development project
```
fresh -c other_runner.conf
```
### Database ORM
This project use [buffalo-pop](github.com/gobuffalo/buffalo-pop/v2), [Soda CLI](https://gobuffalo.io/en/docs/db/toolbox) and [Fizz](https://github.com/gobuffalo/fizz)

- Config database 
```
soda g config
```
```yml
development:
  dialect: postgres
  database: myNameDatabase
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgresUser:postgresPassword@127.0.0.1:5432/myNameDatabaseTest?sslmode=disable"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgresUser:postgresPassword@127.0.0.1:5432/myNameDatabaseProducction?sslmode=disable"}}
```
- Create Database
Development

```
soda create -a
```
* Test
```
soda create -e test
```
* Producction
```
soda create -e producction
```

- Generate Migrations 
```
soda generate fizz name_of_migration
```

- Run Migrations 
```
soda migrate
```
or
```
soda migrate up
```

- Down Migrations 
```
soda migrate down
```
### Graphql
This project use [gqlgen](https://github.com/99designs/gqlgen) 

-Config gqlgen.yml
```yml
# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - ./src/graph/**/*.graphqls

# Where should the generated server code go?
exec:
  filename: ./src/graph/generated/generated.go
  package: generated

# Uncomment to enable federation
# federation:
#   filename: graph/generated/federation.go
#   package: generated

# Where should any generated models go?
model:
  filename: ./src/models/models_gen.go
  package: models

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: ./src/graph/resolvers
  package: graph

# Optional: turn on use `gqlgen:"fieldName"` tags in your models
# struct_tag: json

# Optional: turn on to use []Thing instead of []*Thing
omit_slice_element_pointers: false

# Optional: set to speed up generation time by not performing a final validation pass.
# skip_validation: true

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
 - "github.com/user_git/my_project_app.git/src/models"

models:
  Board:
    model: models.Board
  BoardInput:
    model: models.Board
  User:
    model: models.User
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
```

- Generate grapqhl
every time thats change in the folder graphql run
```
gqlgen generate
```

### Golang console
This project use [gore](https://github.com/motemen/gore)

- Run console
```
gore -autoimport
```
