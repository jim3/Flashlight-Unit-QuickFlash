### M5Stack Flashlight Unit
---

Initially, I had *no* idea how I was going to trigger the flash, but it turned out to be dead simple. The device has built-in flash, 
that includes and a white LED, with a color temperature of 5000-5700K. The "white LED" really does have that nice "_camera flash effect_", 
totally worth the price of this device. I plan to incorporate it into my current project as an alert for various events. Adding PWM for the 
varying brightness is also a must.

The communication interface is GPIO. I tested this on an `Arduino Uno R4 WiFi` (C++) and an `ESP32 DevKitM-1` (TinyGo).
I used GPIO 9 on the Arduino, GPIO 26 on the ESP32. 

This is the *initial* TinyGo "version". I just recently discovered [TinyGo](https://tinygo.org/) and I am extremely excited about it!

| M5Stack GPIO | ESP32 GPIO |
| ------------ | ---------- |
| 5V           | 5V         |
| GND          | GND        |
| S            | GPIO 26    |

---
