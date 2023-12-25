<?php
$hisob = 0;
$N = mt_rand(1, 100);
$K = mt_rand(1, 100);
echo "N = ".$N;
echo "<br>K = ".$K;
while($N >= $K){
	$N -= $K;
	$hisob++;
}
echo "<br>Тақсим = ".$hisob;
echo "<br>Бақия = ".$N;
?>