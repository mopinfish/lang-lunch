object TraitSample {
  def main(args: Array[String]): Unit = {
    val a = new Human("dwango")
    a.printName()
  }
}

trait NameTrait {
  val name: String
  def printName(): Unit = println(name)
}

class Human(val name: String) extends NameTrait

object HumanSample {
  val a = new Human("dwango")
  a.printName()
}
