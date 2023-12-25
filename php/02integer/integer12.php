<?php
$num = mt_rand(100, 999);
$a = intval($num / 100);
$b = intval($num / 10) % 10;
$c = $num % 10;
$mun = $c * 100 + $b * 10 + $a;
echo "Адади серақама: ".$num;
echo "<br>Қисми садӣ: ".$a;
echo "<br>Қисми даҳӣ: ".$b;
echo "<br>Қисми воҳидӣ: ".$c;
echo "<br>Адад бо чаппагӣ: ".$mun;
?>