### Todo

- redesign store, no longer need mutex
- core/server
  - get and set
- sync core/store with db on load
- api versioning
- OpenApi translator
  - Will need YAML parser
- custom output type for compiler
- typescript schema generator
- cli
- web client
  - Objects
    - send
    - get
    - update
  - forms for building and editing
  - group by category and components
- url registry similar to Category
- Object composition?
- error handling for Reader (unkown dirs, file formats, etc)
- tests
  - package
  - api
- compression file format
- each component should have logging level outputs
- generators
  - protobuf
  - openapi
  - langs
    - typescript
    - golang
- lsp
- validate component names
- schema validators (easy extensions and built-in)

### Decisions

- logging
  - log output levels
  - logger choice
- allow generators to be importable from generators pkg manager
  - create public repo ala NPM
- concurrent dir scanning (will need benchmark)
- do basic validation (names, regex, etc) at parsing

### Hierarchy (bottom-up)

`Props`->`Objects`->`Routes`->`Actions`
