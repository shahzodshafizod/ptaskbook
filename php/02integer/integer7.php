<?php
$N = mt_rand(10, 99);
$a = intval($N / 10);
$b = $N % 10;
$sum = $a + $b;
$zarb = $a * $b;
echo "Адади дурақама: ".$N;
echo "<br>Қисми даҳӣ: ".$a;
echo "<br>Қисми воҳидӣ: ".$b;
echo "<br>Суммаи рақамҳо: ".$sum;
echo "<br>ҲосилиЗарби рақамҳо: ".$zarb;
?>