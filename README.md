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

The current period, if applicable, is highlighted in red.

This program has dependencies on many GNU/Linux utilities such as <code>curl</code>, <code>grep</code>, <code>sed</code>, and <code>tput</code>.  These are most likely installed, but if not are required for this program to function.

Originally forked from [jcschefer](https://github.com/jcschefer) and later migrated into standalone repository for further personal development.

Uses TJHSST Intranet3 (Ion) Web API.
