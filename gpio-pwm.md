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


### fuer servo-motoren


- port fuer pwm konfigurieren

        gpio mode 1 pwm

- pwm auf mark:space ratio einstellen

        gpio pwm-ms

- clock teiler -> 19.2 MHz / 400 -> 48 kHz slots pro sekunde

        gpio pwmc 400

- ein pwm durchlauf besteht aus 1000 slots (range)

        gpio pwmr 1000

