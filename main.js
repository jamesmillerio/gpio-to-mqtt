var gpio = require("pi-gpio"),
    Promise = require("promise"),
    config = require("./config.json");

if(config == null || config.pins == null) { return; }

var promises = [];

//Convert pins to promises.
for(var i = 0; i < config.pins.length; i++) {

  var pin = config.pins[i];

  promises.push(new Promise(function(resolve, reject) {
    gpio.open(pin.pin, null, function(err) {

      if(err) { console.log("Error opening pin: " + err); }

      resolve(pin)

    });
  }));

}

Promise.race(promises).then(function(pin) {
  console.log("PIN!");
  console.log(pin);
})
