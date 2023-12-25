<?php
$hisob = 1;
$i = 2;
$miq = 0;
$afz = true;
$kam = true;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
while($hisob <= $K){
	echo "Nabor".$hisob."<br>";
	do{
		$a = mt_rand(-100, 100);
	}
	while($a == 0);
	echo "n1 = ".$a."<br>";
	do{
		$b = mt_rand(-100, 100);
		echo "n".$i." = ".$b."<br>";
		if($b != 0){
			if($a <= $b)
				$kam = false;
			if($a >= $b)
				$afz = false;
		}
		$a = $b;
		$i++;
	}
	while($b != 0);
	$mant = $afz || $kam;
	if($mant)
		$miq++;
	$i = 2;
	$kam = true;
	$afz = true;
	$hisob++;
}
echo "Миқдори маҷмӯъҳои талабкардашуда = ".$miq;
 ?>