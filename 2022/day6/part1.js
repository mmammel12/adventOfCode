const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day6/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);
const input = data[0];

let chars = "";
let num = 1;
for (let i = 0; i < input.length; i++) {
    const char = input[i];
    chars += char;
    if (chars.length === 5) {
        chars = chars.slice(1);
    }
    if (chars.length === 4) {
        const set = new Set(chars.split(""));
        if (set.size === 4) {
            num += i;
            break;
        }
    }
}

console.log(num);
