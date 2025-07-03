### Todo

- ape/Components/Component redesign
- validate component names
- handle refs
- client
  - render components
  - form for creating
- sync core/store with db on load
- compile project dir
- api versioning
- OpenApi translator
- typescript schema generator
- cli
- web client
  - objects
    - send
    - get
    - update
  - forms for building and editing
  - group by category and components
- url registry similar to Category
- error handling for Reader (unkown dirs, file formats, etc)
- tests
  - package
  - api
- compression file format
- generators
  - protobuf
  - openapi
  - langs
    - typescript
    - golang
- lsp
- schema validators (easy extensions and built-in)

### Decisions

- logging
  - log output levels
  - logger choice
- allow generators to be importable from generators pkg manager
  - create public repo a la NPM
- custom output type for compiler
- Object composition

### Hierarchy (bottom-up)

`Props`->`Objects`->`Routes`->`Actions`

### Compilation

1. Create each component individually (with no regards to relationships)
   - Much easier to understand compilation methods this way
   - Separately store relationship data

2. Recursively walk and link/add ids
