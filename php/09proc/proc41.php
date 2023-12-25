<?php
function Sin1($x, $e){
        $dar = $x;
		$sum = $x;
        $fact = 1;
		$i = 1;
		$alomat = 1;
        while($dar / $fact > $e){
                $alomat *= -1;
                $fact *= (2*$i) * (2*$i+1);
                $dar *= $x * $x;
                $sum += $alomat * $dar / $fact;
                $i++;
        }
		$sum = number_format($sum, 2);
        return $sum;
}
$hisob = 1;
$x = mt_rand(-10, 10);
echo "x = ".$x."<br>";
$e = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
while($hisob <= 6){
	$index = mt_rand(0, 8);
	$nat = Sin1($x, $e[$index]);
	echo "<br>e = ".$e[$index];
	echo "<br>Sin1(".$x.", ".$e[$index].") = ".$nat."<br><br>";
	$hisob++;
}
?>