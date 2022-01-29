<?php
error_reporting(0);
header("content-type:application/json");
$content=file_get_contents('php://input');
$data=json_decode($content,true);
include_once "php/config.php";
include_once "controller.php";
include_once "info.php";
include_once "record.php";
include_once "warning.php";