package UserInputValidator

import (
	"strings"
)

func Validateuserinput(firstname string, lastname string, emailid string, ticket_order uint, current_avail_ticket uint, mob_no string) (bool, bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidmail := strings.Contains(emailid, "@")
	isValidOrder := ticket_order >= 0 && ticket_order <= current_avail_ticket
	isValidMob := len((mob_no)) == 10
	return isValidMob, isValidName, isValidOrder, isValidmail
}
