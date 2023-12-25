<?php
do{ 
	$N = mt_rand();
}
while($N < 1000);
$m = $N % 1000;
$sadi = intval($m / 100);
echo "Адади аз 999 калон: ".$N;
echo "<br>Қисми садӣ: ".$sadi;
?>