package main

import (
  "os"
  "github.com/jolatechno/go-lc3/src/lc3"
  "github.com/jolatechno/go-lc3/src/cpu"
  "github.com/jolatechno/go-lc3/src/interfaces"
)

func main() {
  args := os.Args[1:]
	if len(args) < 2 { /* check if there are enough arguments */
		panic("Missing arguments")
	}

  var used_cpu interfaces.Cpu /* cpu var */

  switch cpu_name := args[0]; cpu_name { /* see the used cpu */
  case "lc3":
    used_cpu = &lc3.Lc3cpu

  default:
    panic("cpu not understood")
  }

  imgs := args[1:]
  for _, img := range imgs { /* for each iamge */
    _, err := cpu.LoadFile(used_cpu, img) /* Load the image file loadFile */
    if err != nil { /* throwing an error */
        panic(err)
    }
  }

  err := cpu.Run(used_cpu) /* run the vm */
  if err != nil { /* throwing an error */
      panic(err)
  }
}
