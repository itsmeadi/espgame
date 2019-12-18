-- MySQL dump 10.13  Distrib 8.0.16, for osx10.13 (x86_64)
--
-- Host: localhost    Database: ESPGAME
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `answers`
--

DROP TABLE IF EXISTS `answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `answers` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `question_id` int(5) DEFAULT NULL,
  `answer_text` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `media_url` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answers`
--

LOCK TABLES `answers` WRITE;
/*!40000 ALTER TABLE `answers` DISABLE KEYS */;
INSERT INTO `answers` VALUES (1,1,'','./upload/9b36a7e0-a309-457b-8edb-50be0b6b2365'),(2,1,'','./upload/e663a00e-c692-4493-960c-a378a1081548'),(3,1,'','./upload/66a9cb7a-5880-4561-ad1f-348160d1e4ff'),(4,1,'','./upload/ff1e346a-b4b0-46bb-9ae4-df25a0253f38'),(5,1,'','./upload/74e5a698-1f14-4a59-a80c-cccf645cdab2'),(6,2,'','./upload/eb662c39-dae9-4ae2-807e-a3802ac2ad35'),(7,2,'','./upload/9c0c7238-57ca-4324-99b4-b389d6807b0d'),(8,2,'','./upload/e6597a5e-9e1f-450c-bef7-dcf9c9fa9581'),(9,2,'','./upload/c7ff514b-dded-4ecb-bd47-ef65c12e181f'),(10,2,'','./upload/bab2edf8-e0c4-462e-8b09-cf6f62890313'),(11,3,'','./upload/10ad284a-8dfa-4409-89cd-921ef8a51ef9'),(12,3,'','./upload/631d6a21-7c3f-49fb-87f8-808bc548b518'),(13,3,'','./upload/0f23ce3f-7431-4762-915d-d81af3d0b3ad'),(14,3,'','./upload/bfd1da9c-8656-4e6a-b13f-915688f0c4c3'),(15,3,'','./upload/a8641830-a668-4437-80f3-ec12ed7d792f'),(16,4,'','./upload/1770c86d-8824-4820-8781-836496be8408'),(17,4,'','./upload/47ce0169-0ceb-40c0-bd1e-0f7aaaf50884'),(18,4,'','./upload/2ab682e4-2b72-4a2a-8f3b-103355ed3504'),(19,4,'','./upload/dd4886bd-d3dc-433b-abc5-4542efccbad1'),(20,4,'','./upload/4a212617-7663-410d-8015-9185d74cce1b'),(21,5,'','./upload/d14e4c2f-cd32-4a03-9b09-796b29763972'),(22,5,'','./upload/6168e6bb-0d35-4950-ac92-504e398a95e7'),(23,5,'','./upload/9717ad8e-8923-4b7a-b4f0-6ca938e6349c'),(24,5,'','./upload/1d93e066-b765-4aef-a698-e0abe08ded3e'),(25,5,'','./upload/9e79a45b-cd40-470b-9fec-47264a511a11'),(26,6,'','./upload/8a975b90-29a9-4c97-8428-0e85c466d3cc'),(27,6,'','./upload/d4852801-bcd2-48ac-8aa7-2cad1261b488'),(28,6,'','./upload/1ae62ed4-ad22-483b-bc69-637e5b7852ba'),(29,6,'','./upload/44bc6769-fd55-4bc6-9c39-bb0e0df593c2'),(30,6,'','./upload/3a788a08-39e7-4223-ad1c-d261770a5c41'),(31,7,'','./upload/b51980d9-ee3e-49d2-8df4-33c610c0248a'),(32,7,'','./upload/98099280-b99b-422e-91ff-5540a990db90'),(33,7,'','./upload/5e4fd7ed-25dd-49bd-b3d0-279d0e32d7d9'),(34,7,'','./upload/77406a0c-13ba-4960-80de-5d54fbc3ffa0'),(35,7,'','./upload/1d1b2af1-2fde-45a9-a570-c96928fb131d'),(36,8,'','./upload/c66869a8-7506-473e-8a79-050f4a3e885d'),(37,8,'','./upload/97127acf-1e8a-44ef-af79-0f7d5ffcbaed'),(38,8,'','./upload/894c4d52-59a4-46a8-b1d2-e680b7a4e2c6'),(39,8,'','./upload/afa44306-d107-422d-9b55-5c85374f9451'),(40,8,'','./upload/52e545b8-d640-4f0c-acc1-ff4409901adc'),(41,9,'','./upload/0d3e5162-da7a-4093-819d-7069b63d8b41'),(42,9,'','./upload/8234dea0-b7de-414f-992b-bb32098d2000'),(43,9,'','./upload/65031d33-430b-4335-8ebd-7abbea5d73e2'),(44,9,'','./upload/da19daab-8d57-458b-aaf2-ad435afd725d'),(45,9,'','./upload/425a4cd0-b366-4ebb-af49-97b3715c8bd9'),(46,10,'','./upload/5d9501b0-020c-43fa-ac00-f2c192a7c4eb'),(47,10,'','./upload/921c2b5e-712f-466c-8b2e-db677b0df9a3'),(48,10,'','./upload/8078968a-d680-4598-8da8-999fff8db515'),(49,10,'','./upload/6e055d7b-8a5c-41aa-bdfe-e41b569afeeb'),(50,10,'','./upload/1015e93d-a0c3-4e5e-96a4-cc6292aa9ca9'),(51,11,'','./upload/f5e14ccb-3ee1-40a3-b2e1-b7f46cf27863'),(52,11,'','./upload/640a2b28-0764-49e0-9cf8-2a0615d5b10e'),(53,11,'','./upload/8e09fa0b-51d0-4a24-aeca-b5c49ba5b552'),(54,11,'','./upload/a36f8278-57c0-4415-ab59-6bee2d556d33'),(55,11,'','./upload/3c06bcea-d89e-4fb3-9a3f-c07ba994e239'),(56,12,'','./upload/a4e05188-6fac-4318-b35b-a80c1026aae4'),(57,12,'','./upload/9f7b4098-0d6d-4f6f-9ba0-abc1de2a5764');
/*!40000 ALTER TABLE `answers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `group` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) COLLATE utf8mb4_general_ci DEFAULT '0',
  `answered_by_users` int(2) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `questions`
--

DROP TABLE IF EXISTS `questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `questions` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `question_text` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `media_url` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `answered_by_users` int(2) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `questions`
--

LOCK TABLES `questions` WRITE;
/*!40000 ALTER TABLE `questions` DISABLE KEYS */;
INSERT INTO `questions` VALUES (1,'Match the images','./upload/edf561ef-c0e4-41d1-b8a2-c78d975037df',2),(2,'Match the images','./upload/0aebe25e-8b69-46bd-ae33-0bd7d292dd78',2),(3,'Match the images','./upload/dba18631-a35a-487a-ae71-e84c41bd4c27',2),(4,'Match the images','./upload/c457ce80-95e9-4065-9094-ee7325cdc7d4',2),(5,'Match the images','./upload/f20e2244-2884-49a9-95e8-db98aca0523b',2),(6,'Match the images','./upload/43f5f7ef-68ac-499e-87af-02e4eafdf7ab',0),(7,'Match the images','./upload/3378bae3-2512-4a0f-99fe-80547dd15573',0),(8,'Match the images','./upload/9f3b9569-3078-446f-8c88-0128236cc6f2',0),(9,'Match the images','./upload/8661c8d0-a753-4c28-85ab-f0cd0bfa0b3e',0),(10,'Match the images','./upload/0f47ed4e-d76a-470d-9faa-f07d4a9ff56d',0),(11,'Match the images','./upload/c51e04af-252e-4116-b510-4542c9cd72fd',0),(12,'Match the images','./upload/1d3f5dbd-31bc-4527-92b2-eed1b1e97056',0);
/*!40000 ALTER TABLE `questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `id` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `usertype` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('adi2','adi2','12cf356b0f0d6802815bfd02b84a9418',''),('aditya','aditya','057829fa5a65fc1ace408f490be486ac',''),('admin','admin','21232f297a57a5a743894a0e4a801fc3','admin');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_questions_answers`
--

DROP TABLE IF EXISTS `user_questions_answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_questions_answers` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `question_id` int(5) DEFAULT NULL,
  `answer_id` int(5) DEFAULT NULL,
  `correctness` int(2) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`user_id`,`question_id`,`answer_id`),
  KEY `q_idx` (`question_id`),
  KEY `uid_idx` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_questions_answers`
--

LOCK TABLES `user_questions_answers` WRITE;
/*!40000 ALTER TABLE `user_questions_answers` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_questions_answers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-12-19  2:47:17
