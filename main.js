var gpio = require("rpi-gpio"),
    Promise = require("promise"),
    config = require("./config.json");

if(config == null || config.pins == null) { return; }

var promises = [];

//Convert pins to promises.
for(var i = 0; i < config.pins.length; i++) {

  var pin = config.pins[i];

  promises.push(new Promise(function(resolve, reject) {

    //Close the pin just in case it's open from a previous session.
    gpio.close(pin.pin);

    //Open the pin.
    gpio.open(pin.pin, "out", function(err) {

      if(err) throw err;

      resolve(pin);

    });
  }));

}

Promise.all(promises).then(function(pins) {

  if(pins == null) return;

  for(var i = 0; i < pins.length; i++) {

    var pin = pins[i];

    gpio.read(pin.pin, function(err, value) {

      if(err) throw err;

      console.log("Pin " + pin.pin + ": " + value);

    });

  }

}).then(function(pins) {

  for(var i = 0; i < pins.length; i++) {

    var pin = pins[i];

    gpio.close(pin.pin);

  }

});
