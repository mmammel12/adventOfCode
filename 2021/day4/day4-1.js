const { readFile } = require("../readFile");

const checkForWin = (board, draw) => {
  let win = false;
  let total = 0;
  for (let row = 0; row < board.length; row++) {
    win = board[row].every((item) => item === "x");
    if (win) break;
  }
  if (!win) {
    for (let i = 0; i < 5; i++) {
      let hits = 0;
      for (let j = 0; j < 5; j++) {
        if (board[j][i] === "x") {
          hits++;
          if (hits === 5) {
            win = true;
            break;
          }
        }
      }
    }
  }

  if (win) {
    board.forEach((row) =>
      row.forEach((item) => {
        if (item !== "x") {
          total += parseInt(item);
        }
      })
    );
    throw new Error(`${total * parseInt(draw)}`);
  }
};

const checkDraws = (draws, boards) => {
  draws.forEach((draw) => {
    Object.keys(boards).forEach((boardNum) => {
      boards[boardNum].forEach((row) => {
        const hitIndex = row.findIndex((num) => num === draw);
        if (hitIndex >= 0) {
          row[hitIndex] = "x";
          checkForWin(boards[boardNum], draw);
        }
      });
    });
  });
};

const main = () => {
  let draws = [];
  let boards = {};

  readFile(
    "day4\\input.txt",
    (fileData) => {
      draws = fileData.shift().split(",");
      fileData.shift();
      let board = [];
      let count = 0;
      fileData.forEach((line) => {
        if (line === "") {
          boards[count] = board;
          board = [];
          count++;
        } else {
          board.push(line.split(" "));
        }
      });
      boards[count] = board;
    },
    (err) => console.log(err)
  );

  try {
    checkDraws(draws, boards);
  } catch (err) {
    console.log(err);
  }
};

main();
