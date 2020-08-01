-- phpMyAdmin SQL Dump
-- version 4.7.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: 2019-08-29 18:54:33
-- 服务器版本： 5.6.35
-- PHP Version: 7.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `monteblog_beego`
--

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE `article` (
  `article_id` int(11) NOT NULL COMMENT '文章id',
  `article_class_id` int(11) NOT NULL COMMENT '分类id',
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `content` text NOT NULL COMMENT '文章内容',
  `time` int(11) NOT NULL COMMENT '发布时间',
  `update_time` int(11) NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章内容';

--
-- 转存表中的数据 `article`
--

INSERT INTO `article` (`article_id`, `article_class_id`, `title`, `content`, `time`, `update_time`) VALUES
(1, 1, 'PHP是不是世界上最好的语言', '                  <p>\n	PHP是不是世界上最好的语言？\n</p>\n<p>\n	1.是不是？\n</p>\n<p>\n	2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n<p>\n	PHP是不是世界上最好的语言？\n1.是不是？\n2.到底是不是？\n</p>\n                ', 1567042074, 1567042074),
(2, 2, '第一篇文章2', '第一篇文章内容', 0, 1),
(3, 1, '第一篇文章3', '第一篇文章内容', 0, 1),
(4, 1, '第一篇文章4', '第一篇文章内容', 0, 1),
(5, 1, '第一篇文章5', '第一篇文章内容', 0, 1),
(6, 1, '第一篇文章6', '第一篇文章内容', 0, 1),
(7, 1, '第一篇文章7', '第一篇文章内容', 0, 1),
(8, 1, '第一篇文章8', '第一篇文章内容', 0, 1),
(9, 1, '第一篇文章9', '第一篇文章内容', 0, 1),
(10, 1, '第一篇文章10', '第一篇文章内容', 0, 1),
(11, 1, '第一篇文章11', '第一篇文章内容', 0, 1),
(12, 1, '第一篇文章12', '第一篇文章内容', 0, 1),
(13, 1, '第一篇文章13', '第一篇文章内容', 0, 1),
(14, 1, '第一篇文章14', '第一篇文章内容', 0, 1),
(15, 1, '第一篇文章15', '第一篇文章内容', 0, 1),
(16, 1, '第一篇文章16', '第一篇文章内容', 0, 1),
(18, 0, '', '888', 0, 0),
(19, 4, '测试富文本图片上传', '<div align=\"center\">\n	<img src=\"/static/upload/8569787226622152728.png\" alt=\"\" /><br />\n</div>', 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `article_class`
--

CREATE TABLE `article_class` (
  `article_class_id` int(11) NOT NULL COMMENT '文章分类id',
  `name` varchar(255) NOT NULL COMMENT '分类名称'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类';

--
-- 转存表中的数据 `article_class`
--

INSERT INTO `article_class` (`article_class_id`, `name`) VALUES
(1, 'PHP'),
(2, 'Go语言'),
(3, 'Python'),
(4, '前端'),
(5, '运维');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `admin` int(11) NOT NULL DEFAULT '0' COMMENT '1管理员0普通成员'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`user_id`, `username`, `password`, `admin`) VALUES
(1, 'admin', '123456', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`article_id`);

--
-- Indexes for table `article_class`
--
ALTER TABLE `article_class`
  ADD PRIMARY KEY (`article_class_id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`user_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `article`
--
ALTER TABLE `article`
  MODIFY `article_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章id', AUTO_INCREMENT=20;
--
-- 使用表AUTO_INCREMENT `article_class`
--
ALTER TABLE `article_class`
  MODIFY `article_class_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章分类id', AUTO_INCREMENT=20;
--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id', AUTO_INCREMENT=2;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
