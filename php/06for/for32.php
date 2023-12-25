<?php
$An = 1;
$N = mt_rand(1, 50);
echo "Mиқдори ададҳои пайдарпайӣ: ".$N;
echo "<br>қиматҳои пайдарпайӣ:<br><br>";
for($i = 1; $i <= $N; $i++){
	$An = ($An + 1) / $i;
	echo "A".$i." = ".number_format($An, 2)."<br>";
}
?>