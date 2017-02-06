name := """slick-mysql-app"""

version := "1.0-SNAPSHOT"

lazy val root = (project in file(".")).enablePlugins(PlayScala)

scalaVersion := "2.11.7"

libraryDependencies ++= Seq(
  jdbc,
  cache,
  ws,
  "org.scalatestplus.play" %% "scalatestplus-play" % "1.5.1" % Test,
  "com.typesafe.slick" %% "slick" % "3.1.0-M2",
  "com.typesafe.slick" %% "slick-codegen" % "3.1.0-M2",
  "mysql" % "mysql-connector-java" % "5.1.34",
)
