package example

import slick.driver.MySQLDriver.api._

object CodeGen extends App {
  val slickDriver = "com.mysql.jdbc.Driver"
  val jdbcDriver = "org.mysql.Driver"
  val url = "jdbc:mysql://localhost/skill_search"
  val outputDir = "src/main/scala"
  val pkg = "example"
  val user = "otsukano"
  val password = "lain1896"

  slick.codegen.SourceCodeGenerator.main(
    Array(slickDriver, jdbcDriver, url, outputDir, pkg, user, password))
}
