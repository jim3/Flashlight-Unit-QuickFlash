### M5Stack Flashlight Unit
---

The device has built-in flash, includes the AW3641 driver and a white LED, with a color temperature of 5000-5700K
The "white LED" really does have that nice "_camera flash effect_", totally worth the price of this device, and
turns are to be _perfect_ for an alert light that's triggered by an _event_ or _alarm_.

The communication interface is GPIO. I tested this on a `Arduino Uno R4 WiFi` (C++) and a `ESP32 DevKitM-1` (TinyGo).
I used GPIO 9 on the Arduino, GPIO 26 on the ESP32.

`M5Stack` > `ESP32 DevKitM-1`

| M5Stack GPIO | ESP32 GPIO |
| ------------ | ---------- |
| 5V           | 5V         |
| GND          | GND        |
| S            | GPIO 26    |

---
