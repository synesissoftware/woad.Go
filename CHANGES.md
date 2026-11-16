# woad.Go - Changes <!-- omit in toc -->


## 0.2.0-alpha1 - 17th November 2026

* added `BLINKING` and `BOLD`, byte-for-byte with **woad** **0.3.0**;
* added cached `ColourPolicy`, `Refresh`, `SupportsColour`, and `SeqFor` for the portable colour contract (Unix attachment via **golang.org/x/term**, and `TERM=dumb`; no console probe);
* added decision rows 1-22; **examples/colour** and **examples/libver** emit sequences only when the writer supports colour;
* left Windows virtual-terminal mode and mintty PTY recognition for a later alpha;


## 0.1.1 - 20th August 2026

* enforced Synesis Go import order via **gci** (**.golangci.yml**, **examples/.golangci.yml**);
* moved example companion markdown to `examples/<name>.md` (alongside `examples/<name>/main.go`);
* version string updated for the 0.1.1 release;


## 0.1.0 - 20th August 2026

* added **Version()** (replacing the **Version** constant), formed by **ver2go.CombineVersion()**;
* updated **ver2go** to 0.2.0-beta1;
* version string updated for the 0.1.0 release;


## 0.0.1 - 20th August 2026

* Initial release;
* Added fixed ANSI SGR colour constants (**RESET**, **FG_***, **BG_***, including bright variants);
* Added version API via **ver2go**;
* Added Synesis Go boilerplate (CI, scripts, docs, examples);


<!-- ########################### end of file ########################### -->
