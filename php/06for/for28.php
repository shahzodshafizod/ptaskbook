<?php
$juft = 2;
$toq = 1;
$N = mt_rand(1, 50);
$X = array(-0.9, -0.8, -0.7, -0.6, -0.5, -0.4, -0.3, -0.2, -0.1, 0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$index = mt_rand(0, 18);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X[$index]."<br>";
$sum = 1 + $X[$index] / 2;
for($i = 2; $i <= $N; $i++){
	$juft *= (2 * $i);
	$toq *= (2 * $i - 3);
	$sum += pow(-1, $i - 1) * $toq * pow($X[$index], $i) / $juft;
}
echo "Cумма = ".number_format($sum, 2);
echo "<br>sqrt(".(1 + $X[$index]).") = ".number_format(sqrt(1 + $X[$index]), 2);
?>