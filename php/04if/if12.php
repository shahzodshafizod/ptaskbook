<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$c = mt_rand(-100, 100);
if($a > $b)
	if($b < $c)
		$min = $b;
	else
		$min = $c;
else if($a < $b)
	if($a < $c)
		$min = $a;
	else
		$min = $c;
else if($a < $c)
	if($a < $b)
		$min = $a;
	else
		$min = $b;
echo "Aдади якӯм: ".$a;
echo "<br>Aдади дуюм: ".$b;
echo "<br>Aдади сеюм: ".$c;
echo "<br>Aдади хурдтарин: ".$min;
?>