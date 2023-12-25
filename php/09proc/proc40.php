<?php
function Exp1($x, $e){
        $dar = 1;
		$sum = 1;
        $fact = 1;
		$i = 1;
        while($dar / $fact > $e){
                $dar *= $x;
                $fact *= $i;
                $sum += $dar / $fact;
                $i++;
        }
        return $sum;
}
$hisob = 1;
$x = mt_rand(-10, 10);
echo "x = ".$x;
$e = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
while($hisob <= 6){
	$index = mt_rand(0, 8);
	$nat = Exp1($x, $e[$index]);
	$nat = number_format($nat, 2);
	echo "<br>e = ".$e[$index];
	echo "<br>Exp1(".$x.",".$e[$index].") = ".$nat."<br><br>";
	$hisob++;
}
?>