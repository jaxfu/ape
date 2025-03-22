1. do not alter interfaces, only implementations

2. each top-level package in ape/engine/core should have
   no outside dependencies.

   - if neccesary, tie together in ape/engine/core

3. ape/engine/core should have no outside dependencies

   - if neccesary, tie together in main()

4. ape/engine/compiler should have no outside dependencies

5. compilation process should be stateless/functional

   - new compiler will be stack allocated for each
     compilation with ape/engine/compiler.NewCompiler()

6. ape/components should have no outside dependencies

7. utility packages in ape/engine/pkg should
   have no interdependencies

   - they may have dependencies on ape/engine/core as needed
     but never vice versa obviously

8. keep everything on the stack aside from when absolutely necessary,
   e.g. mutexes, stateful recievers (which should be eschewed),
   or representing "optional" values in structs

9. top-level packages should expose explicity from
   package/interface.go

   - create an package/interface.go file in the root
   - implementation in package/internal/
   - no implementations in package/interface.go
   - import from package/internal/ and mirror/expose
