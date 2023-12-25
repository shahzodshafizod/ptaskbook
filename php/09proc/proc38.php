<?php
function Power2($A, $N){
        $dar = 1;
        for($co = 1; $co <= $N; $co++)
                $dar *= $A;
        return $dar;
}
$A = mt_rand(-20, 20);
$K = mt_rand(-10, 10);
$L = mt_rand(-10, 10);
$M = mt_rand(-10, 10);
echo "A = ".$A;
echo "<br>K = ".$K;
echo "<br>L = ".$L;
echo "<br>M = ".$M;
echo "<br>A^K = ".Power2($A, $K);
echo "<br>A^L = ".Power2($A, $L);
echo "<br>A^M = ".Power2($A, $M);
?>