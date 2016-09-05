object LasyValSample {
  def main(args: Array[String]): Unit = {
    Def.printZ() // => 3
    Def.y = 4
    Def.printZ() // => 5
    Val.printZ() // => 3
    Val.y = 4
    Val.printZ() // => 3
  }
}
trait A {
  val foo: String
}

trait B extends A {
//  lazy val bar = foo + "World" // もしくは def bar でもよい
  def bar = foo + "World!!"
}

class C extends B {
  val foo = "Hello"

  def printBar(): Unit = println(bar)
}

object Def {
  val x = 1
  var y = 2
  def z = x + y
  def printZ():Unit = {
    println(z)
  }
}

object Val {
  val x = 1
  var y = 2
  val z = x + y
  def printZ():Unit = {
    println(z)
  }
}
