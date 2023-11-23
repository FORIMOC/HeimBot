-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-08-08 02:50:18
-- 服务器版本： 5.7.34-log
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `heimbot`
--

-- --------------------------------------------------------

--
-- 表的结构 `admins`
--

CREATE TABLE `admins` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(191) COLLATE utf8mb4_bin NOT NULL,
  `password` longtext COLLATE utf8mb4_bin NOT NULL,
  `level` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `admins`
--

INSERT INTO `admins` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `password`, `level`) VALUES
(1, '2023-08-07 16:32:07.498', '2023-08-07 16:32:07.498', NULL, '2097517935', '$2a$10$SWkiIQYj8ImwSYlQ0LbsROByjbTB5jzP0O1WrbZM2IlF5rZ03WpVC', 0);

-- --------------------------------------------------------

--
-- 表的结构 `fd_cache_messages`
--

CREATE TABLE `fd_cache_messages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(2047) COLLATE utf8mb4_bin NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `fd_cache_messages`
--

INSERT INTO `fd_cache_messages` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id`, `username`, `content`) VALUES
(67, '2023-08-07 16:37:49.985', '2023-08-07 16:37:49.985', NULL, '104', '10012', 'user12', '那也是玉兰卡');

-- --------------------------------------------------------

--
-- 表的结构 `fd_groups`
--

CREATE TABLE `fd_groups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `sender_id` varchar(255) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `fd_groups`
--

INSERT INTO `fd_groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `sender_id`) VALUES
(1, '2023-07-01 16:32:45.772', '2023-08-07 16:32:45.772', NULL, '839784935', '10009'),
(2, '2023-07-01 16:32:46.776', '2023-08-07 16:32:45.776', NULL, '519843267', '10010');

-- --------------------------------------------------------

--
-- 表的结构 `fd_messages`
--

CREATE TABLE `fd_messages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(2047) COLLATE utf8mb4_bin NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `fd_messages`
--

INSERT INTO `fd_messages` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id`, `username`, `content`) VALUES
(1, '2023-07-01 16:32:45.765', '2023-08-07 16:32:45.765', NULL, '104', '10009', 'user9', '这是咱们学校勤工俭学服务中心的QQ群，平时要是没有课了的话，有想勤工俭学的可以进群了解一下群号：839784935'),
(2, '2023-07-01 16:32:46.775', '2023-08-07 16:32:45.775', NULL, '104', '10010', 'user10', '同学们好！这是我们学校勤工俭学协会服务中心的QQ群，想做兼职的同学可以加入群聊了解一下，明天将发布新的兼职信息！群号：519843267。'),
(3, '2023-07-02 16:31:36.765', '2023-08-07 16:32:45.765', NULL, '104', '10009', 'user9', '这是咱们学校勤工俭学服务中心的QQ群，平时要是没有课了的话，有想勤工俭学的可以进群了解一下群号：839784935'),
(4, '2023-07-02 16:31:40.775', '2023-08-07 16:32:45.775', NULL, '104', '10010', 'user10', '同学们好！这是我们学校勤工俭学协会服务中心的QQ群，想做兼职的同学可以加入群聊了解一下，明天将发布新的兼职信息！群号：519843267。'),
(5, '2023-07-03 16:34:26.765', '2023-08-07 16:32:45.765', NULL, '104', '10009', 'user9', '这是咱们学校勤工俭学服务中心的QQ群，平时要是没有课了的话，有想勤工俭学的可以进群了解一下群号：839784935'),
(6, '2023-07-03 16:34:31.775', '2023-08-07 16:32:45.775', NULL, '104', '10010', 'user10', '同学们好！这是我们学校勤工俭学协会服务中心的QQ群，想做兼职的同学可以加入群聊了解一下，明天将发布新的兼职信息！群号：519843267。'),
(7, '2023-07-04 16:32:23.765', '2023-08-07 16:32:45.765', NULL, '104', '10009', 'user9', '这是咱们学校勤工俭学服务中心的QQ群，平时要是没有课了的话，有想勤工俭学的可以进群了解一下群号：839784935'),
(8, '2023-07-04 16:32:30.775', '2023-08-07 16:32:45.775', NULL, '104', '10010', 'user10', '同学们好！这是我们学校勤工俭学协会服务中心的QQ群，想做兼职的同学可以加入群聊了解一下，明天将发布新的兼职信息！群号：519843267。'),
(9, '2023-07-05 16:31:52.765', '2023-08-07 16:32:45.765', NULL, '104', '10009', 'user9', '这是咱们学校勤工俭学服务中心的QQ群，平时要是没有课了的话，有想勤工俭学的可以进群了解一下群号：839784935'),
(10, '2023-07-05 16:31:54.771', '2023-08-07 16:32:45.775', NULL, '104', '10010', 'user10', '同学们好！这是我们学校勤工俭学协会服务中心的QQ群，想做兼职的同学可以加入群聊了解一下，明天将发布新的兼职信息！群号：519843267。');

-- --------------------------------------------------------

--
-- 表的结构 `groups`
--

