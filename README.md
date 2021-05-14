# nothing

`nothing` - do nothing and succeed (or fail with a requested exit code)

similar to [`TRUE(1)`](https://man7.org/linux/man-pages/man1/true.1.html) and [`FALSE(1)`](https://man7.org/linux/man-pages/man1/false.1.html) but more

# usage

1. optionally: rename executable to any name you like (i.e. `Adobe Updater.exe`)
2. optionally set env var `<EXECUTABLE-NAME>_EXIT=<exitcode>`. Name is uppercase, without non-alphanumeric characters and without file extension (i.e. `ADOBEUPDATER`)
3. run the executable. (If the env var isn't set it will exit with `0`).

Example (Windows `cmd`):
```
$ move nothing.exe "Adobe Updater.exe"
$ set ADOBEUPDATER_EXIT=99
$ "Adobe Updater.exe"
$ echo %errorlevel%
99
```

# avoiding security confirmations

The Windows and Mac executables from the releases are codesigned and should be trusted by most anti-virus tools.

On MacOS you still may need to remove the quarantine attribute after download to avoid the one-time "it's downloaded, trust it?" popup:

```
xattr -c <file>
```

# why?

I quickly needed a renamable, code-signed dummy `.exe` on windows to replace a misbehaving execuable on many machines.