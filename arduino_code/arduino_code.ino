const int NUM_SLIDERS = 2;
const int silder1 = A0, silder2 = A1;
int global1 = 0, global2 = 0;
int sliderVal1 = 0, sliderVal2 = 0;

void setup() {
  pinMode(silder1, INPUT);
  pinMode(silder2, INPUT);
  Serial.begin(9600);
}

void loop() {
  sliderVal1 =  map(analogRead(silder1), 0, 1023, 0, 100);
  sliderVal2 =  map(analogRead(silder2), 0, 1023, 0, 100);

  if(global1 != sliderVal1 || global2 != sliderVal2){
    String builtString = String((int)sliderVal1) + "," + String((int)sliderVal2);
    Serial.println(builtString);
    global1 = sliderVal1;
    global2 = sliderVal2;
  }
  // Serial.println("none");
  delay(100);
}
