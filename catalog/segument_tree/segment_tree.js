"use strict"

class SegmentTree {
  constructor(xs) {
    const n = xs.length
    this.n = n

    this.nodes = new Array(2*n - 1)
    const m = this.nodes.length

    // 葉に値をセット
    for (var i = 1; i <= n; i++) {
      this.nodes[m-i] = xs[n-i]
    }

    // 後ろから順に値をセット
    for (var k = m - n - 1; k >= 0; k--) {
      // 左の子
      const left = this.nodes[2 * k + 1]
      // 右の子
      const right = this.nodes[2 * k + 2]
      // 自分の値
      this.nodes[k] = Math.min(left, right)
    }
  }
  update(i, x) {
    // 葉の段にアクセス
    const start = (this.nodes.length - n)
    const k = start + i
    this.nodes[k] = x
    for (var j = Math.floor((k-1)/2); j >= 0; j = Math.floor((j-1)/2)) {
      // 左の子
      const left = this.nodes[2 * j + 1]
      // 右の子
      const right = this.nodes[2 * j + 2]
      // 自分の値
      const v = Math.min(left, right)
      this.nodes[j] = v
    }
  }
  query(a, b) {
    const start = this.nodes.length - n
    return this._query(start+a, start+b, 0, start, this.nodes.length)
  }
  _query(a, b, k, l, r) {
    // [a, b) と [l, r] が関係ない
    if (b <= l || a >= r) return Number.MAX_SAFE_INTEGER
    // [a, b) 内に収まってる
    else if (a <= l && r <= b) return this.nodes[k]
    // [a, b) と [l, r] が一部交わってる
    else {
      return Math.min(
        this._query(a, b, 2*k + 1, l, Math.floor((r+l)/2)),
        this._query(a, b, 2*k + 2, Math.floor((r+l)/2), r)
      )
    }
  }
}

/**
 * % node main.js
 * > 4
 * > 5
 * > 3
 * > 7
 * > 9
 *   xs:5,3,7,9
 *   SegmentTree { n: 4, nodes: [
 *       3, 3, 7, 5,
 *       3, 7, 9
 *     ] }
 *   update i=1 x=9
 *   SegmentTree { n: 4, nodes: [
 *       5, 5, 7, 5,
 *       9, 7, 9
 *     ] }
 *   query [0, 3]
 *   5
 */
let count = 0
let n = 0
const xs = []
const reader = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
})
reader.on("line", (line) => {
  if (count === 0) {
    n = Number(line.trim())
  } else {
    const x = Number(line.trim())
    xs.push(x)
    if (count === n) {
      console.log(`xs:${xs}`)
      const tree = new SegmentTree(xs)
      console.log(tree)
      console.log(`update i=1 x=9`)
      tree.update(1, 9)
      console.log(tree)
      const min = tree.query(0, 3)
      console.log(`query [0, 3]`)
      console.log(min)
      process.exit(0)
    }
  }
  count++
})
