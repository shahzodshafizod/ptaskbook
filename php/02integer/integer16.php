<?php
$abc = mt_rand(100, 999);
$a = intval($abc / 100);
$b = intval($abc / 10) % 10;
$c = $abc % 10;
$acb = $a * 100 + $c * 10 + $b;
echo "Адади серақама: ".$abc;
echo "<br>Қисми садӣ: ".$a;
echo "<br>Қисми даҳӣ: ".$b;
echo "<br>Қисми воҳидӣ: ".$c;
echo "<br>Адад пас аз иҷрои амалиёт: ".$acb;
?>