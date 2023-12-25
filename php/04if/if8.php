<?php
do{
	$a = mt_rand(-100, 100);
	$b = mt_rand(-100, 100);
}
while($a == $b);
$min = $a;
$max = $b;
if($a > $b){
	$min = $b;
	$max = $a;
}
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Адади калонтарин: ".$max;
echo "<br>Адади хурдтарин: ".$min;
?>