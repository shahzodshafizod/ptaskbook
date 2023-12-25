<?php
$num = mt_rand(10, 99);
$a = intval($num / 10);
$b = $num % 10;
$mun = $b * 10 + $a;
echo "Адади дурақама: ".$num;
echo "<br>Қисми даҳӣ: ".$a;
echo "<br>Қисми воҳидӣ: ".$b;
echo "<br>Адад бо чаппагӣ: ".$mun;
?>