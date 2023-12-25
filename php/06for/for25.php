<?php
$sum = 0;
$N = mt_rand(1, 50);
$X = array(-0.9, -0.8, -0.7, -0.6, -0.5, -0.4, -0.3, -0.2, -0.1, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$index = mt_rand(0, 17);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X[$index]."<br>";
for($i = 1; $i <= $N; $i++)
	$sum += pow(-1, $i - 1) * pow($X[$index], $i) / $i;
echo "Сумма = ".number_format($sum, 2);
echo "<br>log(".$X.") = ".number_format(log($X[$index]), 2);
?>