CREATE TABLE `groups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(191) COLLATE utf8mb4_bin NOT NULL,
  `group_name` longtext COLLATE utf8mb4_bin NOT NULL,
  `status` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `sum_msg_num` bigint(20) NOT NULL DEFAULT '0',
  `avg_msg_len` double NOT NULL DEFAULT '0',
  `member_num` bigint(20) NOT NULL DEFAULT '0',
  `max_member_num` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `groups`
--

INSERT INTO `groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `group_name`, `status`, `sum_msg_num`, `avg_msg_len`, `member_num`, `max_member_num`) VALUES
(1, '2023-08-07 16:32:13.721', '2023-08-07 16:37:52.070', NULL, '104', '群4', '2', 1020, 27.59701492537314, 534, 1000);

-- --------------------------------------------------------

--
-- 表的结构 `group_cache_messages`
--

CREATE TABLE `group_cache_messages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(2047) COLLATE utf8mb4_bin NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_cache_messages`
--

INSERT INTO `group_cache_messages` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id`, `username`, `content`) VALUES
(21, '2023-07-05 16:33:59.074', '2023-08-07 16:33:59.074', NULL, '104', '10015', 'user15', '还能逛街'),
(22, '2023-07-05 16:34:04.098', '2023-08-07 16:34:04.098', NULL, '104', '10011', 'user11', '万达好像离了几公里哈[CQ:face,id=...]'),
(23, '2023-07-05 16:34:06.104', '2023-08-07 16:34:09.104', NULL, '104', '10012', 'user12', '我咋去年没这个价'),
(24, '2023-07-05 16:34:17.161', '2023-08-07 16:34:14.124', NULL, '104', '10016', 'user16', '来我宿舍[CQ:face,id=...]'),
(25, '2023-07-05 16:34:21.150', '2023-08-07 16:34:19.150', NULL, '104', '10012', 'user12', '坐公交就行'),
(26, '2023-07-05 16:34:24.175', '2023-08-07 16:34:24.175', NULL, '104', '10015', 'user15', '万达打车过来十几块'),
(27, '2023-07-05 16:34:32.210', '2023-08-07 16:34:29.210', NULL, '104', '10012', 'user12', '打胶吗'),
(28, '2023-07-05 16:34:34.229', '2023-08-07 16:34:34.229', NULL, '104', '10016', 'user16', '开学季正常会贵一点'),
(29, '2023-07-05 16:34:37.250', '2023-08-07 16:34:39.250', NULL, '104', '10013', 'user13', '那个防学长的图呢'),
(30, '2023-07-05 16:34:44.282', '2023-08-07 16:34:44.282', NULL, '104', '10012', 'user12', '我空间里'),
(31, '2023-07-05 16:34:52.302', '2023-08-07 16:34:49.302', NULL, '104', '10012', 'user12', '（bushi）'),
(32, '2023-07-05 16:34:54.330', '2023-08-07 16:34:54.330', NULL, '104', '10013', 'user13', '[CQ:image,file=...]'),
(33, '2023-07-05 16:34:59.360', '2023-08-07 16:34:59.360', NULL, '104', '10012', 'user12', '[CQ:image,file=...]'),
(34, '2023-07-05 16:35:02.363', '2023-08-07 16:35:04.363', NULL, '104', '10012', 'user12', '[CQ:image,file=...]'),
(35, '2023-07-05 16:35:12.397', '2023-08-07 16:35:09.397', NULL, '104', '10012', 'user12', '是这样的'),
(36, '2023-07-05 16:35:14.416', '2023-08-07 16:35:14.416', NULL, '104', '10013', 'user13', '男孩子也要保护好自己啊'),
(37, '2023-07-05 16:35:19.448', '2023-08-07 16:35:19.448', NULL, '104', '10016', 'user16', '[CQ:image,file=...]'),
(38, '2023-07-05 16:35:20.480', '2023-08-07 16:35:24.480', NULL, '104', '10012', 'user12', '噗，男孩子还怕被噶了？'),
(39, '2023-07-05 16:35:28.487', '2023-08-07 16:35:29.487', NULL, '104', '10011', 'user11', '新生开学还要军训吗'),
(40, '2023-07-05 16:35:34.491', '2023-08-07 16:35:34.491', NULL, '104', '10017', 'user17', '她也是曾经这么跟我说的'),
(41, '2023-07-05 16:35:41.517', '2023-08-07 16:35:39.517', NULL, '104', '10015', 'user15', '话说你们能住进新宿舍吗，羡慕'),
(42, '2023-07-05 16:35:44.526', '2023-08-07 16:35:44.526', NULL, '104', '10016', 'user16', '不然呢？'),
(43, '2023-07-05 16:35:46.540', '2023-08-07 16:35:49.540', NULL, '104', '10013', 'user13', '？？？？？？'),
(44, '2023-07-05 16:35:54.546', '2023-08-07 16:35:54.546', NULL, '104', '10012', 'user12', '这个问题问得好'),
(45, '2023-07-05 16:35:59.568', '2023-08-07 16:35:59.568', NULL, '104', '10016', 'user16', '难不成让20，21再军训一次？'),
(46, '2023-07-05 16:36:08.595', '2023-08-07 16:36:04.595', NULL, '104', '10013', 'user13', '高三那个还是大学那个'),
(47, '2023-07-05 16:36:09.601', '2023-08-07 16:36:09.601', NULL, '104', '10012', 'user12', '就像问人离开了一氧化二氢能不能活'),
(48, '2023-07-05 16:36:17.633', '2023-08-07 16:36:14.633', NULL, '104', '10016', 'user16', '今年大一军训，明年大二军训是吧'),
(49, '2023-07-05 16:36:19.641', '2023-08-07 16:36:19.641', NULL, '104', '10011', 'user11', '本科生和研究生都要军训吗'),
(50, '2023-07-05 16:36:22.656', '2023-08-07 16:36:24.656', NULL, '104', '10013', 'user13', '军训 今年大1训 明年大2训 后年大3'),
(51, '2023-07-05 16:36:29.686', '2023-08-07 16:36:29.686', NULL, '104', '10015', 'user15', '研究生不用训'),
(52, '2023-07-05 16:36:38.704', '2023-08-07 16:36:34.704', NULL, '104', '10013', 'user13', '研究生不逊'),
(53, '2023-07-05 16:36:41.708', '2023-08-07 16:36:39.708', NULL, '104', '10011', 'user11', '[CQ:image,file=...]'),
(54, '2023-07-05 16:36:44.730', '2023-08-07 16:36:44.730', NULL, '104', '10012', 'user12', '饶了23届吧'),
(55, '2023-07-05 16:36:49.748', '2023-08-07 16:36:49.748', NULL, '104', '10015', 'user15', '都老胳膊老腿的'),
(56, '2023-07-05 16:36:51.756', '2023-08-07 16:36:54.760', NULL, '104', '10015', 'user15', '训坏了不好'),
(57, '2023-07-05 16:36:59.775', '2023-08-07 16:36:59.775', NULL, '104', '10012', 'user12', '本科生逊，一语双关是吧'),
(58, '2023-07-05 16:37:07.792', '2023-08-07 16:37:04.792', NULL, '104', '10013', 'user13', '训'),
(59, '2023-07-05 16:37:09.805', '2023-08-07 16:37:09.805', NULL, '104', '10011', 'user11', '咱们学校是不是进宿舍要刷脸'),
(60, '2023-07-05 16:37:17.837', '2023-08-07 16:37:14.819', NULL, '104', '10012', 'user12', '不是'),
(61, '2023-07-05 16:37:19.843', '2023-08-07 16:37:19.843', NULL, '104', '10013', 'user13', '是'),
(62, '2023-07-05 16:37:20.866', '2023-08-07 16:37:24.866', NULL, '104', '10012', 'user12', '玉兰卡'),
(63, '2023-07-05 16:37:25.873', '2023-08-07 16:37:29.873', NULL, '104', '10012', 'user12', '？'),
(64, '2023-07-05 16:37:38.905', '2023-08-07 16:37:34.905', NULL, '104', '10012', 'user12', '研究生宿舍？'),
(65, '2023-07-05 16:37:39.937', '2023-08-07 16:37:39.937', NULL, '104', '10013', 'user13', '宿舍楼？'),
(66, '2023-07-05 16:37:44.956', '2023-08-07 16:37:44.956', NULL, '104', '10011', 'user11', '[CQ:image,file=...]'),
(67, '2023-07-05 16:37:51.986', '2023-08-07 16:37:49.986', NULL, '104', '10012', 'user12', '那也是玉兰卡');

-- --------------------------------------------------------

--
-- 表的结构 `group_graphs`
--

CREATE TABLE `group_graphs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id1` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `username1` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id2` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `username2` longtext COLLATE utf8mb4_bin NOT NULL,
  `value1` double NOT NULL DEFAULT '0',
  `value2` double NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_graphs`
--

INSERT INTO `group_graphs` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id1`, `username1`, `user_id2`, `username2`, `value1`, `value2`) VALUES
(1, '2023-08-07 16:32:23.758', '2023-08-07 16:37:49.991', NULL, '104', '10012', 'user12', '10011', 'user11', 24.9, 12.3),
(2, '2023-08-07 16:32:33.802', '2023-08-07 16:36:19.651', NULL, '104', '10009', 'user9', '10011', 'user11', 1, 2.5),
(3, '2023-08-07 16:32:33.809', '2023-08-07 16:36:09.607', NULL, '104', '10009', 'user9', '10012', 'user12', 1.5, 7.5),
(4, '2023-08-07 16:32:38.813', '2023-08-07 16:36:19.653', NULL, '104', '10010', 'user10', '10011', 'user11', 1, 2.5),
(5, '2023-08-07 16:32:38.817', '2023-08-07 16:36:09.610', NULL, '104', '10010', 'user10', '10012', 'user12', 1.5, 7.5),
(6, '2023-08-07 16:32:38.824', '2023-08-07 16:32:38.824', NULL, '104', '10010', 'user10', '10009', 'user9', 0.5, 0),
(7, '2023-08-07 16:32:53.876', '2023-08-07 16:37:44.983', NULL, '104', '10013', 'user13', '10011', 'user11', 13.1, 3.6),
(8, '2023-08-07 16:32:53.882', '2023-08-07 16:37:50.004', NULL, '104', '10013', 'user13', '10012', 'user12', 20.7, 7),
(9, '2023-08-07 16:32:53.890', '2023-08-07 16:36:24.668', NULL, '104', '10013', 'user13', '10009', 'user9', 4.5, 0),
(10, '2023-08-07 16:32:53.897', '2023-08-07 16:36:24.671', NULL, '104', '10013', 'user13', '10010', 'user10', 4.5, 0),
(11, '2023-08-07 16:33:13.923', '2023-08-07 16:37:09.819', NULL, '104', '10014', 'user14', '10011', 'user11', 4.8, 4.4),
(12, '2023-08-07 16:33:13.928', '2023-08-07 16:37:14.840', NULL, '104', '10014', 'user14', '10012', 'user12', 8.4, 6.6),
(13, '2023-08-07 16:33:13.938', '2023-08-07 16:33:34.012', NULL, '104', '10014', 'user14', '10009', 'user9', 2, 0),
(14, '2023-08-07 16:33:13.944', '2023-08-07 16:33:34.016', NULL, '104', '10014', 'user14', '10010', 'user10', 2, 0),
(15, '2023-08-07 16:33:13.963', '2023-08-07 16:37:19.858', NULL, '104', '10014', 'user14', '10013', 'user13', 0.8, 6.6),
(16, '2023-08-07 16:33:44.026', '2023-08-07 16:37:44.975', NULL, '104', '10015', 'user15', '10011', 'user11', 8.2, 2),
(17, '2023-08-07 16:33:44.031', '2023-08-07 16:37:49.997', NULL, '104', '10015', 'user15', '10012', 'user12', 13, 6),
(18, '2023-08-07 16:33:44.041', '2023-08-07 16:36:29.695', NULL, '104', '10015', 'user15', '10009', 'user9', 3, 0),
(19, '2023-08-07 16:33:44.045', '2023-08-07 16:36:29.699', NULL, '104', '10015', 'user15', '10010', 'user10', 3, 0),
(20, '2023-08-07 16:33:44.053', '2023-08-07 16:37:39.954', NULL, '104', '10015', 'user15', '10013', 'user13', 3.6, 2.4),
(21, '2023-08-07 16:34:14.125', '2023-08-07 16:37:44.980', NULL, '104', '10016', 'user16', '10011', 'user11', 7.2, 0.7),
(22, '2023-08-07 16:34:14.131', '2023-08-07 16:37:50.001', NULL, '104', '10016', 'user16', '10012', 'user12', 12.6, 2.8),
(23, '2023-08-07 16:34:14.138', '2023-08-07 16:36:14.647', NULL, '104', '10016', 'user16', '10009', 'user9', 3, 0),
(24, '2023-08-07 16:34:14.141', '2023-08-07 16:36:14.651', NULL, '104', '10016', 'user16', '10010', 'user10', 3, 0),
(25, '2023-08-07 16:34:14.151', '2023-08-07 16:37:39.960', NULL, '104', '10016', 'user16', '10013', 'user13', 1.2, 0.7),
(26, '2023-08-07 16:35:34.492', '2023-08-07 16:35:34.517', NULL, '104', '10017', 'user17', '10011', 'user11', 1.2, 0),
(27, '2023-08-07 16:35:34.497', '2023-08-07 16:35:34.515', NULL, '104', '10017', 'user17', '10012', 'user12', 2.1, 0),
(28, '2023-08-07 16:35:34.502', '2023-08-07 16:35:34.502', NULL, '104', '10017', 'user17', '10009', 'user9', 0.5, 0),
(29, '2023-08-07 16:35:34.506', '2023-08-07 16:35:34.506', NULL, '104', '10017', 'user17', '10010', 'user10', 0.5, 0),
(30, '2023-08-07 16:35:34.513', '2023-08-07 16:35:34.513', NULL, '104', '10017', 'user17', '10013', 'user13', 0.2, 0),
(31, '2023-08-07 16:36:49.755', '2023-08-07 16:36:54.778', NULL, '104', '10015', 'user15', '10014', 'user14', 4.4, 0);

