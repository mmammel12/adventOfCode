const fs = require("fs");

const readFile = (path, callback, errorCallback) => {
  try {
    console.log(`${__dirname}\\${path}`);
    const data = fs.readFileSync(`${__dirname}\\${path}`, "utf-8");
    if (callback != null) callback(data.split("\r\n"));
  } catch (err) {
    if (errorCallback != null) errorCallback(err);
  }
};

exports.readFile = readFile;
