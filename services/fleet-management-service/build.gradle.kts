plugins {
	java
	id("org.springframework.boot") version "4.0.2"
	id("io.spring.dependency-management") version "1.1.7"
}

group = "com.fleetpulse"
version = "0.0.1-SNAPSHOT"
description = "Fleet Management Service"

java {
	toolchain {
		languageVersion = JavaLanguageVersion.of(21)
	}
}

repositories {
	mavenCentral()
}

extra["springCloudVersion"] = "2025.1.0"

dependencies {
	implementation("org.springframework.boot:spring-boot-starter-webmvc")
	implementation("org.mybatis.spring.boot:mybatis-spring-boot-starter:4.0.1")
	implementation("org.springframework.cloud:spring-cloud-starter-config")
	implementation("org.springframework.cloud:spring-cloud-starter-netflix-eureka-client")
	runtimeOnly("org.postgresql:postgresql")
	testImplementation("org.springframework.boot:spring-boot-starter-webmvc-test")
	testImplementation("org.mybatis.spring.boot:mybatis-spring-boot-starter-test:4.0.1")
	testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

dependencyManagement {
	imports {
		mavenBom("org.springframework.cloud:spring-cloud-dependencies:${property("springCloudVersion")}")
	}
}

tasks.withType<Test> {
	useJUnitPlatform()
}
