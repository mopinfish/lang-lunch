object HelloWorld {
  def main(args: Array[String]): Unit = {
//    isInfant()
//    count9()
//    forloop()
    findTriangle()
  }

  def findTriangle(): Unit = {
    for(a <- 1 to 1000; b <- 1 until 1000; c <- 1 to 1000) {
      if (a * a == b * b + c * c) {
        println("a = " + a + " b = " + b + " c = " + c)
      }
    }
  }

  def forloop(): Unit = {
    for(x <- 1 to 5; y <- 1 until 5; z <- 2 to 4; if x != y) {
      println("x = " + x + " y = " + y + " z = " + z) 
    }
  }

  def count9(): Unit = {
    var i = 0;
    do {
      println(i)
      i += 1
    } while(i < 10)
  }

  def isInfant(): Unit = {
    var age: Int = 7
    var isSchoolStarted: Boolean = false

    if (1 <= age && age <= 6) {
      isSchoolStarted = false
    } else {
      isSchoolStarted = true
    }

    if (isSchoolStarted) {
      println("幼児ではありません")
    } else {
      println("幼児です")
    }
  }
}
