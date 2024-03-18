#include <Serial.h>

const int flashlightPin = 9;  // D9 GPIO

void setup() {
    Serial.begin(9600);              // Start serial communication at 9600 baud rate
    pinMode(flashlightPin, OUTPUT);  // Set the control pin as output
    Serial.println("Starting test program...");
}

// Quick flash, used for alerts
void loop() {
    Serial.println("Turning on LED...");
    digitalWrite(flashlightPin, HIGH);
    delay(1);  // Wait for 1ms second
    Serial.println("Turning off LED...");
    digitalWrite(flashlightPin, LOW);
    delay(5000);  // Wait for 5000ms
}
