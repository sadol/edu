// helper package for the main package above; contains constants and function

package greetings

import "time"

const GoodDay = "Good Day"
const GoodNight = "Good Night"

func IsAm() bool {
	return time.Now().Hour() < 12
}
