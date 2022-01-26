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
    $word=preg_replace('/xxx/i','<span class="bad">$0</span>',$word);

    // 键政 filter
    $word=preg_replace('/xxx/i','<span class="jz">$0</span>',$word);

    // 色情 filter
    $word=preg_replace('/xxx/i','<span class="sex">$0</span>',$word);

    // 学术 filter
    $word=preg_replace('/xxx/i','<span class="academic">$0</span>',$word);

    // 恶俗 filter
    $word=preg_replace('/xxx/i','<span class="esu">$0</span>',$word);

    return $word;
}