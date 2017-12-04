# senseBoxPi

Use this to read and submit sensor values from sensors attached to a Raspberry Pi.

Measurements are read directly through sysfs files provided by the drivers included in the Industrial I/O and Hwmon kernel subsystems.

## Supported Sensors
| Type                             | Subsystem/Bus  | required Linux version | Raspbian Version |
|----------------------------------|----------------|------------------------|------------------|
| `bmp280` (Pressure)              | IIO/I2C        | 4.9.60                 | November 2017    |
| `hdc100x` (Humidity/Temperature) | IIO/I2C        | 4.9.60                 | November 2017    |
| `tsl4531` (Light intensity)      | IIO/I2C        | 4.9.60                 | November 2017    |
| `veml6070` (UV light intensity)  | IIO/I2C        | 4.9.60                 | November 2017    |
| `sht3x` (Humidity/Temperature)   | Hwmon/I2C      | 4.9.65                 | November 2017    |

## Prerequisites
Update your Raspbian. Edit your `/boot/config.txt`. Enable I2C by uncommenting `dtparam=i2c_arm=on`.

Add the required device tree overlay instructions. `bmp280` and `hdc100x` support an additional `addr` argument for changing the I2C address.
```
dtoverlay=i2c-sensor,hdc100x,addr=0x43
dtoverlay=i2c-sensor,bmp280,addr=0x77
dtoverlay=i2c-sensor,tsl4531
dtoverlay=i2c-sensor,veml6070
dtoverlay=i2c-sensor,sht3x
```

Poweroff your Pi, connect the sensors with the Pi. Check your sensors if they need 3.3V or 5V. (Standard senseBox sensors are fine with 5V). Wire the sensors SDA and SCL to the Pis SDA and SCL Pins ([usually 3 and 5](https://pinout.xyz/pinout/i2c#)). Wire 5V and GND.

Boot your pi.

Check if the iio devices are populated `$ ls /sys/bus/iio/devices/`. If the output says at least `iio:device0  iio:device1  iio:device2  iio:device3`, you're good to go! Hwmon devices are populated under `/sys/class/hwmon`.

## Installation

Download the latest version from [the releases page](https://github.com/sensebox/senseboxpi/releases)

### From source

Install [Go](https://golang.org/doc/install), then

 `$ go get -u github.com/sensebox/senseboxpi/cmd/senseboxpi`

## Running

`senseboxpi` reads its configuration from a json file. By default, it expects a json file called `senseboxpi_config.json` next to the `senseboxpi` binary.

### Commandline flags
```
Usage of senseboxpi:
  -c string
    	path of the configuration json (shorthand) (default "senseboxpi_config.json")
  -config string
    	path of the configuration json (default "senseboxpi_config.json")
  -csv-output string
    	path to file where measurements in csv format will be appended
  -offline
    	operate offline. Do not upload to server

```

### Configuration JSON
The root keys should be `_id`, `postDomain` with string values and `sensors` as array of objects. Each object in `sensors` has the keys `_id`, `sensorType` and `phenomenon` with string values.
Here is an example:
```json
{
  "_id": "5912f2f051d3460011f57fdd",
  "postDomain": "ingress.osem.vo1d.space",
  "sensors": [
    {
      "_id": "5912f2f051d3460011f57fde",
      "phenomenon": "temperature",
      "sensorType": "hdc100x"
    },
    {
      "_id": "5912f2f051d3460011f57fdf",
      "phenomenon": "humidity",
      "sensorType": "hdc100x"
    },
    {
      "_id": "5912f2f051d3460011f57fe0",
      "phenomenon": "pressure",
      "sensorType": "bmp280"
    },
    {
      "_id": "5912f2f051d3460011f57fe1",
      "phenomenon": "light",
      "sensorType": "tsl4531"
    },
    {
      "_id": "5912f2f051d3460011f57fe2",
      "phenomenon": "uv",
      "sensorType": "veml6070"
    }
  ]
}
```
