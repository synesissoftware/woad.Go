# woad.Go - Example - **colour**


## Summary

Prints sample words in named foreground colours only when standard output supports colour. The words are always printed. The SGR sequences come from `SeqFor`.


## Source

See [examples/colour/main.go](./colour/main.go).


## Execution

```bash
go run ./examples/colour
```

On a colour-capable terminal the words appear in the named colours. When the policy inhibits colour, or standard output is not a terminal, or `TERM` is `dumb` and the policy is automatic, the same lines are printed with no SGR sequences.


<!-- ########################### end of file ########################### -->
