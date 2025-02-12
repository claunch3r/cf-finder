# cf-finder
`cf-finder` is a utility that determines whether an IP address belongs to Cloudflare.

# Usage
```
cf-finder -h
cf-finder -target <ip>
cf-finder -target <ip-1>,<ip-2>
cf-finder -file <targets.txt>
```
This will display help for the tool. Here are all the switches it supports.
```
Usage:
  ./cf-finder [flags]

Flags:
INPUT:
  -target string    input target ip(s) to probe
  -file string      input file containing list of ip addresses to process
```
