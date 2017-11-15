# senseBoxPi

Use this to read and submit sensor values from sensors attached to a Raspberry Pi.

Measurements are read directly through sysfs filed provided by the drivers included in the Industrial I/O kernel subsystem.

Tested with Raspbian Strech Lite with Kernel 4.9.60. Kernel Versions below 4.9.60 lack the needed driver modules.

## Prerequisites
Update your Raspbian to at least Kernel 4.9.60. You can check your current version with `$ uname -r`.
Edit your `/boot/config.txt`. Enable I2C by uncommenting `dtparam=i2c_arm=on`.


For the standard senseBox sensors, add
```
dtoverlay=i2c-sensor,hdc100x,addr=0x43
dtoverlay=i2c-sensor,bmp280,addr=0x77
dtoverlay=i2c-sensor,tsl4531
dtoverlay=i2c-sensor,veml6070
```

Poweroff your Pi, connect the sensors with the Pi. Check your sensors if they need 3.3V or 5V. (Standard senseBox sensors are fine with 5V). Wire the sensors SDA and SCL to the Pis SDA and SCL Pins ([usually 3 and 5](https://pinout.xyz/pinout/i2c#)). Wire 5V and GND.

Boot your pi.

Check if the iio devices are populated `$ ls /sys/bus/iio/devices/`. If the output says at least `iio:device0  iio:device1  iio:device2  iio:device3`, you're good to go!

## Installation

Install [Go](https://golang.org/doc/install), then

 `$ go get -u github.com/sensebox/senseboxpi`

## Supported Sensors
- BMP280 (Pressure)
- HDC100x (Temperature and Humidity)
- TSL4531 (Light intensity)
- VEML6070 (UV intensity)

