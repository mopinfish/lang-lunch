name := """slick3-sample"""

version := "0.1"

lazy val root = (project in file("."))

scalaVersion := "2.11.7"

libraryDependencies ++= Seq(
	"com.typesafe.slick" %% "slick" % "3.1.1",
	"org.slf4j" % "slf4j-nop" % "1.7.20",
	"mysql" % "mysql-connector-java" % "5.1.35"
)
//libraryDependencies ++= Seq(
//  "org.scalatest" %% "scalatest" % "2.2.4" % "test",
//  "com.jsuereth" %% "scala-arm" % "1.4",
//  "com.typesafe.slick" %% "slick" % "3.1.0-M2",
//  "com.typesafe.slick" %% "slick-codegen" % "3.1.0-M2",
//  "com.typesafe" % "config" % "1.3.0",
//  "mysql" % "mysql-connector-java" % "5.1.34"
//)
