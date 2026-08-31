# woad.Go - TODO <!-- omit in toc -->


## Table of Contents <!-- omit in toc -->

- [Functional improvements](#functional-improvements)
- [Performance improvements](#performance-improvements)
- [Packaging improvements](#packaging-improvements)


## Functional improvements

* [x] ~~~SGR colour and reset codes~~~ - ✅;
* [ ] TTY-conditional colour codes (process and per-stream);
* [ ] Windows virtual-terminal gating (OS build + `GetConsoleMode`);


## Performance improvements

* \<none>


## Packaging improvements

* [ ] Before the next official release: confirm **`go.mod`** (`go 1.21`) and the CI Go-version matrix, bump Synesis `require`s to newly published tags, then run **`go mod tidy`** (not against currently published tags). Prior Synesis Go releases, in order:
  * **ver2go**;


<!-- ########################### end of file ########################### -->
