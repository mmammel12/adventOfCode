const { readFile } = require("../../readFile");

let data = [];
readFile(
    "2022\\day1\\input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

let max = 0;
let acc = 0;
data.forEach((value) => {
    if (value === "") {
        max = acc > max ? acc : max;
        acc = 0;
    } else {
        acc += parseInt(value);
    }
});

console.log(max);
