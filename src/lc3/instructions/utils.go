package instructions

import (
  "errors"
  "os"
  "os/exec"
  "github.com/jolatechno/go-lc3/src/interfaces"
  "github.com/jolatechno/go-lc3/src/lc3/registers"
)

/* update flag */
func update_flags(res uint16, regs interfaces.Registers) (err error) {
    if (res == 0) {
        return regs.Write(registers.R_COND, FL_ZRO)
    } else if res >> 15 == 1 { /* a 1 in the left-most bit indicates negative */
        return regs.Write(registers.R_COND, FL_NEG)
    }
    return regs.Write(registers.R_COND, FL_POS)
}

/* sign_extend */
func sign_extend(x uint16, bit_count uint16) uint16 {
    if (x >> (bit_count - 1)) & 1 == 1 {
        x |= (0xFFFF << bit_count);
    }
    return x;
}

/* function to read register and convert it to uint 16 */
func readReg(regs interfaces.Registers, r uint16) (value uint16) {
  v, err := regs.Read(r) /* read register */
  handle(err) /* throw an error */

  value, ok := v.(uint16) /* convert it to uint16 */
  if !ok { /* throw an error */
    panic("can't convert register to int")
  }

  return value
}

/* function to handle panics */
func handle(err error) {
  if err != nil {
    panic(err)
  }
}

/* function to read of off memory */
func readMem(mem interfaces.Memory, address uint16) (value uint16) {
  v, err := mem.Read(address) /* read value */
  handle(err) /* throw an error */

  value, ok := v.(uint16) /* convert it to uint16 */
  if !ok { /* throw an error */
    panic("can't convert register to int")
  }

  return value
}

/* convert any recover to error for return */
func recover_all() (err error) {
  r := recover()
  if r != nil {
    switch x := r.(type) {
    case string:
      return errors.New(x)

    case error:
      return x

    default:
      return errors.New("Unknown panic")
    }
  }

  return nil
}

// GetChar reads one character from Stdin as an uint16
func GetChar() uint16 {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	b := make([]byte, 1)
	os.Stdin.Read(b)
	return uint16(b[0])
}

// IsKeyPressed checks if a key was pressed
func IsKeyPressed() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Size() > 0
}
