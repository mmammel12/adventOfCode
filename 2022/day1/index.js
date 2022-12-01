const { readFile } = require("../../readFile");

let data = [];
readFile(
    "2022\\day1\\input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

let heap = [];
let acc = 0;
data.forEach((value) => {
    if (value === "") {
        if (heap.length < 3) {
            heap.push(acc);
        } else {
            if (acc > heap[0]) {
                heap[0] = acc;
            }
        }
        heap.sort();
        acc = 0;
    } else {
        acc += parseInt(value);
    }
});

const max = heap.reduce((acc, num) => (acc += num), 0);

console.log(max);
