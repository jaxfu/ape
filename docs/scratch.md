### Todo

- v2 rewrite
  - parser
    - make more readable
    - bugfix "?Comment :: .Message"
    - handle enum keys that may be references
  - assembler or assembler
- handle refs
- client
  - render components
  - form for creating
- builtin api versioning
- OpenApi translator
- Typescript generator
- custom output type for compiler
- cli
- web client
  - objects
    - send
    - get
    - update
  - forms for building and editing
  - group by category and components
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
- object composition?

### Hierarchy (bottom-up)

`Props`->`Objects`->`Routes`->`Actions`

### Compilation

1. Create each component individually (with no regards to relationships)
   - Much easier to understand compilation methods this way
   - Separately store relationship data

2. Recursively walk and link/add ids
