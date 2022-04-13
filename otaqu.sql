/*
 Navicat Premium Data Transfer

 Source Server         : local mysql
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : otaqu

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 13/04/2022 15:56:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kupon_prices
-- ----------------------------
DROP TABLE IF EXISTS `kupon_prices`;
CREATE TABLE `kupon_prices` (
  `uuid` char(36) NOT NULL,
  `kupon_id` char(36) NOT NULL,
  `name` varchar(200) NOT NULL DEFAULT '0',
  `price` int(11) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uuid`),
  KEY `kupon_id` (`kupon_id`),
  CONSTRAINT `FK__kupon` FOREIGN KEY (`kupon_id`) REFERENCES `kupons` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of kupon_prices
-- ----------------------------
BEGIN;
INSERT INTO `kupon_prices` VALUES ('13385d20-55d1-467d-a784-51312b01dbd4', '729f4175-9320-4de2-bcc5-49fb26ee369c', '0', 35000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('16d7bc9b-40ec-4775-b9b3-def5b44d4f49', 'e5c8d854-cd28-4e4c-918f-858009d791e1', '0', 24000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('1ec2f91d-fa24-482b-85de-27909d247d0c', 'a2585868-c8f5-414b-95d1-92948bdea1aa', '0', 99000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('4386a5f1-f8b3-4f77-b40c-4a3ba883c409', 'd87b121c-93ee-461b-b6da-84f8cd47b977', '0', 65000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('4c68a13a-45a3-4367-bc24-27c20e625f48', '8748fe96-356c-4ef9-81d4-7b8b36379680', '0', 40000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('61216e00-ff6f-458f-8bff-a94cf2c198d3', '2cab5a7e-69d0-441d-9a4f-136e661b604c', '0', 42000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('c0c9edcc-5538-4b0b-b09a-ba790eb14e10', 'd6249247-c33e-4aa7-94ef-fc46790386b0', '0', 27000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('cce0f6d9-5184-43f3-83ba-3ea191d0b7cb', 'dfba4b2d-5350-4e9b-9b02-7025308693c1', '0', 28000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('d8dfd823-2c46-4ebd-927c-210ebbf8878f', '838eb245-4a61-4bb5-ae40-307fe7bfb46a', '0', 65000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('dcd44ee8-637c-4b6b-9773-bc279b5d5f6d', '82d4daf0-7058-4109-a33a-8a18a835f53e', '0', 18000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('e9043152-96d5-41fe-b7cf-b5821cb471df', 'da7b239a-58e2-4535-9e7d-d163aec57aa3', '0', 15000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupon_prices` VALUES ('ec131b50-a3b6-4739-ad4e-d7238b31815b', '7c15a595-f5ff-40f9-b82a-d14511c15879', '0', 65000, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
COMMIT;

-- ----------------------------
-- Table structure for kupons
-- ----------------------------
DROP TABLE IF EXISTS `kupons`;
CREATE TABLE `kupons` (
  `uuid` char(36) NOT NULL,
  `name` varchar(200) NOT NULL DEFAULT '0',
  `description` text NOT NULL,
  `rating` decimal(2,1) NOT NULL DEFAULT '0.0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of kupons
-- ----------------------------
BEGIN;
INSERT INTO `kupons` VALUES ('2cab5a7e-69d0-441d-9a4f-136e661b604c', 'Transera Waterpark', 'Dapatkan Harga Menarik Untuk Bermain Wahana Air Bersama Keluarga Dan Teman di  Transera Waterpark', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('729f4175-9320-4de2-bcc5-49fb26ee369c', 'Amsterdam Waterpark', 'Ayo Liburan Murah dan Juga Seru Dengan Promo Tiket Masuk Amsterdam Waterpark 40 ribu!!', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('7c15a595-f5ff-40f9-b82a-d14511c15879', 'Taman Legenda Keong Emas TMII', 'Liburan Seru Dengan Bermain Wahana Menyenangkan di Taman Legenda Keong Emas TMII', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('82d4daf0-7058-4109-a33a-8a18a835f53e', 'Lembah Cisadane', 'Liburan bersama keluarga di Lembah Cisadane dengan suasana yang asri.', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('838eb245-4a61-4bb5-ae40-307fe7bfb46a', 'Ohana Waterpark', 'Dapatkan special promo Ohana Waterpark untuk weekdays dan weekend!', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('8748fe96-356c-4ef9-81d4-7b8b36379680', 'Venetian Water Carnaval', 'Liburan seru ala Kota Venezia Italia di Venetian Water Carnaval dengan promo Lakupon.', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('a2585868-c8f5-414b-95d1-92948bdea1aa', 'AYCE Namba-Ya! Suki and Grill', 'Nikmati All You Can Eat BBQ dan Shabu-Shabu khas Jepang di Namba-Ya dengan promo Lakupon.', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('d6249247-c33e-4aa7-94ef-fc46790386b0', 'Batavia Splash Water Adventure', 'Liburan Murah dan Seru bersama Keluarga di Batavia Splash Water Adventure', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('d87b121c-93ee-461b-b6da-84f8cd47b977', 'Outback Steakhouse Jakarta', 'Nikmati Steak Terbaik Super Juicy dan Empuk di Outback Steakhouse Makin Murah Dengan Voucher Senilai 200 ribu ini.', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('da7b239a-58e2-4535-9e7d-d163aec57aa3', 'Tiket Masuk Ancol', 'Liburan di Taman Impian Jaya Ancol. Dapatkan Harga Promo Tiket Masuk Ancol Termurah di Lakupon!', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('dfba4b2d-5350-4e9b-9b02-7025308693c1', 'Green Lake View Waterpark', 'Liburan Bersama Keluarga', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
INSERT INTO `kupons` VALUES ('e5c8d854-cd28-4e4c-918f-858009d791e1', 'Newtown Waterpark', 'Yuk Berlibur Seru Bermain Air Bersama Keluarga Lewat Promo Menarik di Newtown Waterpark!', 9.5, '2022-04-13 15:56:01', '2022-04-13 15:56:01');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
