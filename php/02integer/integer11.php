<?php
$N = mt_rand(100, 999);
$a = intval($N / 100);
$b = intval($N / 10) % 10;
$c = $N % 10;
$sum = $a + $b + $c;
$zarb = $a * $b * $c;
echo "Адади серақама: ".$N;
echo "<br>Қисми садӣ: ".$a;
echo "<br>Қисми даҳӣ: ".$b;
echo "<br>Қисми воҳидӣ: ".$c;
echo "<br>Суммаи рақамҳо: ".$sum;
echo "<br>ҲосилиЗарби рақамҳо: ".$zarb;
?>