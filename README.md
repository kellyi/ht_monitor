ht_monitor
==========

Golang application to broadcast humidity & temperature readings from a Raspberry Pi across MQTT. Requires building the [Sense HAT (B)'s SHTC3 example](https://www.waveshare.com/wiki/Sense_HAT_(B)).

[Here's a related Rails application](https://github.com/kellyi/humidity-monitor) which receives readings across MQTT and sends text messages if the humidity value falls outside a certain range.
