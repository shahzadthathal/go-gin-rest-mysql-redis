-- MySQL dump 10.13  Distrib 5.7.20, for Win64 (x86_64)
--
-- Host: localhost    Database: ms_account_dev
-- ------------------------------------------------------
-- Server version	5.7.20-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `department`
--

DROP TABLE IF EXISTS `department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `department` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) DEFAULT NULL,
  `WORKSTATION_ID` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `DEPARTMENT_WORKSTATION` (`WORKSTATION_ID`),
  CONSTRAINT `DEPARTMENT_WORKSTATION` FOREIGN KEY (`WORKSTATION_ID`) REFERENCES `workstation` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `department`
--

LOCK TABLES `department` WRITE;
/*!40000 ALTER TABLE `department` DISABLE KEYS */;
INSERT INTO `department` VALUES (1,'Sales & Marketing',1),(2,'Research & Development',1),(3,'Administration',1),(4,'Human Resources',2),(5,'Sales & Marketing',3);
/*!40000 ALTER TABLE `department` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user`
--

DROP TABLE IF EXISTS `m_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `m_user` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `PASSWORD` varchar(255) DEFAULT NULL,
  `USER_NAME` varchar(255) DEFAULT NULL,
  `ACCOUNT_EXPIRED` tinyint(1) DEFAULT NULL,
  `ACCOUNT_LOCKED` tinyint(1) DEFAULT NULL,
  `CREDENTIALS_EXPIRED` tinyint(1) DEFAULT NULL,
  `ENABLED` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `USER_USER_NAME` (`USER_NAME`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user`
--

LOCK TABLES `m_user` WRITE;
/*!40000 ALTER TABLE `m_user` DISABLE KEYS */;
INSERT INTO `m_user` VALUES (1,'$2a$08$qvrzQZ7jJ7oy2p/msL4M0.l83Cd0jNsX6AJUitbgRXGzge4j035ha','admin',0,0,0,1),(2,'$2a$08$dwYz8O.qtUXboGosJFsS4u19LHKW7aCQ0LXXuNlRfjjGKwj5NfKSe','budi',0,0,0,1),(3,'$2a$08$kPjzxewXRGNRiIuL4FtQH.mhMn7ZAFBYKB3ROz.J24IX8vDAcThsG','anduk',0,0,0,1),(4,'$2a$08$vVXqh6S8TqfHMs1SlNTu/.J25iUCrpGBpyGExA.9yI.IlDRadR6Ea','haya',0,0,0,1);
/*!40000 ALTER TABLE `m_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user_detail`
--

DROP TABLE IF EXISTS `m_user_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `m_user_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ADDRESS` varchar(255) DEFAULT NULL,
  `DOB` varchar(8) DEFAULT NULL,
  `POB` varchar(255) DEFAULT NULL,
  `PHONE` varchar(255) DEFAULT NULL,
  `EMAIL` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user_detail`
--

LOCK TABLES `m_user_detail` WRITE;
/*!40000 ALTER TABLE `m_user_detail` DISABLE KEYS */;
INSERT INTO `m_user_detail` VALUES (2,'Jl citapen','19950617','Bandung','081232342234','haer@gmail.com',1),(3,'Jl wiradanu','19941210','Jepara','0812234234234','hawr@gmail.com',2),(4,'Jl galungun','19970116','Garut','0812345674234','hasfg@gmail.com',3),(5,'Jl iskandar aa','19910817','Jakartaaaa','081234234234','haahffff@gmail.com',4);
/*!40000 ALTER TABLE `m_user_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user_role`
--

DROP TABLE IF EXISTS `m_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `m_user_role` (
  `USER_ID` bigint(20) DEFAULT NULL,
  `DEPARTMENT_ID` bigint(20) DEFAULT NULL,
  `POSITION` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user_role`
--

LOCK TABLES `m_user_role` WRITE;
/*!40000 ALTER TABLE `m_user_role` DISABLE KEYS */;
INSERT INTO `m_user_role` VALUES (1,1,'Manager'),(2,1,'Staff'),(3,3,'Staff'),(4,5,'Manager');
/*!40000 ALTER TABLE `m_user_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workstation`
--

DROP TABLE IF EXISTS `workstation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `workstation` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) DEFAULT NULL,
  `ADDRESS` varchar(255) DEFAULT NULL,
  `EMAIL` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `WORKSTATION_NAME` (`NAME`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workstation`
--

LOCK TABLES `workstation` WRITE;
/*!40000 ALTER TABLE `workstation` DISABLE KEYS */;
INSERT INTO `workstation` VALUES (1,'Jakarta','Mampang','Jakarta@gmail.com'),(2,'Bandung','Buah Batu','bandung@gmail.com'),(3,'Tasikmalaya','Citapen','tasikmalaya@gmail.com');
/*!40000 ALTER TABLE `workstation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'ms_account_dev'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-05 11:48:29
