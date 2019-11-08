package command

import (
	"fmt"
)


//AreYouSure to command line arg for confirmation
func AreYouSure(f func(), msg string) {
	var sure string

	fmt.Println("[?] Are you sure ... (y/n)")
	fmt.Scanf("%s", &sure)

	switch sure {
	case "Y", "y", "Yes", "yes", "YES":
		f()
		fmt.Println("[+] ", msg)
		break

	default:
		return
	}

}
