#### Install Kotlin compiler

[https://kotlinlang.org/docs/command-line.html](https://kotlinlang.org/docs/command-line.html) points to [https://github.com/JetBrains/kotlin/releases](https://github.com/JetBrains/kotlin/releases) where I downloaded the zip, placed in `/opt/kotlin/` and add the `bin` directory in the system `$PATH`.


#### Install Gradle

Followed the instructions of [https://docs.gradle.org/current/userguide/installation.html#ex-installing-manually](https://docs.gradle.org/current/userguide/installation.html#ex-installing-manually) to install gradle


#### Start a Kotlin project
Create a project using gradle
```sh
mkdir my-project
cd my-project
gradle init --type kotlin-application --dsl kotlin
```
Build and run
```sh
./gradlew build
./gradlew run
```
