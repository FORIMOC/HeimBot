<?php
function filter($word)
{
    // CQ filter
    $word=preg_replace('/\[/','<span class="CQ">[',$word);
    $word=preg_replace('/]/',']</span>',$word);

    // base64 filter
    if (base64_encode(base64_decode($word)) === $word){
        $word="<span class='b64'>".$word."</span> (".base64_decode($word).")";
    }

    // 脏话 filter
    $word=preg_replace('/傻逼|妈的|操你妈|弱智|脑残|智障|狗东西|狗逼|你妈死了|你妈逼/i','<span class="bad">$0</span>',$word);

    // 键政 filter
    $word=preg_replace('/中共|土共|社工|入关|国党|兔友|神友|六四|纳粹|视奸/i','<span class="jz">$0</span>',$word);

    // 色情 filter
    $word=preg_replace('/奶子|阴道|阴蒂|阴茎|贞操锁|乳房/i','<span class="sex">$0</span>',$word);

    // 学术 filter
    $word=preg_replace('/python|javascript|C\+\+|java|ruby|go|php|rust|django|css|html|flask|nodejs|angular|vue|react|cms|wordpress/i','<span class="academic">$0</span>',$word);
    $word=preg_replace('/ctf|acm|数模|美赛|前端|后端|力扣|面试|cpc|leetcode|ieee/i','<span class="academic">$0</span>',$word);
    $word=preg_replace('/sql|注入|ssti|ssrf|xss|rce|撞库/i','<span class="academic">$0</span>',$word);
    $word=preg_replace('/保研|考研|实验室|论文|大作业|绩点/i','<span class="academic">$0</span>',$word);

    // 恶俗 filter
    $word=preg_replace('/杨程瑜|唐嘉雄|夏赫绅|骆泽文|徐志|康哲鸣|方陶松|白睿宁|温腾飞|江润汉/i','<span class="esu">$0</span>',$word);

    return $word;
}