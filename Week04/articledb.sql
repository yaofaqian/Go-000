/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : articledb

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2020-12-16 20:31:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for article_content
-- ----------------------------
DROP TABLE IF EXISTS `article_content`;
CREATE TABLE `article_content` (
  `article_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `article_title` varchar(255) NOT NULL,
  `article_author` varchar(255) DEFAULT NULL,
  `article_publish_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `article_hits` int(11) DEFAULT '0',
  `article_content` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `user_id` int(11) DEFAULT NULL,
  `article_create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`article_id`)
) ENGINE=MyISAM AUTO_INCREMENT=129 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article_content
-- ----------------------------
INSERT INTO `article_content` VALUES ('128', '苹果M1芯片相较于华为麒麟9000芯片是否被严重高估了？', '互联网砖瓦匠', '2020-12-15 20:57:22', '0', '虽然还没有拿到M1实机，但只要对比华为最新的手机芯片——麒麟9000 就清楚了，麒麟9000作为一款拥有153亿晶体管的5nm处理器，对比苹果M1芯片160亿晶体管，M1仅仅只比麒麟多了4%晶体管。这样看下来，M1最多也就是一款手机级别的芯片，放在电脑上很难撑得起庞大的应用内容，如今网上一片人云测评觉得M1能吊打intel，是否有高估的嫌疑？？\r\n\r\n退一步来说，如果像M1这种性能的芯片都能轻松吊打intel，那么华为下一步是否也可以自研笔记本端的arm芯片，以此进一步占领笔记本领域，吞并苹果的市场份额？？', '1', '2020-12-15 20:57:22');
