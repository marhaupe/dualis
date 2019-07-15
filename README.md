# Dualis

Simple hack to check your grades on https://dualis.dhbw.de. Starts a headless browser and takes a screenshot in a matter of seconds.

# Installation

Download the [archive](https://github.com/marhaupe/dualis/releases) and extract it somewhere your `$PATH` is set to, for example to `/usr/local/bin`:

```bash
tar -C /usr/local/bin -xzf <archive>
```

Alternatively, just download it and extract it somewhere. I'm not your mother.

# Usage

If it's in your `$PATH`:
```bash
dualis -u <max.mustermann@dh-karlsruhe.de> -p <password>
```
or just run `dualis` and input the data that is prompted.


If it's not in the your `$PATH`, but you're in the same folder the binary is saved:
```bash
./dualis -u <max.mustermann@dh-karlsruhe.de> -p <password>
```
or just run `./dualis` and input the data that is prompted.
