### M5Stack Flashlight Unit

The device has built-in flash, including AW3641 driver and a white LED, with a color temperature of 5000-5700K.

The "white LED" really does have that nice "_camera flash effect_", totally worth the price of the device and
turns are to be _perfect_ for an alert light that's triggered by an _event_ or _alarm_.

The communication interface is GPIO. I tested this on the Arduino Uno R4 WiFi and it works well. The connection is as follows:

M5Stack -> Arduino Uno R4 WiFi

```
-   5V -> 5V
-   GND -> GND
-   S -> GPIO 9
```

The LED is turned on when pin is `HIGH` and turned off when the GPIO 9 pin is `LOW`.