-- --------------------------------------------------------

--
-- 表的结构 `group_keywords`
--

CREATE TABLE `group_keywords` (
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `keyword` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `freq` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_keywords`
--

INSERT INTO `group_keywords` (`group_id`, `keyword`, `freq`) VALUES
('104', '兼职', 10),
('104', '勤工俭学', 15),
('104', '群号', 10);

-- --------------------------------------------------------

--
-- 表的结构 `group_messages`
--

CREATE TABLE `group_messages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(2047) COLLATE utf8mb4_bin NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `group_times`
--

CREATE TABLE `group_times` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `month` bigint(20) NOT NULL DEFAULT '0',
  `day` bigint(20) NOT NULL DEFAULT '0',
  `hour1` bigint(20) NOT NULL DEFAULT '0',
  `hour2` bigint(20) NOT NULL DEFAULT '0',
  `hour3` bigint(20) NOT NULL DEFAULT '0',
  `hour4` bigint(20) NOT NULL DEFAULT '0',
  `hour5` bigint(20) NOT NULL DEFAULT '0',
  `hour6` bigint(20) NOT NULL DEFAULT '0',
  `hour7` bigint(20) NOT NULL DEFAULT '0',
  `hour8` bigint(20) NOT NULL DEFAULT '0',
  `hour9` bigint(20) NOT NULL DEFAULT '0',
  `hour10` bigint(20) NOT NULL DEFAULT '0',
  `hour11` bigint(20) NOT NULL DEFAULT '0',
  `hour12` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_times`
--

INSERT INTO `group_times` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `month`, `day`, `hour1`, `hour2`, `hour3`, `hour4`, `hour5`, `hour6`, `hour7`, `hour8`, `hour9`, `hour10`, `hour11`, `hour12`) VALUES
(1, '2023-07-01 16:32:18.746', '2023-08-07 16:37:49.985', NULL, '104', 7, 1, 1, 0, 0, 0, 6, 50, 46, 96, 56, 54, 13, 10),
(2, '2023-07-02 00:00:00.000', NULL, NULL, '104', 7, 2, 0, 4, 0, 0, 15, 13, 26, 12, 16, 12, 13, 20),
(3, '2023-07-03 00:00:00.000', NULL, NULL, '104', 7, 3, 8, 0, 2, 0, 51, 29, 121, 4, 9, 7, 4, 6),
(4, '2023-07-04 00:00:00.000', NULL, NULL, '104', 7, 4, 0, 0, 0, 0, 16, 12, 12, 17, 30, 44, 22, 5),
(5, '2023-07-05 00:00:00.000', NULL, NULL, '104', 7, 5, 0, 1, 1, 5, 3, 12, 37, 21, 65, 0, 3, 0);

-- --------------------------------------------------------

--
-- 表的结构 `group_users`
--

CREATE TABLE `group_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `freq` bigint(20) NOT NULL DEFAULT '0',
  `value` double NOT NULL DEFAULT '0',
  `keyword_num` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_users`
--

INSERT INTO `group_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id`, `username`, `freq`, `value`, `keyword_num`) VALUES
(1, '2023-08-07 16:32:18.746', '2023-08-07 16:37:44.986', NULL, '104', '10011', 'user11', 6, 28, 0),
(2, '2023-08-07 16:32:23.754', '2023-08-07 16:37:50.005', NULL, '104', '10012', 'user12', 9, 62.3, 0),
(3, '2023-08-07 16:32:33.797', '2023-08-07 16:32:33.814', NULL, '104', '10009', 'user9', 1, 2.5, 15),
(4, '2023-08-07 16:32:38.807', '2023-08-07 16:32:38.825', NULL, '104', '10010', 'user10', 1, 3, 20),
(5, '2023-08-07 16:32:53.875', '2023-08-07 16:37:39.965', NULL, '104', '10013', 'user13', 7, 52.5, 0),
(6, '2023-08-07 16:33:13.923', '2023-08-07 16:33:34.030', NULL, '104', '10014', 'user14', 2, 18, 0),
(7, '2023-08-07 16:33:44.025', '2023-08-07 16:36:54.785', NULL, '104', '10015', 'user15', 4, 35.2, 0),
(8, '2023-08-07 16:34:14.124', '2023-08-07 16:36:14.668', NULL, '104', '10016', 'user16', 2, 27, 0),
(9, '2023-08-07 16:35:34.491', '2023-08-07 16:35:34.519', NULL, '104', '10017', 'user17', 1, 4.5, 0);

-- --------------------------------------------------------

--
-- 表的结构 `group_user_times`
--

CREATE TABLE `group_user_times` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `month` bigint(20) NOT NULL DEFAULT '0',
  `day` bigint(20) NOT NULL DEFAULT '0',
  `hour1` bigint(20) NOT NULL DEFAULT '0',
  `hour2` bigint(20) NOT NULL DEFAULT '0',
  `hour3` bigint(20) NOT NULL DEFAULT '0',
  `hour4` bigint(20) NOT NULL DEFAULT '0',
  `hour5` bigint(20) NOT NULL DEFAULT '0',
  `hour6` bigint(20) NOT NULL DEFAULT '0',
  `hour7` bigint(20) NOT NULL DEFAULT '0',
  `hour8` bigint(20) NOT NULL DEFAULT '0',
  `hour9` bigint(20) NOT NULL DEFAULT '0',
  `hour10` bigint(20) NOT NULL DEFAULT '0',
  `hour11` bigint(20) NOT NULL DEFAULT '0',
  `hour12` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `group_user_times`
--

INSERT INTO `group_user_times` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `user_id`, `month`, `day`, `hour1`, `hour2`, `hour3`, `hour4`, `hour5`, `hour6`, `hour7`, `hour8`, `hour9`, `hour10`, `hour11`, `hour12`) VALUES
(1, '2023-07-01 16:32:18.746', '2023-08-07 16:37:44.958', NULL, '104', '10011', 7, 1, 1, 6, 4, 1, 6, 6, 1, 4, 6, 2, 4, 1),
(2, '2023-07-01 16:32:23.750', '2023-08-07 16:37:49.985', NULL, '104', '10012', 7, 1, 5, 3, 8, 4, 1, 5, 8, 5, 3, 3, 1, 7),
(3, '2023-07-01 16:32:33.798', '2023-08-07 16:32:33.801', NULL, '104', '10009', 7, 1, 0, 0, 0, 0, 0, 0, 0, 0, 40, 0, 0, 0),
(4, '2023-07-01 16:32:38.808', '2023-08-07 16:32:38.813', NULL, '104', '10010', 7, 1, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 0),
(5, '2023-07-01 16:32:53.876', '2023-08-07 16:37:39.938', NULL, '104', '10013', 7, 1, 3, 6, 1, 1, 3, 6, 1, 4, 8, 5, 5, 0),
(6, '2023-07-01 16:33:13.922', '2023-08-07 16:33:33.995', NULL, '104', '10014', 7, 1, 4, 2, 1, 3, 3, 2, 2, 5, 8, 4, 1, 1),
(7, '2023-07-01 16:33:44.025', '2023-08-07 16:36:54.761', NULL, '104', '10015', 7, 1, 6, 4, 6, 7, 8, 4, 6, 3, 6, 1, 1, 4),
(8, '2023-07-01 16:34:14.124', '2023-08-07 16:36:14.635', NULL, '104', '10016', 7, 1, 4, 2, 1, 6, 3, 8, 2, 5, 4, 7, 6, 2),
(9, '2023-07-01 16:35:34.491', '2023-08-07 16:35:34.494', NULL, '104', '10017', 7, 1, 8, 3, 3, 3, 6, 6, 6, 2, 1, 6, 2, 2),
(10, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10011', 7, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
(11, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10012', 7, 2, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 1),
(12, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10013', 7, 2, 1, 2, 1, 1, 1, 2, 2, 1, 2, 2, 1, 2),
(13, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10014', 7, 2, 1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1),
(14, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10015', 7, 2, 2, 2, 1, 2, 2, 2, 1, 1, 2, 2, 2, 1),
(15, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10016', 7, 2, 2, 1, 2, 2, 2, 1, 1, 2, 1, 1, 1, 1),
(16, '2023-07-02 00:00:00.000', NULL, NULL, '104', '10017', 7, 2, 1, 2, 1, 4, 2, 2, 1, 2, 1, 4, 1, 2),
(17, NULL, NULL, NULL, '104', '10011', 7, 3, 1, 2, 2, 4, 2, 4, 4, 4, 3, 4, 3, 2),
(18, NULL, NULL, NULL, '104', '10012', 7, 3, 4, 4, 1, 3, 1, 2, 2, 3, 1, 3, 1, 3),
(19, NULL, NULL, NULL, '104', '10013', 7, 3, 4, 3, 3, 3, 1, 3, 2, 2, 6, 4, 1, 3),
(20, NULL, NULL, NULL, '104', '10014', 7, 3, 3, 1, 1, 4, 5, 2, 1, 3, 6, 1, 1, 6),
(21, NULL, NULL, NULL, '104', '10015', 7, 3, 5, 4, 3, 3, 4, 1, 1, 6, 2, 2, 4, 2),
(22, NULL, NULL, NULL, '104', '10016', 7, 3, 4, 4, 4, 3, 1, 4, 2, 3, 2, 2, 6, 3),
(23, NULL, NULL, NULL, '104', '10017', 7, 3, 1, 1, 2, 6, 1, 6, 4, 3, 3, 3, 2, 1),
(24, NULL, NULL, NULL, '104', '10011', 7, 4, 1, 4, 2, 1, 4, 1, 2, 1, 1, 1, 4, 2),
(25, NULL, NULL, NULL, '104', '10012', 7, 4, 2, 2, 2, 4, 2, 2, 1, 1, 2, 2, 2, 4),
(26, NULL, NULL, NULL, '104', '10013', 7, 4, 1, 2, 1, 2, 2, 4, 2, 2, 1, 1, 3, 2),
(27, NULL, NULL, NULL, '104', '10014', 7, 4, 2, 2, 1, 3, 2, 2, 2, 2, 2, 2, 2, 2),
(28, NULL, NULL, NULL, '104', '10015', 7, 4, 1, 4, 2, 2, 1, 1, 1, 0, 2, 2, 2, 2),
(29, NULL, NULL, NULL, '104', '10016', 7, 4, 1, 3, 1, 4, 1, 1, 1, 2, 2, 1, 4, 2),
(30, NULL, NULL, NULL, '104', '10017', 7, 4, 2, 3, 2, 1, 1, 2, 3, 1, 1, 4, 4, 0),
(31, NULL, NULL, NULL, '104', '10011', 7, 5, 1, 2, 2, 3, 4, 5, 6, 1, 5, 1, 9, 1),
(32, NULL, NULL, NULL, '104', '10012', 7, 5, 8, 7, 1, 7, 6, 8, 3, 1, 1, 2, 1, 0),
(35, NULL, NULL, NULL, '104', '10013', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(36, NULL, NULL, NULL, '104', '10014', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(37, NULL, NULL, NULL, '104', '10015', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(38, NULL, NULL, NULL, '104', '10016', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(39, NULL, NULL, NULL, '104', '10017', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(40, NULL, NULL, NULL, '104', '10009', 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 40, 0, 0, 0),
(41, NULL, NULL, NULL, '104', '10010', 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 0),
(42, NULL, NULL, NULL, '104', '10009', 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 40, 0, 0, 0),
(43, NULL, NULL, NULL, '104', '10010', 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 0),
(44, NULL, NULL, NULL, '104', '10009', 7, 4, 0, 0, 0, 0, 0, 0, 0, 0, 40, 0, 0, 0),
(45, NULL, NULL, NULL, '104', '10010', 7, 4, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 0),
(46, NULL, NULL, NULL, '104', '10009', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 40, 0, 0, 0),
(47, NULL, NULL, NULL, '104', '10010', 7, 5, 0, 0, 0, 0, 0, 0, 0, 0, 60, 0, 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `keyword_lists`
--

CREATE TABLE `keyword_lists` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `keyword` varchar(191) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `keyword_lists`
--

INSERT INTO `keyword_lists` (`id`, `created_at`, `updated_at`, `deleted_at`, `keyword`) VALUES
(1, '2023-08-07 17:32:45.025', '2023-08-07 17:32:45.025', NULL, '勤工俭学'),
(2, '2023-08-07 17:32:52.534', '2023-08-07 17:32:52.534', NULL, '兼职'),
(3, '2023-08-07 17:34:33.568', '2023-08-07 17:34:33.568', NULL, '群号');

-- --------------------------------------------------------

--
-- 表的结构 `key_users`
--

CREATE TABLE `key_users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(191) COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin,
  `remark` longtext COLLATE utf8mb4_bin,
  `sum_msg_num` bigint(20) NOT NULL DEFAULT '0',
  `avg_msg_len` double NOT NULL DEFAULT '0',
  `status` varchar(191) COLLATE utf8mb4_bin NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `key_users`
--

INSERT INTO `key_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `username`, `remark`, `sum_msg_num`, `avg_msg_len`, `status`) VALUES
(1, '2023-08-07 16:32:45.773', '2023-08-07 17:41:51.976', NULL, '10009', 'user9', '', 1, 149, '0'),
(2, '2023-08-07 16:32:45.778', '2023-08-07 17:42:07.153', NULL, '10010', 'user10', '', 1, 185, '0');

-- --------------------------------------------------------

--
-- 表的结构 `key_user_cache_messages`
--

CREATE TABLE `key_user_cache_messages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `username` longtext COLLATE utf8mb4_bin NOT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `group_name` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(2047) COLLATE utf8mb4_bin NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `key_user_groups`
--

CREATE TABLE `key_user_groups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `group_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `group_name` longtext COLLATE utf8mb4_bin NOT NULL,
  `freq` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `key_user_keywords`
--

CREATE TABLE `key_user_keywords` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `keyword` longtext COLLATE utf8mb4_bin NOT NULL,
  `freq` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `key_user_times`
--

CREATE TABLE `key_user_times` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext COLLATE utf8mb4_bin NOT NULL,
  `month` bigint(20) NOT NULL DEFAULT '0',
  `day` bigint(20) NOT NULL DEFAULT '0',
  `hour1` bigint(20) NOT NULL DEFAULT '0',
  `hour2` bigint(20) NOT NULL DEFAULT '0',
  `hour3` bigint(20) NOT NULL DEFAULT '0',
  `hour4` bigint(20) NOT NULL DEFAULT '0',
  `hour5` bigint(20) NOT NULL DEFAULT '0',
  `hour6` bigint(20) NOT NULL DEFAULT '0',
  `hour7` bigint(20) NOT NULL DEFAULT '0',
  `hour8` bigint(20) NOT NULL DEFAULT '0',
  `hour9` bigint(20) NOT NULL DEFAULT '0',
  `hour10` bigint(20) NOT NULL DEFAULT '0',
  `hour11` bigint(20) NOT NULL DEFAULT '0',
  `hour12` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `ts_cache_results`
--

CREATE TABLE `ts_cache_results` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `amplitude1` double NOT NULL DEFAULT '0',
  `amplitude2` double NOT NULL DEFAULT '0',
  `amplitude3` double NOT NULL DEFAULT '0',
  `freq1` double NOT NULL DEFAULT '0',
  `freq2` double NOT NULL DEFAULT '0',
  `freq3` double NOT NULL DEFAULT '0',
  `phase1` double NOT NULL DEFAULT '0',
  `phase2` double NOT NULL DEFAULT '0',
  `phase3` double NOT NULL DEFAULT '0',
  `acf1` double NOT NULL DEFAULT '0',
  `acf2` double NOT NULL DEFAULT '0',
  `acf3` double NOT NULL DEFAULT '0',
  `mse_mean` double NOT NULL DEFAULT '0',
  `mse_variance` double NOT NULL DEFAULT '0',
  `outliers` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `outliers_mse` double NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `ts_cache_times`
--

CREATE TABLE `ts_cache_times` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `month` bigint(20) NOT NULL DEFAULT '0',
  `day` bigint(20) NOT NULL DEFAULT '0',
  `hour1` bigint(20) NOT NULL DEFAULT '0',
  `hour2` bigint(20) NOT NULL DEFAULT '0',
  `hour3` bigint(20) NOT NULL DEFAULT '0',
  `hour4` bigint(20) NOT NULL DEFAULT '0',
  `hour5` bigint(20) NOT NULL DEFAULT '0',
  `hour6` bigint(20) NOT NULL DEFAULT '0',
  `hour7` bigint(20) NOT NULL DEFAULT '0',
  `hour8` bigint(20) NOT NULL DEFAULT '0',
  `hour9` bigint(20) NOT NULL DEFAULT '0',
  `hour10` bigint(20) NOT NULL DEFAULT '0',
  `hour11` bigint(20) NOT NULL DEFAULT '0',
  `hour12` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `ts_cache_times`
--

INSERT INTO `ts_cache_times` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_id`, `month`, `day`, `hour1`, `hour2`, `hour3`, `hour4`, `hour5`, `hour6`, `hour7`, `hour8`, `hour9`, `hour10`, `hour11`, `hour12`) VALUES
(1, '2023-07-01 00:00:00.000', NULL, NULL, '104', 7, 1, 1, 0, 0, 0, 6, 50, 46, 96, 56, 54, 13, 10),
(2, '2023-07-02 00:00:00.000', NULL, NULL, '104', 7, 2, 0, 4, 0, 0, 15, 13, 26, 12, 16, 12, 13, 20),
(3, '2023-07-03 00:00:00.000', NULL, NULL, '104', 7, 3, 3, 8, 0, 2, 0, 15, 29, 121, 4, 9, 4, 6),
(4, '2023-07-04 00:00:00.000', NULL, NULL, '104', 7, 4, 0, 0, 0, 0, 16, 12, 12, 17, 30, 44, 22, 5),
(5, '2023-07-05 00:00:00.000', NULL, NULL, '104', 7, 5, 0, 1, 1, 5, 3, 12, 37, 21, 2, 0, 3, 0);

--
-- 转储表的索引
--

--
-- 表的索引 `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD KEY `idx_admins_deleted_at` (`deleted_at`);

--
-- 表的索引 `fd_cache_messages`
--
ALTER TABLE `fd_cache_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_fd_cache_messages_deleted_at` (`deleted_at`);

--
-- 表的索引 `fd_groups`
--
ALTER TABLE `fd_groups`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `GroupID_SenderID` (`group_id`,`sender_id`),
  ADD KEY `idx_fd_groups_deleted_at` (`deleted_at`);

--
-- 表的索引 `fd_messages`
--
ALTER TABLE `fd_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_fd_messages_deleted_at` (`deleted_at`);

--
-- 表的索引 `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `group_id` (`group_id`),
  ADD KEY `idx_groups_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_cache_messages`
--
ALTER TABLE `group_cache_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_group_cache_messages_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_graphs`
--
ALTER TABLE `group_graphs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UserID1_UserID2` (`user_id1`,`user_id2`),
  ADD KEY `idx_group_graphs_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_keywords`
--
ALTER TABLE `group_keywords`
  ADD UNIQUE KEY `GroupID_Keyword` (`group_id`,`keyword`);

--
-- 表的索引 `group_messages`
--
ALTER TABLE `group_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_group_messages_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_times`
--
ALTER TABLE `group_times`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `GroupID_Month_Day` (`group_id`,`month`,`day`),
  ADD KEY `idx_group_times_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_users`
--
ALTER TABLE `group_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `GroupID_UserID` (`group_id`,`user_id`),
  ADD KEY `idx_group_users_deleted_at` (`deleted_at`);

--
-- 表的索引 `group_user_times`
--
ALTER TABLE `group_user_times`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `GroupID_UserID_Month_Day` (`group_id`,`user_id`,`month`,`day`),
  ADD KEY `idx_group_user_times_deleted_at` (`deleted_at`);

--
-- 表的索引 `keyword_lists`
--
ALTER TABLE `keyword_lists`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD KEY `idx_keyword_lists_deleted_at` (`deleted_at`);

--
-- 表的索引 `key_users`
--
ALTER TABLE `key_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD KEY `idx_key_users_deleted_at` (`deleted_at`);

--
-- 表的索引 `key_user_cache_messages`
--
ALTER TABLE `key_user_cache_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_key_user_cache_messages_deleted_at` (`deleted_at`);

--
-- 表的索引 `key_user_groups`
--
ALTER TABLE `key_user_groups`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_key_user_groups_deleted_at` (`deleted_at`);

--
-- 表的索引 `key_user_keywords`
--
ALTER TABLE `key_user_keywords`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_key_user_keywords_deleted_at` (`deleted_at`);

--
-- 表的索引 `key_user_times`
--
ALTER TABLE `key_user_times`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_key_user_times_deleted_at` (`deleted_at`);

--
-- 表的索引 `ts_cache_results`
--
ALTER TABLE `ts_cache_results`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_ts_cache_results_deleted_at` (`deleted_at`);

--
-- 表的索引 `ts_cache_times`
--
ALTER TABLE `ts_cache_times`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `GroupID_Month_Day` (`group_id`,`month`,`day`),
  ADD KEY `idx_ts_cache_times_deleted_at` (`deleted_at`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `fd_cache_messages`
--
ALTER TABLE `fd_cache_messages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=68;

--
-- 使用表AUTO_INCREMENT `fd_groups`
--
ALTER TABLE `fd_groups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `fd_messages`
--
ALTER TABLE `fd_messages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- 使用表AUTO_INCREMENT `groups`
--
ALTER TABLE `groups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `group_cache_messages`
--
ALTER TABLE `group_cache_messages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=68;

--
-- 使用表AUTO_INCREMENT `group_graphs`
--
ALTER TABLE `group_graphs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- 使用表AUTO_INCREMENT `group_messages`
--
ALTER TABLE `group_messages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `group_times`
--
ALTER TABLE `group_times`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `group_users`
--
ALTER TABLE `group_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- 使用表AUTO_INCREMENT `group_user_times`
--
ALTER TABLE `group_user_times`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=48;

--
-- 使用表AUTO_INCREMENT `keyword_lists`
--
ALTER TABLE `keyword_lists`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `key_users`
--
ALTER TABLE `key_users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `key_user_cache_messages`
--
ALTER TABLE `key_user_cache_messages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `key_user_groups`
--
ALTER TABLE `key_user_groups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `key_user_keywords`
--
ALTER TABLE `key_user_keywords`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `key_user_times`
--
ALTER TABLE `key_user_times`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `ts_cache_results`
--
ALTER TABLE `ts_cache_results`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `ts_cache_times`
--
ALTER TABLE `ts_cache_times`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
