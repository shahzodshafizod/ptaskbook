<?php
function Power1($A, $B){
        $nat = 0;
        if($A > 0)
                $nat = exp($B * log($A));
        return $nat;
}
$A = mt_rand(-20, 20);
$B = mt_rand(-20, 20);
$C = mt_rand(-20, 20);
$P = mt_rand(-10, 10);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>P = ".$P;
echo "<br>A^P = ".Power1($A, $P);
echo "<br>B^P = ".Power1($B, $P);
echo "<br>C^P = ".Power1($C, $P);
?>