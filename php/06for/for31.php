<?php
$An = 2;
$N = mt_rand(1, 50);
echo "Mиқдори ададҳои пайдарпайӣ: ".$N;
echo "<br>Қиматҳои пайдарпайиҳо:<br>";
for($i = 1; $i <= $N; $i++){
	$An = 2 + 1 / $An;
	echo "A".$i." = ".number_format($An, 2)."<br>";
}
?>