### 10.7.25

- rewrote grammar to v2
- began v2 rewrite w/ ape/compiler/internal/lexer

### 8.7.25

- began major refactor of entire project

### 7.7.25

- started ape/compiler/internal/assembler
  - test added to compiler_test.go

### 3.7.25

- wip ape/compiler/parser steps rewrite
  - depth and indentation check working

### 30.6.25

- wip ape/compiler/parser steps rewrite
  - components working
  - enum members working

### 29.6.25

- wip ape/compiler/parser steps rewrite
  - comments working

### 28.6.25

- wip ape/compiler/parser steps rewrite

### 22.6.25

- ape/compiler/parser
  - classify node working for comp, constraint, comment

### 15.6.25

- ape/compiler/lexer rewrite to handle crlf

### 3.6.25

- ape/compiler/lexer redesigned

### 24.5.25

- ape/compiler/parser working w/ comments & components

### 19.5.25

- rewrote lexer

### 18.5.25

- began redesign into two-phase parser

### 8.5.25

- ape/compiler/parser comments

### 7.5.25

- compile redesign started
  - simpler lexical tokens

### 6.5.25

- ape/compiler/analyzer started

### 5.5.25

- ape/compiler/formatter completed

### 4.5.25

- added enum support to ape/compiler/parser

### 3.5.25

- ape/compiler/parser basic func

### 2.5.25

- ape/pkg/stack generic
- start ape/compiler/parser recursive redesign

### 1.5.25

- start ape/compiler redesign
  - lexer completed

### 30.4.25

- OpenApi unmarshal Schemas -> Props

### 29.4.25

- init OpenApi unmarshaller

### 28.4.25

- GetComponents on client

### 27.4.25

- Get/Create routes working

### 26.4.25

- ape/engine/core EventBus
- CreateComponent

### 25.4.25

- ape/engine/core/validator

### 24.4.25

- init client

### 23.4.25

- made engine/core/store threadsafe
- Linker working for ComponentId

### 22.4.25

- Compiler redesign
- Store.Categories

### 21.4.25

- Core/Store
- Db

### 20.4.25

- Assembler.AssembleResponses
- Completed Assembler.AssembleRoute
- Completed Compiler/Assembler

### 19.4.25

- Assembler.AssembleRequest
- Assembler.AssembleMessageBody

### 18.4.25

- Major comonent redesign
  - Requests, MessageBody, and Response

### 17.4.25

- Assembler.AssembleProp
- init Assembler.AssembleRequest

### 16.4.25

- Parser completed for all types

### 15.4.25

- Parser.ParseProp

### 14.4.25

- major redesign

### 13.4.25

- Assembler.AssembleRoute

### 12.4.25

- Scanner.ScanRoute

### 11.4.25

- Routes schemas

### 10.4.25

- init tests

### 9.4.25

- FileHandler
- major refactor
  - Transformation -> Assembler
  - Scanning -> Scanner

### 8.4.25

- store ComponentId (lowercased) and DisplayId (original)
- Refs linked properly by Repo
- refactor Reader under Core
- track and validate categories
  - Transformer.rawMetadata -> Transformer.handleCategories
- Server init

### 7.4.25

- Reader.readObjectFile bugfixes
- Repo.storeOpts Ref bugfix
- Reader logs print Props under Object

### 5.4.25

- Store basics
- Cache basics
- added AllComponents type
- altered Reader to return AllComponents
- fixed Repo opts handling
- ComponentId as RefTag

### 4.4.25

- Objects fetched properly from Repo

### 3.4.25

- Reader basics
- db handle Opts

### 2.4.25

- db initialization
- Repository

### 31.3.25

- Scanner

### 22.3.25

- v0.1 finished through Routes

### 19.3.25

- completed basic v0.1 spec
