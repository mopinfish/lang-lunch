class Point(val x: Int, val y: Int) {
  def +(p: Point): Point = {
    new Point(x + p.x, y + p.y)
  }
  override def toString(): String = "(" + x + ", " + y + ")"
}

class Calculator {
  def add(x: Int)(y: Int): Int = x + y
}

object HelloWorld {
  def main(args: Array[String]): Unit = {
//    isInfant()
//    count9()
//    forloop()
//    findTriangle()
//    sampleMatchSyntax()
//    listPatternMatch()
//    randomString()
    createPointClass()
    addSample()
    inheritSample()
  }

  def inheritSample():Unit = {
    class Aprinter() {
      def print(): Unit = {
        println("A")
      }
    }
    class Bprinter() extends Aprinter {
      override def print(): Unit = {
        println("B")
      }
    }
    new Aprinter().print
    new Bprinter().print
  }

  def addSample():Unit = {
    val calc = new Calculator()
    println(calc.add(2)(3))
    val calc2 = calc.add(2) _
    println(calc2)
    println(calc2(3))
  }

  def createPointClass(): Unit = {
    val p1 = new Point(36, 34)
    val p2 = new Point(2, 2)
    println(p1.toString())
    println(p1 + p2 + p2)
  }

  def randomString(): Unit = {
    for(i <- 1 to 1000) {
      val str = new scala.util.Random(new java.security.SecureRandom()).alphanumeric.take(5).toList match {
        case List(a, b, c, d, _) => List(a, b, c, d, a).mkString
      }
      println(str)
    }
  }

  def listPatternMatch(): Unit = {
    val arr = List("A", "B", "C", "D", "E")
    arr match {
      case List(a, "B", c, d, e) =>
        println("a = " + a)
        println("c = " + c)
        println("d = " + d)
        println("e = " + e)
      case _ =>
        println("nothing")
    }
  }

  def sampleMatchSyntax(): Unit = {
    val taro = "Taro"
    taro match {
      case "Taro" => println("Male")
      case "Jira" | "Baro" => println("BeMale")
      case "Hanako" => println("Female")
      // wild card: it is mean 'default' in other langages.
      case _ => println("Other")
    }
  }

  def findTriangle(): Unit = {
    for(a <- 1 to 1000; b <- 1 until 1000; c <- 1 to 1000; if (a * a == b * b + c * c)) {
      println("a = " + a + " b = " + b + " c = " + c)
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
