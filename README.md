# etld1

**A tiny command-line tool that extracts the registrable domain (eTLD+1) from URLs or hostnames using the [Mozilla Public Suffix List](https://publicsuffix.org/).**

Example:

```
sub.service.example.co.uk → example.co.uk
https://deep.nested.example.com/path → example.com
```

---

## Install

### Prebuilt binaries

Download the latest release for your platform from the [Releases](../../releases) page.

### From source

```bash
git clone https://github.com/Firefishy/etld1.git
cd etld1
go build -o etld1 .
```

Or with GoReleaser installed:

```bash
goreleaser build --snapshot --single-target --clean
```

---

## Usage

Pipe domains or URLs into the program. Each line outputs the base registrable domain.

```bash
echo 'sub.service.example.co.uk' | ./etld1
# example.co.uk

echo 'https://deep.nested.example.com/path' | ./etld1
# example.com

cat urls.txt | ./etld1
```

Lines that do not parse as a URL or domain are silently skipped.

---

## Why?

When normalising URLs, it’s often important to reduce them to the domain users actually control.
For example, `co.uk` is controlled by the registry, but `example.co.uk` is registrable.
This tool uses the official PSL (via Go’s `x/net/publicsuffix`) to get the correct base domain.

---

## License

MIT
