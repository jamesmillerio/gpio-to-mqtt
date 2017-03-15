var fs = require("fs");

var noop = function() {};

exports = {
  open: function(pin, callback) {

    var gpioBase = "/sys/class/gpio";
    var pinPath = gpioBase + "/gpio" + pin;
    var direction = "out";

    callback = callback || noop;

    fs.open(gpioBase + "/export", "w", null, function(openError, fd) {

            if(openError) {
              callback(openError);
              return;
            }

            fs.write(fd, pin, function(writeError, w, s) {

                    fs.close(fd);

                    if(writeError) {
                      callback(writeError);
                      return;
                    }

                    fs.open(pinPath + "/direction", "w", null, function(directionError, fDir) {

                            if(directionError) {
                              callback(directionError);
                              return;
                            }

                            fs.write(fDir, direction, function(directionWriteError, w, s) {

                                    if(directionWriteError) {
                                      callback(directionWriteError);
                                      return;
                                    }

                                    fs.close(fDir);

                                    callback();

                            });

                    });

            });

    });

  },
  read: function(pin) {

  },
  close: function(pin, callback) {
    var gpioBase = "/sys/class/gpio";
    var pinValuePath = gpioBase + "/gpio" + pin + "/value";

    callback = callback || noop;

    fs.open(pinValuePath, "r", function(openError, fd) {

      if(openError) {
        callback(openError, false);
        return;
      }

      fs.readFile(fd, function(readError, data) {

        if(readError) {
          callback(readError, false);
          return;
        }

        console.log(data);

        callback(null, data == "1");

      });

    });


  }
};
