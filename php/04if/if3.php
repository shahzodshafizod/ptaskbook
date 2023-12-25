<?php
$N = mt_rand(-100, 100);
echo "Адади додашуда: ".$N;
if($N > 0)
	$N++;
else if($N < 0)
	$N -= 2;
else
	$N = 10;
echo "<br>Адад пас аз иҷрои шарт: ".$N;
?>