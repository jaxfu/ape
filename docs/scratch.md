### Todo

- Linker
  - need to redesign Store to link properly
- Routes
  - Store
- OpenApi translator
  - Will need YAML parser
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

- need better abstraction for Props vs Opts, too confusing currently
- separation between Components intended to be shown in APE format
  vs Props w/ backend data
- concurrent dir scanning (will need benchmark)
- allow generators to be importable from generators pkg manager
  - create public repo ala NPM
- logging
  - log output levels
  - logger choice
- probably should do basic validatin (names, regex, etc) FileHandler
  in compilation

### Hierarchy (bottom-up)

`Props`->`Objects`->`Routes`->`Actions`

### Ingest Flow

**_From Project Dir:_**

Project Dir -> File -> FileHandler -> Bytes ->
Scanner -> Assembler -> AllComponents -> Store

**_From Api Request:_**

Request -> Bytes -> Scanner -> Assembler ->
AllComponents -> Store
