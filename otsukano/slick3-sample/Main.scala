object Main extends App {
println(Database.forConfig("skill-search-local"))
//  val db: JdbcBackend.DatabaseDef = Database.forConfig("skill-search-local")
//  import slick.driver.MySQLDriver.api._
//
//  val sql = sql"select * from user".as[(Int, String, String)]
//  val f = db.run(sql)
//
//  Await.result(f, Duration.Inf) foreach println
}
