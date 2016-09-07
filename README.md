# Schedule

Get the current TJHSST bell schedule from the command line.

To install, run:

```bash
$ make install
```

Installation will likely require <code>sudo</code>.

After installation, run the command:

```bash
$ sched
```

To uninstall, run (also likely requires sudo):

```bash
$ make uninstall
```

This program has dependencies on many GNU/Linux utilities such as <code>curl</code>, <code>grep</code>, and <code>sed</code>.  These are most likely installed, but if not are required for this program to function.

Uses TJHSST Intranet3 (Ion) Web API.
