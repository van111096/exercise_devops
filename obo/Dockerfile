
FROM maven:ibmjava-alpine as spring-boot-task7
COPY . .
RUN mvn package -Dmaven.test.skip=true

FROM openjdk:8-alpine
COPY --from=spring-boot-task7 /target/obo-stadium-0.0.1-SNAPSHOT.jar .
CMD java -Djava.security.egd=file:/dev/./urandom -jar obo-stadium-0.0.1-SNAPSHOT.jar