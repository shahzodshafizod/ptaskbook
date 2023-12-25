<?php
$juft = $toq = 1;
$N = mt_rand(1, 50);
$X = array(-0.9, -0.8, -0.7, -0.6, -0.5, -0.4, -0.3, -0.2, -0.1, 0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$index = mt_rand(0, 18);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X[$index]."<br>";
$sum = $X[$index];
for($i = 1; $i <= $N; $i++){
	$toq *= (2 * $i - 1);
	$juft *= (2 * $i);
	$sum += $toq * pow($X[$index], 2 * $i + 1) / ($juft * (2 * $i + 1));
}
echo "Cумма = ".number_format($sum, 2);
echo "<br>arcsin(".$X.") = ".number_format(asin($X[$index]), 2);
?>