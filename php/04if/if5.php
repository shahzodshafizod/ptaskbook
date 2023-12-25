<?php
do{
	$a = mt_rand(-100, 100);
	$b = mt_rand(-100, 100);
	$c = mt_rand(-100, 100);
}
while(($a == 0) || ($b == 0) || ($c == 0));
$plus = 0;
$minus = 3;
if($a > 0)
	$plus++;
if($b > 0)
	$plus++;
if($c > 0)
	$plus++;
$minus -= $plus;
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Адади сеюм: ".$c;
echo "<br>Миқдори ададҳои мусбӣ: ".$plus;
echo "<br>Миқдори ададҳои манфӣ: ".$minus;
?>