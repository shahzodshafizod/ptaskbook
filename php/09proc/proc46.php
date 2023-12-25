<?php
function KTU2($A, $B){
        while(($A != 0) && ($B != 0))
                if($A > $B)
                        $A %= $B;
                else
                        $B %= $A;
        return ($A + $B);
}
$A = mt_rand(1, 100);
$B = mt_rand(1, 100);
$C = mt_rand(1, 100);
$D = mt_rand(1, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>D = ".$D;
echo "<br>KTU2(".$A.", ".$B.") = ".KTU2($A, $B);
echo "<br>KTU2(".$A.", ".$C.") = ".KTU2($A, $C);
echo "<br>KTU2(".$A.", ".$D.") = ".KTU2($A, $D);
?>