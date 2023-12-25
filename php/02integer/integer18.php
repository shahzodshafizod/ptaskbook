<?php
do{
	$N = mt_rand();
}
while($N < 1000);
$m = $N % 10000;
$hazor = intval($m / 1000);
echo "Адади аз 999 калон: ".$N;
echo "<br>Қисми ҳазорӣ: ".$hazor;
?>