"use strict"

/**
 * @param {number[]} ws
 * @param {number} p
 * @param {number} k
 */
const check = (ws, p, k) => {
  let sum = 0
  let kCount = 1
  let sumMax = 0
  for (let i = 0; i < ws.length; i++) {
    if (ws[i] > p) {
      // p大きさが足りない
      return -1
    }

    if (sum + ws[i] > p) {
      if (kCount === k) {
        // p大きさが足りない
        return -1
      } else {
        // 次のトラック
        kCount++
        sum = 0
      }
    }
    sum = sum + ws[i]

    // 1つのトラックあたりの最大積載量を更新
    if (sum > sumMax) sumMax = sum
  }
  // 荷物全部つめた

  // pをもっと小さくできるかもしれない
  return 1
}

/**
 * @param {number[]} ws
 * @param {number} k
 * @param {number} left
 * @param {number} right
 */
const search = (ws, k, left, right) => {
  if (left === right) return left
  let p = Math.floor((right + left) / 2)
  // console.log(`p:${p} left:${left} right:${right} k:${k}`)
  const result = check(ws, p, k)
  if (result < 0) return search(ws, k, p+1, right) // pは右側
  else if (result > 0) return search(ws, k, left, p) // pは左側
  else return p
}

let rows = 0
let n = 0
let k = 0
const ws = []
let sumW = 0
const reader = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
})
reader.on("line", (line) => {
  rows++
  if (rows === 1) {
    ;[n, k] = line
      .trim()
      .split(" ")
      .map((x) => Number(x))
  } else {
    const w = Number(line.trim())
    ws.push(w)
    sumW = sumW + w
    if (rows - 1 === n) {
      const result = search(ws, k, 1, sumW)
      console.log(result)
      process.exit(0)
    }
  }
})

