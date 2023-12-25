<?php
$sum = 0;
$N = mt_rand(1, 50);
echo "N = ".$N."<br>";
for($i = 1; $i <= $N; $i++)
	$sum += 1 / $i;
echo "Сумма: ".number_format($sum, 2);
?>