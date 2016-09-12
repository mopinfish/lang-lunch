// build.sbt

lazy val root = (project in file(".")).
  settings(
    name := "sandbox",
    version := "0.0.1",
		scalaVersion := "2.11.8",
		scalacOptions ++= Seq("-deprecation", "-feature", "-unchecked", "-Xlint")
  )
