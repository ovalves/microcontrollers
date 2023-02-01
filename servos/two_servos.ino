#include <Servo.h>

Servo servoL;
Servo servoR;

void setup() {
    // configurando a comunicação via porta serial com velocidade de 9600bps(baud).
    Serial.begin(9600);
    Serial.flush();

    // Anexa os servos (físico), nos pino 5 e 6
    servoL.attach(5);
    servoR.attach(6);

    delay(1000);

    // Coloca os servos na posição inicial
    servoL.write(0);
    servoR.write(0);

    Serial.println("STARTING SERVOS...");
}

void loop() {
    // Lê o valor do potenciômetro
    int angle = analogRead(0);

    // Mapeia os valores de 0 a 90 graus
    angle = map(angle, 0, 1023, 0, 90);

    Serial.println(angle);

    // Escreve o ângulo para os servos
    servoL.write(angle);
    servoR.write(angle);
}
