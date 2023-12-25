<?php
function DigitCountSum($K, &$C, &$S){
	$S = 0;
	$C = 0;
	while($K > 0){
		$S += ($K % 10);
		$C++;
		$K = intval($K / 10);
	}
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand();
	DigitCountSum($N, $miq, $sum);
	echo "Адади додашуда: ".$N;
	echo "<br>Миқдори рақамҳои адад: ".$miq;
	echo "<br>Суммаи рақамҳои адад: ".$sum."<br><br>";
	$hisob++;
}
?>