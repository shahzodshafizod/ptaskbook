<?php
function SumRange($A, $B){
        $sum = 0;
        if($A < $B)
			for($i = $A; $i <= $B; $i++)
                $sum += $i;
        return $sum;
}
$A = mt_rand(1, 100);
$B = mt_rand(1, 100);
$C = mt_rand(1, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
$AB = SumRange($A, $B);
$BC = SumRange($B, $C);
echo "<br>Сумма аз A то B: ".$AB;
echo "<br>Сумма аз B то C: ".$BC;
?>