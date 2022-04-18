<?php
function OCR($lang, $img_src){
    if (!isset($img_src)){
        exit('{"reply": "未设置图片！"}');
    }
    switch ($lang){
        case '中文':
            $lang='chi_sim';
            break;
        default:
        case '英文':
            $lang='eng';
            break;
    }
    $script_src='./util/OCR.py';
    $cmd="python3 ".$script_src." ".$img_src." ".$lang;
    exec($cmd,$ret);
    return preg_replace('/ /','',implode('\n', $ret));
}

function OCR_array($lang, $img_src_array){
    if (!isset($img_src_array)){
        exit('{"reply": "未设置图片！"}');
    }
    switch ($lang){
        case '中文':
            $lang='chi_sim';
            break;
        default:
        case '英文':
            $lang='eng';
            break;
    }
    $result = [];
    $script_src='./util/OCR.py';
    foreach ($img_src_array as $key => $img_src){
        exec("python3 ".$script_src." ".$img_src." ".$lang,$ret[$key]);
        $result[$key]=preg_replace('/ /','',implode('\n',$ret[$key]));
    }
    return $result;
}