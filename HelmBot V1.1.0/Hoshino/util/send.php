<?php
function send($sender, $target_list, $result){
    switch ($sender){
        case 'Hoshino':
            $num='1';
            break;
        case 'Komari':
            $num='2';
            break;
        case 'Kalina':
        default:
            $num='0';
            break;
    }
    foreach ($target_list as $key => $value) {
        $url='http://127.0.0.1:570'.$num.'/send_group_msg?group_id='.$value.'&message='.$result;
        fopen($url,'r');
    }
}