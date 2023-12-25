<?php
$sum = 0;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= $N; $i++)
	$sum += pow(-1, $i + 1) * (1 + $i * 0.1);
echo "Сумма: ".$sum;
?>