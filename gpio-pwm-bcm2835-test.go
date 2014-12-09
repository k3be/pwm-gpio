package main

//
// example golang code to control pwm per bcm2835 library
//
//
//
// depends at compile time on bcm2835 lib from: http://www.airspayce.com/mikem/bcm2835/
// pwm example in c: http://www.airspayce.com/mikem/bcm2835/pwm_8c-example.html
//
// compile:
//   * install lib bcm2835 (./configure && make &&  make install)
//   * compile this file (go build)
//
// execute:
//   * ./gpio-pwm-bcm2835-test <FROM> <TO> [<STEP_WIDTH>] [<DELAY>]
//

// #cgo LDFLAGS: -lbcm2835
// #include <bcm2835.h>
import "C"
import "os"
import "fmt"
import "strconv"

const (
	PWM_CHANNEL = 0
	RANGE       = 1024
	// http://www.airspayce.com/mikem/bcm2835/group__constants.html#ga63c029bd6500167152db4e57736d0939
	PIN = C.RPI_V2_GPIO_P1_12 // Version 2, Pin P1-12, can be PWM channel 0 in ALT FUN 5.
)

func main() {
	//
	// parse args
	//
	if len(os.Args) < 3 {
		fmt.Printf("%s: <FROM> <TO> [<STEP_WIDTH>] [<DELAY>]\n", os.Args[0])
		fmt.Printf("       FROM: start value\n")
		fmt.Printf("         TO: end value\n")
		fmt.Printf(" STEP_WIDTH: step width - from ... to (default 10)\n")
		fmt.Printf("      DELAY: step delay in millis (default 1000)\n")
		os.Exit(1)
	}

	from, err := strconv.Atoi(os.Args[1])
	failOnErr(err, "invalid <FROM>")

	to, err := strconv.Atoi(os.Args[2])
	failOnErr(err, "invalid <TO>")

	stepWidth := 10
	// optional: stepWidth
	if len(os.Args) >= 4 {
		stepWidth, err = strconv.Atoi(os.Args[3])
		failOnErr(err, "invalid <STEP_WIDTH>")
	}

	delayMillis := 1000
	// optional: delayMillis
	if len(os.Args) >= 5 {
		delayMillis, err = strconv.Atoi(os.Args[4])
		failOnErr(err, "invalid <DELAY>")
	}

	//
	// setup bcm2835 for pwm
	//
	if C.bcm2835_init() == 0 {
		// bcm2835_init prints the failure reason on stderr
		fmt.Println("bcm2835 init error")
		os.Exit(1)
	}
	defer C.bcm2835_close()

	// mode pwm
	// http://www.airspayce.com/mikem/bcm2835/group__gpio.html#gaf866b136c0a9fd4cca4065ce51eb4495
	C.bcm2835_gpio_fsel(PIN, C.BCM2835_GPIO_FSEL_ALT5)

	// clock divider
	// http://www.airspayce.com/mikem/bcm2835/group__pwm.html#ga4487f4e26e57ea3697a57cf52b8de35b
	C.bcm2835_pwm_set_clock(C.BCM2835_PWM_CLOCK_DIVIDER_512)

	// markspace
	// http://www.airspayce.com/mikem/bcm2835/group__pwm.html#ga7e2eddac472b6d81e0ba7dbf165672cb
	C.bcm2835_pwm_set_mode(PWM_CHANNEL, 1, 1)

	// range
	// http://www.airspayce.com/mikem/bcm2835/group__pwm.html#ga9f5ca04d4f859d96f1205a03c79de7ce
	C.bcm2835_pwm_set_range(PWM_CHANNEL, RANGE)

	//
	// action
	//
	fmt.Printf("loop from: %d to: %d with step width: %d and delay: %d ms\n", from, to, stepWidth, delayMillis)
	for d := from; d <= to; d += stepWidth {
		fmt.Printf("send %d\n", d)
		C.bcm2835_pwm_set_data(PWM_CHANNEL, C.uint32_t(d))

		C.bcm2835_delay(C.uint(delayMillis))
	}
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err.Error())
		os.Exit(1)
	}
}
