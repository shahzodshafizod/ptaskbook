<?php
$hisob = 1;
$i = 2;
$miq = 0;
$mant = true;
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
		if(($a >= $b) && ($b != 0))
				$mant = false;
		$a = $b;
		$i++;
	}
	while($b != 0);
	if($mant)
		$miq++;
	$i = 2;
	$hisob++;
	$mant = true;
}
echo "Миқдори маҷмӯъҳои талабкардашуда = ".$miq;
?>