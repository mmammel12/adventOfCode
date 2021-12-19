const fs = require("fs");

const readFile = (path, callback, errorCallback) => {
  try {
    const data = fs.readFileSync(path, "utf-8");
    callback(data.split("\r\n"));
  } catch (err) {
    errorCallback(err);
  }
};

const findMostCommonBit = (data, position) => {
  let mostCommon = { 1: 0, 0: 0 };
  data.forEach((bits) =>
    bits[position] === "1" ? mostCommon["1"]++ : mostCommon["0"]++
  );
  return mostCommon["1"] >= mostCommon["0"] ? "1" : "0";
};

const findElementsWithCommonBit = (data, bit, position) => {
  return data.filter((bits) => bits[position] === bit);
};

const findOxygenAndCO2 = (data) => {
  let oxygenList = [...data];
  let CO2List = [...data];

  let currentBitPosition = 0;
  while (oxygenList.length > 1 || CO2List.length > 1) {
    if (oxygenList.length > 1) {
      oxygenList = findElementsWithCommonBit(
        oxygenList,
        findMostCommonBit(oxygenList, currentBitPosition),
        currentBitPosition
      );
    }
    if (CO2List.length > 1) {
      CO2List = findElementsWithCommonBit(
        CO2List,
        findMostCommonBit(CO2List, currentBitPosition) === "1" ? "0" : "1",
        currentBitPosition
      );
    }
    currentBitPosition++;
  }

  return parseInt(oxygenList[0], 2) * parseInt(CO2List[0], 2);
};

const main = () => {
  let binaryData = [];
  let error = false;
  let errorMsg = "";
  readFile(
    "./2021/day3/input.txt",
    (data) => {
      binaryData = data;
    },
    (err) => {
      errorMsg = err;
      error = true;
    }
  );

  if (!error) {
    console.log(findOxygenAndCO2(binaryData));
  } else {
    console.log(errorMsg);
  }
};

main();
