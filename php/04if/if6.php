<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$max = $a;
if($a < $b)
	$max = $b;
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Адади калонтарин: ".$max;
?>