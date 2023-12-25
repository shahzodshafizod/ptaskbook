<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$min_index = 1;
if($a > $b)
	$min_index = 2;
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Рақами адади хурдтарин: ".$min_index;
?>