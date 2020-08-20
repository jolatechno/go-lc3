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

For example here to run a simple hello world program on an lc3 vm you can use

```
./main lc3 hello.obj
```

## Cpus

### Lc3

The lc3 vm was inspired by [Write your Own Virtual Machine](https://justinmeiners.github.io/lc3-vm/index.html) and [ziggy42/gLC3](https://github.com/ziggy42/gLC3/).

Images are compiled using [lc3web](https://wchargin.github.io/lc3web/).

#### Usage

To simply run an lc3 vm image, use :

```go
package main

import (
  "github.com/jolatechno/go-lc3/src/lc3"
)

func main() {
  instructions.DebugLc3Instructions = true /* remove if you don't want debuging */

  lc3cpu := lc3.New() /* defining a new lc3 cpu */

  _, err := cpu.LoadFile(lc3cpu, "your_image_filename") /* test the loadFile function */
  if err != nil { /* throwing an error */
      panic(err)
  }

  err = cpu.Run(lc3cpu) /* test the run function */
  if err != nil { /* throwing an error */
      panic(err)
  }
}
```

To read a value off of the memory of `cpu` at `address` use :

```go
value, err := cpu.Memory().Read(address) /* read value off of memory */
```

To write `value` to it use :

```go
err := cpu.Memory().Write(address, value) /* write value to memory */
```

And to reset it use :

```go
cpu.Memory().Reset() /* reset memory */
```

The same thing can be applied to registers :

```go
value, err := cpu.Registers().Read(register) /* read value off of registers */

err := cpu.Registers().Write(register, value) /* write value to registers */

cpu.Registers().Reset() /* reset registers */
```

with the added option of `fetching` the pc register and incrementing it at the same time :

```go
pc, err := cpu.Registers().Fetch() /* reading the pc register and incrementing it */
```
