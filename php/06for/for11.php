<?php
$sum = 0;
$N = mt_rand(1, 50);
echo "N = ".$N."<br>";
for($i = 0; $i <= $N; $i++)
	$sum += pow($N + i, 2);
echo "Сумма: ".$sum;
?>