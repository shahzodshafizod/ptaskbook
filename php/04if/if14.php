<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$c = mt_rand(-100, 100);
echo "Aдади якӯм: ".$a;
echo "<br>Aдади дуюм: ".$b;
echo "<br>Aдади сеюм: ".$c;
if($a > $b)
	if($b > $c){
		$min = $c;
		$max = $a;
	}
	else{
		$min = $b;
		if(a < $c)
			$max = $c;
		else
			$max = $a;
	}
else if($a < $b)
	if($a > $c){
		$min = $c;
		$max = $b;
	}
	else{
		$min = $a;
		if($c < $b)
			$max = $b;
		else
			$max = $c;
	}
else if($a > $c)
	if($c > $b){
		$min = $b;
		$max = $a;
	}
	else{
		$min = $c;
		if($a < $b)
			$max = $b;
		else
			$max = $a;
	}
echo "<br>Адади хурдтарин: ".$min;
echo "<br>Адади калонтарин: ".$max;
?>