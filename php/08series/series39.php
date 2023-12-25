<?php
$hisob = 1;
$i = 3;
$miq = 0;
$mant = true;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
while($hisob <= $K){
	echo "Nabor".$hisob."<br>";
	do{
		$a = mt_rand(-100, 100);
		$b = mt_rand(-100, 100);
	}
	while(($a == 0) || ($b == 0));
	echo "n1 = ".$a."<br>";
	echo "n2 = ".$b."<br>";
	do{
		$c = mt_rand(-100, 100);
		echo "n".$i." = ".$c."<br>";
		if($c != 0){
			if(($b > $a) && ($b < $c) || ($b < $a) && ($b > $c))
				$mant = false;
		}
		$a = $b;
		$b = $c;
		$i++;
	}
	while($c != 0);
	if($mant)
		$miq++;
	$i = 3;
	$mant = true;
	$hisob++;
}
echo "Миқдори маҷмӯъҳои аррашакл = ".$miq;
 ?>