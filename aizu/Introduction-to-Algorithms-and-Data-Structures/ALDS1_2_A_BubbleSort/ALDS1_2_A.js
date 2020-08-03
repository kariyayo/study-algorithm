"use strict"

/**
 * @param {any[]} xs
 */
const printArray = (xs) => {
  console.log(xs.join(" "))
}

/**
 * @param {number} n
 * @param {any[]} xs
 */
const bubbleSort = (n, xs) => {
  for (let i = 0; i < xs.length; i++) {
    for (let j = xs.length-1; j > i; j--) {
      if (xs[j] < xs[j-1]) {
        const tmp = xs[j-1]
        xs[j-1] = xs[j]
        xs[j] = tmp
        count++
      }
    }
  }
}

let count = 0

let n = -1
const xs = []
const reader = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});
reader.on('line', function (line) {
  if (n === -1) {
    n = Number(line)
  } else {
    const xs = line.split(" ").map(x => Number(x))
    bubbleSort(n, xs)
    printArray(xs)
    console.log(count)
    process.exit(0)
  }
})
