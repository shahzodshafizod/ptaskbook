<?php
function KTU2($A, $B){
        while(($A != 0) && ($B != 0))
                if($A > $B)
                        $A %= $B;
                else
                        $B %= $A;
        return ($A + $B);
}
function XKU2($A, $B){
        return $A * ($B / KTU2($A, $B));
}
$A = mt_rand(1, 100);
$B = mt_rand(1, 100);
$C = mt_rand(1, 100);
$D = mt_rand(1, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>D = ".$D;
echo "<br>XKU2(".$A.", ".$B.") = ".XKU2($A, $B);
echo "<br>XKU2(".$A.", ".$C.") = ".XKU2($A, $C);
echo "<br>XKU2(".$A.", ".$D.") = ".XKU2($A, $D);
?>