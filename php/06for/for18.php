<?php
$sum = 0;
$N = mt_rand(1, 50);
$A = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади додашуда: ".$A."<br>";
for($i = 0; $i <= $N; $i++)
	$sum += pow(-1, $i) * pow($A, $i);
echo "Сумма = ".$sum;
?>