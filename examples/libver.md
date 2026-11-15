# woad.Go - Example - **libver**


## Summary

Displays the **woad.Go** version and the **ver2go** version, then terminates. The library name, version text, `Version` label, and packed value are coloured the same way as the C **test.scratch.versions** printer. Sequences are emitted only when standard output supports colour.


## Source

See [examples/libver/main.go](./libver/main.go).


## Execution

```bash
go run ./examples/libver
```

On a colour-capable terminal the name is blue, the version text and packed value are green, and the `Version` label is cyan. Otherwise the same lines are printed with no SGR sequences.


<!-- ########################### end of file ########################### -->
