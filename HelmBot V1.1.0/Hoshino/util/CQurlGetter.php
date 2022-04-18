<?php
function getCQurl($CQcode){
    preg_match_all('/\[CQ:image,file=.*?,url=(.*?)]/',$CQcode,$match);
    return $match[1][0];
}

function getCQurl_array($CQcode){
    preg_match_all('/\[CQ:image,file=.*?,url=(.*?)]/',$CQcode,$match);
    return $match[1];
}