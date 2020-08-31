"use strict"

const solver = (n, as) => {
  if (n === 0) return as
  else if (n === 1) {
    const p1 = as[0]
    const p2 = as[as.length - 1]
    const s = [ (2*p1[0] + p2[0]) / 3, (2*p1[1] + p2[1]) / 3 ]
    const t = [ (p1[0] + 2*p2[0]) / 3, (p1[1] + 2*p2[1]) / 3 ]
    const u = [ 
      (t[0]-s[0]) * 1/2            - (t[1]-s[1]) * Math.sqrt(3)/2 + s[0],
      (t[0]-s[0]) * Math.sqrt(3)/2 + (t[1]-s[1]) * 1/2            + s[1]
    ]
    return [p1, s, u, t, p2]
  } else {
    const ans = solver(n-1, as)
    const result = []
    for (var i = 0; i < ans.length-1; i++) {
      var l = solver(1, [ans[i], ans[i+1]])
      for (var j = 0; j < l.length-1; j++) {
        result.push(l[j])
      }
    }
    result.push(ans[ans.length-1])
    return result
  }
}

const reader = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
})
reader.on("line", (line) => {
  const n = Number(line.trim())
  const result = solver(n, [
    [0, 0],
    [100, 0]
  ])
  console.log(result.slice(1).reduce((acc, a) => `${acc}\n${a[0].toFixed(8)} ${a[1].toFixed(8)}`, `${result[0][0].toFixed(8)} ${result[0][1].toFixed(8)}`))
  process.exit(0)
})
