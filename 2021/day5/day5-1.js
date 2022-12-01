const { readFile } = require("../readFile");

const main = () => {
  let data = [];
  let lines = {};
  readFile(
    "day5\\input.txt",
    (fileData) => (data = fileData),
    (err) => console.log(err)
  );
  console.log(data);
  data.forEach((item) => {
    const points = item.split(" -> ");
    console.log(points);
  });
};

main();
