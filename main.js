var gpio = require("pi-gpio"),
    config = require("./config.json");

if(config == null) { return; }

if(config.pins == null || config.pins.legnth <= 0) {
  return;
}

var openPin = function(pin) {

  if(pin == null) return;

  gpio.open(pin.pin, null, function(err) {

    if(err) { console.log("Error opening pin: " + err); }

    console.log(pin);

  });

};

var configurePins = function(c) {
  return new Promise();
  /*for(var i = 0; i < c.pins.length; i++) {

    var pin = c.pins[i];

    openPin(pin);
  }*/


};

configurePins(config);
