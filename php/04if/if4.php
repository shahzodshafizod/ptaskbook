<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$c = mt_rand(-100, 100);
$plus = 0;
if($a > 0)
	$plus++;
if($b > 0)
	$plus++;
if($c > 0)
	$plus++;
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Адади сеюм: ".$c;
echo "<br>Миқдори ададҳои мусбӣ: ".$plus;
?>