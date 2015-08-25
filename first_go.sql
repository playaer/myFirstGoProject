-- phpMyAdmin SQL Dump
-- version 4.0.10deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Aug 25, 2015 at 10:25 AM
-- Server version: 5.5.44-0ubuntu0.14.04.1
-- PHP Version: 5.5.9-1ubuntu4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `first_go`
--
CREATE DATABASE IF NOT EXISTS `first_go` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `first_go`;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `full_name`, `address`, `phone`, `email`, `password`, `hash`, `is_active`, `token`) VALUES
(1, 'Mikhail Yarotski', 'Lesnoj, Aleksandrova Str., 8-280', '<script>Alert()</script>', 'playaer80@gmail.com', '698d51a19d8a121ce581499d7b701668', '', 1, '71f7bc8bfcc339b23f5efc800bf1fd2c'),
(2, 'Mikhail Yarotski111', 'addr1', '2342532563', 'playaerpg@gmail.com', '698d51a19d8a121ce581499d7b701668', '', 1, '6cebad2b7d508d452496f8edec35abe8');

--
-- Dumping data for table `user_update`
--

INSERT INTO `user_update` (`id`, `user_id`, `updated_at`, `old_data`, `new_data`) VALUES
(1, 1, '2015-08-24 11:51:36', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}'),
(2, 1, '2015-08-24 11:53:18', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}'),
(3, 1, '2015-08-24 11:54:58', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}'),
(4, 1, '2015-08-24 11:57:16', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski111","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}'),
(5, 1, '2015-08-24 11:58:15', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}'),
(6, 1, '2015-08-24 12:00:38', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotskiwww","Phone":"+375292278461"}'),
(7, 1, '2015-08-24 12:02:30', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotskiwww","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}'),
(8, 1, '2015-08-24 12:20:11', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"+375292278461"}', '{"Address":"Lesnoj, Aleksandrova Str., 8-280","FullName":"Mikhail Yarotski","Phone":"\\u003cscript\\u003eAlert()\\u003c/script\\u003e"}'),
(9, 2, '2015-08-24 13:04:46', '{"Address":"addr","FullName":"Mikhail Yarotski111","Phone":"2342532563"}', '{"Address":"addr1","FullName":"Mikhail Yarotski111","Phone":"2342532563"}');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
