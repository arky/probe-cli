# OONI Probe CLI

[![linux-debian-packages](https://github.com/ooni/probe-cli/workflows/linux-debian-packages/badge.svg)](https://github.com/ooni/probe-cli/actions?query=workflow%3Alinux-debian-packages)

The next generation OONI Probe Command Line Interface.

## User setup

1. Go [into the releases](https://github.com/ooni/probe-cli/releases) and download the
release for your architecture and platform

2. Extract the tarball with `tar xvzf ooniprobe_*.tar.gz`

3. Copy the `ooniprobe` binary into a location in your `$PATH`, for example
`/usr/local/bin/ooniprobe`

4. Run `ooniprobe run` to perform all the tests

Optional:

Add a crontab entry (on linux) to run `ooniprobe` daily at a random time:

```bash
(crontab -l 2>/dev/null; echo "$(( ( RANDOM % 60 )  + 1 )) $(( ( RANDOM % 24 )  + 1 )) * * * ooniprobe run") | crontab -
```

On macOS you can configure OONI Probe to run automatically using launchd.

Below is a sample launchd script, that should be placed inside of
`~/Library/LaunchAgents/org.ooni.probe.cli.plist`.

Be sure to replace `/PATH/TO/BINARY/ooniprobe` with the actual install location of the
`ooniprobe` binary and `/PATH/TO/CONFIG/config-100sites.json` with the location of a file
which limits the testing to 100 URLs.

You may also want to adjust the locations of the logs.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>org.ooni.probe.daily-run</string>

    <key>KeepAlive</key>
    <false/>
    <key>RunAtLoad</key>
    <true/>

    <key>Program</key>
    <string>/PATH/TO/BINARY/ooniprobe</string>
    <key>ProgramArguments</key>
    <array>
        <string>--config="/PATH/TO/CONFIG/config-100sites.json"</string>
        <string>--batch</string>
        <string>run</string>
    </array>

    <key>StartInterval</key>
    <integer>3600</integer>

    <key>StandardErrorPath</key>
    <string>/tmp/ooniprobe-cli.err</string>

    <key>StandardOutPath</key>
    <string>/tmp/ooniprobe-cli.out</string>
</dict>
</plist>
```

Once you have written the file, you can enable `ooniprobe` to run automatically by
doing: `launchctl load org.ooni.probe.cli.plist`.

## Reporting issues

Please, report issues with this codebase at https://github.com/ooni/probe.

## Development setup

Be sure you have golang >= 1.14 and a C compiler (when developing for Windows, you
need Mingw-w64 installed). The most basic build command is:

```bash
go build -v ./cmd/ooniprobe
```

To compile a release used the `build.sh` script. For more information

```bash
./build.sh help
```

The output generated by this command should provide you with updated information
regarding the pre-requisites for building (and cross-building) `ooniprobe` as well
as useful information regarding cross compiling.

To update bundled binary data use:

```bash
./updatebindata.sh
```

## Updating dependencies

`go get -u -v ./... && go mod tidy`

## Releasing

1. update binary data as described above;

2. update `internal/version/version.go`;

3. make sure you have updated dependencies;

4. run `./build.sh release` and follow instructions.
