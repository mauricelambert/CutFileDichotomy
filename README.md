![CutFileDichotomy logo](https://mauricelambert.github.io/info/go/security/CutFileDichotomy_small.gif "CutFileDichotomy logo")

# CutFileDichotomy

## Description

This tool cuts a target file recursively.

> I used this script to identify malware in archive file without any compression.
>> I didn't know the algorithm to extract files from a very large archive containing a virus. So I cut the archive file several times with this script to scan the smallest file containing the malware. The antivirus detected all the files containing the virus, so I analysed the smallest file and reversed the malware.

## Requirements

### Download binary

 - *No requirements*

### Compilation

 - Go
 - Go Standard library

## Installation

### Download

Download the executable from [Github](https://github.com/mauricelambert/CutFileDichotomy/releases/latest/).

### Compilation

```bash
git clone https://github.com/mauricelambert/CutFileDichotomy.git
cd CutFileDichotomy
go build CutFileDichotomy.go
```

## Usages

```bash
CutFileDichotomy [filename] (size)
CutFileDichotomy malware.bin
CutFileDichotomy malware.bin 2048
```

## Links

 - [Executable - Github](https://github.com/mauricelambert/MaliciousFileDetector/releases/latest/)

## Licence

Licensed under the [GPL, version 3](https://www.gnu.org/licenses/).
