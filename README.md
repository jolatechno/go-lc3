# go-lc3

The goal of this project is to make a modular golang VM that can at first run LC3, but can then be made to run other instruction set.

## Build

To build the main command use `go build` inside of this directory.

## Usage

To run one or more images on one of the `cpus`, simply use

```
./main cpu image1.filename image2.filename ...
```

Each image will be loaded in order before before the vm gets initialized.

## Cpus

### Lc3

The lc3 vm was inspired by [Write your Own Virtual Machine](https://justinmeiners.github.io/lc3-vm/index.html) and [ziggy42/gLC3](https://github.com/ziggy42/gLC3/).

Images are compiled using [lc3web](https://wchargin.github.io/lc3web/).

#### Usage

```go
package main

import (
  "github.com/jolatechno/go-lc3/src/lc3"
)

func main() {
  instructions.DebugLc3Instructions = true /* remove if you don't want debuging */

  _, err := cpu.LoadFile(&lc3.Lc3cpu, "your_image_filename") /* test the loadFile function */
  if err != nil { /* throwing an error */
      panic(err)
  }

  err = cpu.Run(&lc3.Lc3cpu) /* test the run function */
  if err != nil { /* throwing an error */
      panic(err)
  }
}
```
