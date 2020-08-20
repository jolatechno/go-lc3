package main

import (
  "os"
  "os/exec"
  "testing"
)

func TestLc3Cpu(t *testing.T) {
  cmd := exec.Command("go", "run", "main.go", "lc3", "hello.obj") /* command to test the main command */
  cmd.Stdout = os.Stdout /* redirect output to stdout */

  err := cmd.Run() /* run the command */
  if err != nil { /* thrown an error */
    t.Error(err)
  }
}
