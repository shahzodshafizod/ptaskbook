<?php
$zarb = 1;
$hisob = 1;
while($hisob <= 10){
	$N = mt_rand(1, 100) / 4;
	$N = number_format($N, 2);
	echo "N".$hisob." = ".$N."<br>";
	$zarb *= $N;
	$hisob++;
}
echo "Зарби ададҳо = ".$zarb;
?>