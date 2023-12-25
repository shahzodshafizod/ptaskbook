<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$c = mt_rand(-100, 100);
echo "Aдади якӯм: ".$a;
echo "<br>Aдади дуюм: ".$b;
echo "<br>Aдади сеюм: ".$c;
if(($a < $b) && ($a < $c))
	$sum = $b + $c;
else if(($a > $b) && ($b < $c))
	$sum = $a + $c;
else
	$sum = $a + $b;
echo "<br>Суммаи ду ададҳои калонтарин: ".$sum;
?>