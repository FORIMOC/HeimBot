<?php
include_once 'util/util.php';
$server_url = 'http://127.0.0.1:8740';

for ($i = 0; $i < 3; $i++){
    $json = json_decode(file_get_contents('data/data_group'.($i+1).'.json'), true);
    for ($j = 0; $j < sizeof($json['data']); $j++){
        http_post_json($server_url, json_encode($json['data'][$j]));
        sleep(5);
    }
}

echo "Test data generation completed\n";