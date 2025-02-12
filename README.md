# cf-finder
`cf-finder` is a utility that determines whether an IP address belongs to Cloudflare.
![1](https://github.com/user-attachments/assets/9875cbd8-b7cd-4e1e-a7fb-5bbb245d1187)

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
