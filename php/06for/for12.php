<?php
$zarb = 1;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= $N; $i++)
	$zarb *= (1 + $i * 0.1);
echo "ҲосилиЗарб: ".$zarb;
?>