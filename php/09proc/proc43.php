<?php
function Ln1($x, $e){
        $dar = $x;
		$sum = $x;
        $i = 1;
		$alomat = 1;
        while(abs($dar / $i) < $e){
                $i++;
                $dar *= $x;
                $alomat *= -1;
                $sum += $alomat * $dar / $i;
        }
		$sum = number_format($sum, 2);
        return $sum;
}
$x = mt_rand(-10, 10);
echo "x = ".$x."<br>";
$e = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$hisob = 1;
while($hisob <= 6){
	$index = mt_rand(0, 8);
	$nat = Ln1($x, $e[$index]);
	echo "<br>e = ".$e[$index];
	echo "<br>Ln1(1+".$x.") = ".$nat."<br><br>";
	$hisob++;
}
?>