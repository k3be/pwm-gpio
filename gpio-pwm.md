# gpio

- zum konfigurieren / ansteuern der gpio's: [WiringPi](http://wiringpi.com/wiringpi-and-the-raspberry-pi-model-b/)

# port uebersicht

        root@ks:/home/pi# gpio readall
        +-----+-----+---------+------+---+-Model B2-+---+------+---------+-----+-----+
        | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
        +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
        |     |     |    3.3v |      |   |  1 || 2  |   |      | 5v      |     |     |
        |   2 |   8 |   SDA.1 |   IN | 1 |  3 || 4  |   |      | 5V      |     |     |
        |   3 |   9 |   SCL.1 |   IN | 1 |  5 || 6  |   |      | 0v      |     |     |
        |   4 |   7 | GPIO. 7 |   IN | 0 |  7 || 8  | 1 | ALT0 | TxD     | 15  | 14  |
        |     |     |      0v |      |   |  9 || 10 | 1 | ALT0 | RxD     | 16  | 15  |
        |  17 |   0 | GPIO. 0 |   IN | 0 | 11 || 12 | 0 | ALT5 | GPIO. 1 | 1   | 18  |
        |  27 |   2 | GPIO. 2 |   IN | 0 | 13 || 14 |   |      | 0v      |     |     |
        |  22 |   3 | GPIO. 3 |   IN | 0 | 15 || 16 | 0 | IN   | GPIO. 4 | 4   | 23  |
        |     |     |    3.3v |      |   | 17 || 18 | 0 | IN   | GPIO. 5 | 5   | 24  |
        |  10 |  12 |    MOSI |   IN | 0 | 19 || 20 |   |      | 0v      |     |     |
        |   9 |  13 |    MISO |   IN | 0 | 21 || 22 | 0 | IN   | GPIO. 6 | 6   | 25  |
        |  11 |  14 |    SCLK |   IN | 0 | 23 || 24 | 0 | IN   | CE0     | 10  | 8   |
        |     |     |      0v |      |   | 25 || 26 | 0 | IN   | CE1     | 11  | 7   |
        +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
        |  28 |  17 | GPIO.17 |   IN | 0 | 51 || 52 | 0 | IN   | GPIO.18 | 18  | 29  |
        |  30 |  19 | GPIO.19 |   IN | 0 | 53 || 54 | 0 | IN   | GPIO.20 | 20  | 31  |
        +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
        | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
        +-----+-----+---------+------+---+-Model B2-+---+------+---------+-----+-----+



## pwm

- beim pi nur ein port fuer pwm: GPIO 1 (phys port: 12)
- default takt 19.2 MHz

### BCM2835 library

  http://www.airspayce.com/mikem/bcm2835/index.html
   

>  The BCM2835 supports hardware PWM on a limited subset of GPIO pins. This bcm2835 library provides functions for configuring and controlling PWM output on these pins.
>
>  The BCM2835 contains 2 independent PWM channels (0 and 1), each of which be connnected to a limited subset of GPIO pins. The following GPIO pins may be connected to the following PWM channels (from section 9.5):
>  * GPIO PIN RPi pin PWM Channel ALT FUN
>  * 12 0 0
>  * 13 1 0
>  * 18 1-12 0 5
>  * 19 1 5
>  * 40 0 0
>  * 41 1 0
>  * 45 1 0
>  * 52 0 1
>  * 53 1 1
>  *
>
>  In order for a GPIO pin to emit output from its PWM channel, it must be set to the Alt Function given above. Note carefully that current versions of the Raspberry Pi only expose one of these pins (GPIO 18 = RPi Pin 1-12) on the IO headers, and therefore this is the only IO pin on the RPi that can be used for PWM. Further it must be set to ALT FUN 5 to get PWM output.
>
>  Both PWM channels are driven by the same PWM clock, whose clock dvider can be varied using bcm2835_pwm_set_clock(). Each channel can be separately enabled with bcm2835_pwm_set_mode(). The average output of the PWM channel is determined by the ratio of DATA/RANGE for that channel. Use bcm2835_pwm_set_range() to set the range and bcm2835_pwm_set_data() to set the data in that ratio
>
>  Each PWM channel can run in either Balanced or Mark-Space mode. In Balanced mode, the hardware sends a combination of clock pulses that results in an overall DATA pulses per RANGE pulses. In Mark-Space mode, the hardware sets the output HIGH for DATA clock pulses wide, followed by LOW for RANGE-DATA clock pulses.
>
>  The PWM clock can be set to control the PWM pulse widths. The PWM clock is derived from a 19.2MHz clock. You can set any divider, but some common ones are provided by the BCM2835_PWM_CLOCK_DIVIDER_* values of bcm2835PWMClockDivider.
>
>  For example, say you wanted to drive a DC motor with PWM at about 1kHz, and control the speed in 1/1024 increments from 0/1024 (stopped) through to 1024/1024 (full on). In that case you might set the clock divider to be 16, and the RANGE to 1024. The pulse repetition frequency will be 1.2MHz/1024 = 1171.875Hz.
>

  
### fuer servo-motoren


- port fuer pwm konfigurieren

        gpio mode 1 pwm

- pwm auf mark:space ratio einstellen

        gpio pwm-ms

- clock teiler -> 19.2 MHz / 400 -> 48 kHz slots pro sekunde

        gpio pwmc 400

- ein pwm durchlauf besteht aus 1000 slots (range)

        gpio pwmr 1000



#### servo werte

*per test programm: gpio-pwm-bcm2835-test - clock: 512, range: 1024*

|Servo       |MIN|MAX|
|:-----------|--:|--:|
|AMAX AM5811G| 20| 80|