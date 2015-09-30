val file = new java.io.PrintWriter("out/testdata")
scala.util.Random.shuffle(0 to 999999).map(x => f"$x%07d").foreach(file.println)
file.close()
