package cycle

import "github.com/zarexalvindaria/leftpad-cycle"

var DEFAULT_CHAR = ' '

func DoublePad(s string, i int) string {
	return leftpad.Format(s+s, i)
}
