<?php
error_reporting(0);
header("content-type:application/json");
$content=file_get_contents('php://input');
$data=json_decode($content,true);

include_once 'util/send.php';
include_once 'util/filter.php';
include_once 'util/CQurlGetter.php';
include_once 'util/OCR.php';
include_once 'database/config.php';
include_once 'controller/controller.php';
include_once 'src/record.php';
include_once 'src/warning.php';
include_once 'src/history.php';