package colors

import "fmt"

func Red(str string) string {
	return fmt.Sprintf("\033[31m%v\033[0m", str)
}

func Green(str string) string {
	return fmt.Sprintf("\033[32m%v\033[0m", str)
}

func Yellow(str string) string {
	return fmt.Sprintf("\033[33m%v\033[0m", str)
}

func Blue(str string) string {
	return fmt.Sprintf("\033[34m%v\033[0m", str)
}

func Purple(str string) string {
	return fmt.Sprintf("\033[35m%v\033[0m", str)
}

func Cyan(str string) string {
	return fmt.Sprintf("\033[36m%v\033[0m", str)
}

func White(str string) string {
	return fmt.Sprintf("\033[37m%v\033[0m", str)
}

func RedBold(str string) string {
	return fmt.Sprintf("\033[31;1m%v\033[0m", str)
}

func GreenBold(str string) string {
	return fmt.Sprintf("\033[32;1m%v\033[0m", str)
}

func YellowBold(str string) string {
	return fmt.Sprintf("\033[33;1m%v\033[0m", str)
}

func BlueBold(str string) string {
	return fmt.Sprintf("\033[34;1m%v\033[0m", str)
}

func PurpleBold(str string) string {
	return fmt.Sprintf("\033[35;1m%v\033[0m", str)
}

func CyanBold(str string) string {
	return fmt.Sprintf("\033[36;1m%v\033[0m", str)
}

func WhiteBold(str string) string {
	return fmt.Sprintf("\033[37;1m%v\033[0m", str)
}
