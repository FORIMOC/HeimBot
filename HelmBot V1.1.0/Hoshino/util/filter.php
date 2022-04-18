<?php
function filter($str){
    if (preg_match('/drop|or|insert|update|select|delete|\|/i',$str)){
        return true;
    }else{
        return false;
    }
}